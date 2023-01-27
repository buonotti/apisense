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
	err := writeStatus(UP)
	if err != nil {
		return err
	}

	err = writePid(os.Getpid())
	if err != nil {
		return err
	}

	log.DaemonLogger.Infof("daemon started")

	ctx := context.Background()
	ctx, contextCancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, unix.SIGINT, unix.SIGHUP, unix.SIGTERM)

	defer endRun(signalChan, contextCancel)

	go d.signalListener(signalChan, contextCancel, ctx)

	if runOnStart {
		log.DaemonLogger.Info("running work function on start")
		d.work()
	}

	cronScheduler := cron.New()
	workFunctionId, err := cronScheduler.AddFunc(viper.GetString("daemon.interval"), d.work)

	defer cronScheduler.Remove(workFunctionId)

	cronScheduler.Run()
	return nil
}

// signalListener listens on the given channel for signals or the end of the
// given context and then stops the daemon or reloads the config
func (d daemon) signalListener(signalChan chan os.Signal, cancel context.CancelFunc, ctx context.Context) {
	for {
		select {
		case s := <-signalChan:
			switch s {
			case unix.SIGHUP:
				log.DaemonLogger.Info("received SIGHUP, reloading configuration")
				errors.HandleError(d.Pipeline.RefreshItems())
			case unix.SIGINT:
				log.DaemonLogger.Info("received SIGINT, stopping daemon")
				errors.HandleError(writeStatus(DOWN))
				errors.HandleError(writePid(-1))
				cancel()
				os.Exit(0)
			case unix.SIGTERM:
				log.DaemonLogger.Info("received SIGTERM, stopping daemon")
				errors.HandleError(writeStatus(DOWN))
				errors.HandleError(writePid(-1))
				cancel()
				os.Exit(0)
			}
		case <-ctx.Done():
			log.DaemonLogger.Infof("context done, exiting signal handler")
			errors.HandleError(writeStatus(DOWN))
			errors.HandleError(writePid(-1))
			os.Exit(0)
		}
	}
}

// endRun runs at the end of the work function and cancels the context via the
// given cancel function and writes the daemon status
func endRun(signalChan chan os.Signal, cancel context.CancelFunc) {
	log.DaemonLogger.Info("stopping daemon")
	errors.HandleError(writeStatus(DOWN))
	errors.HandleError(writePid(-1))
	signal.Stop(signalChan)
	cancel()
}

// work runs the validation pipeline and logs the results
func (d daemon) work() {
	err := d.Pipeline.RefreshItems()
	errors.HandleError(err)

	result := d.Pipeline.Validate()

	d.logResults(result)

	reportData, err := conversion.Json().Convert(result)
	if err != nil {
		errors.HandleError(errors.CannotSerializeItemError.Wrap(err, "cannot serialize report"))
	}

	reportPath := validation.ReportLocation() + string(filepath.Separator) + time.Now().Format("02-01-2006T15:04:05.000Z") + ".report.json"
	err = os.WriteFile(reportPath, reportData, 0644)
	if err != nil {
		errors.HandleError(errors.CannotWriteFileError.Wrap(err, "cannot write report to file"))
	}

	log.DaemonLogger.Infof("generated report (%s)", reportPath)
}

// logResults logs the results of the validation pipeline to the output file or stdout using the log.DaemonLogger
func (d daemon) logResults(report validation.Report) {
	for _, validatedEndpoint := range report.Endpoints {
		for _, result := range validatedEndpoint.TestCaseResults {
			for _, validatorResult := range result.ValidatorResults {
				log.DaemonLogger.Infof("validation result for validator '%s' on endpoint %s (%s)", validatorResult.Name, validatedEndpoint.EndpointName, result.Url)
				if validatorResult.Message != "" {
					log.DaemonLogger.Errorf("error: %s", validatorResult.Message)
				} else {
					log.DaemonLogger.Infof("nothing to report")
				}
			}
		}
	}
}
