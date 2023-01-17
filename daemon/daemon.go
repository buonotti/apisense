package daemon

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"golang.org/x/sys/unix"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation"
)

// daemon provides simple daemon operations, and it holds the validation.Pipeline
// to perform the validation of the data
type daemon struct {
	Pipeline *validation.Pipeline // Pipeline is the validation pipeline the daemon uses
}

// run starts the daemon and with it the cron scheduler that runs the work
// function in the preferred interval. It also starts a goroutine to listen to
// incoming signals and stop the daemon when a SIGINT is received, and it reloads
// the configuration when a SIGHUP is received.
func (d daemon) run(runOnStart bool) error {
	// write the daemon status to the control file
	err := writeStatus(UP)
	if err != nil {
		return err
	}

	// write the daemon pid to the control file
	err = writePid(os.Getpid())
	if err != nil {
		return err
	}

	log.DaemonLogger.Infof("Daemon started")

	// create a context to be used to stop the daemon
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// create a channel to listen to signals and set the signals the daemon reacts to
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, unix.SIGINT, unix.SIGHUP, unix.SIGTERM)

	// when the work function ends write the daemon status and pid to the control file and stop listening to signals
	// and then cancel the context with the cancel function we got upon creating the context
	defer func() {
		log.DaemonLogger.Info("Stopping daemon")
		errors.HandleError(writeStatus(DOWN))
		errors.HandleError(writePid(-1))
		signal.Stop(signalChan)
		cancel()
	}()

	// Start a goroutine that reacts to the incoming signals. The function runs an infinite loop that
	// reads from the signal channel and reacts to the signals it receives. When a SIGINT is received
	// it stops the daemon, when a SIGHUP is received it reloads the configuration.
	// it also reads from the context and if the context is marked as done it exits
	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case unix.SIGHUP:
					log.DaemonLogger.Info("Received SIGHUP, reloading configuration")
					errors.HandleError(d.Pipeline.RefreshItems())
				case unix.SIGINT:
					log.DaemonLogger.Info("Received SIGINT, stopping daemon")
					errors.HandleError(writeStatus(DOWN))
					errors.HandleError(writePid(-1))
					cancel()
					os.Exit(0)
				case unix.SIGTERM:
					log.DaemonLogger.Info("Received SIGTERM, stopping daemon")
					errors.HandleError(writeStatus(DOWN))
					errors.HandleError(writePid(-1))
					cancel()
					os.Exit(0)
				}
			case <-ctx.Done():
				log.DaemonLogger.Infof("Context done, exiting signal handler")
				errors.HandleError(writeStatus(DOWN))
				errors.HandleError(writePid(-1))
				os.Exit(0)
			}
		}
	}()

	if runOnStart {
		log.DaemonLogger.Info("Running work function on start")
		d.work()()
	}

	// create a new cron scheduler that runs the work function in the preferred interval.
	// The interval is read from the configuration file
	// It then runs the scheduler. Run is a blocking function, so it will run until the context is marked as done
	c := cron.New()
	id, err := c.AddFunc(viper.GetString("daemon.interval"), d.work())
	defer c.Remove(id)
	c.Run()
	return nil
}

// work returns a function that runs the validation pipeline and logs the results
func (d daemon) work() func() {
	return func() {
		// refresh the pipeline items to make sure we are using the latest configuration and have the latest data
		err := d.Pipeline.RefreshItems()
		errors.HandleError(err)

		// validate the pipeline items
		result := d.Pipeline.Validate()

		// log the results
		d.logResults(result)

		// serialize the results to json and write them to the report file with the current timestamp as name
		reportData, err := conversion.Json().Convert(result)
		if err != nil {
			errors.HandleError(errors.CannotSerializeItemError.Wrap(err, "Cannot serialize report"))
		}

		reportPath := validation.ReportLocation() + string(filepath.Separator) + time.Now().Format("02-01-2006T15:04:05.000Z") + ".report.json"
		err = os.WriteFile(reportPath, reportData, 0644)
		if err != nil {
			errors.HandleError(errors.CannotWriteFileError.Wrap(err, "Cannot write report to file"))
		}

		log.DaemonLogger.Infof("Generated report (%s)", reportPath)
	}
}

// logResults logs the results of the validation pipeline to the output file or stdout using the log.DaemonLogger
func (d daemon) logResults(report validation.Report) {
	for _, validatedEndpoint := range report.Results {
		for _, result := range validatedEndpoint.Results {
			for _, validatorResult := range result.ValidatorsOutput {
				log.DaemonLogger.Infof("Validation result for validator '%s' on endpoint %s (%s)", validatorResult.Validator, validatedEndpoint.EndpointName, result.Url)
				if validatorResult.Error != "" {
					log.DaemonLogger.Errorf("Error: %s", validatorResult.Error)
				} else {
					log.DaemonLogger.Infof("Nothing to report")
				}
			}
		}
	}
}
