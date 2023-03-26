package daemon

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/pipeline"
)

// daemon provides simple daemon operations, and it holds the validation.Pipeline
// to perform the validation of the data
type daemon struct {
	Pipeline *pipeline.Pipeline // Pipeline is the validation pipeline the daemon uses
}

// run starts the daemon and with it the cron scheduler that runs the work
// function in the preferred interval. It also starts a goroutine to listen to
// incoming signals and stop the daemon when a SIGINT is received, and it reloads
// the configuration when a SIGHUP is received.
func (d daemon) run(runOnStart bool) error {
	err := writeStatus(UpStatus)
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
	signal.Notify(signalChan, os.Kill)

	defer endRun(signalChan, contextCancel)

	go d.signalListener(signalChan, contextCancel, ctx)

	if runOnStart {
		log.DaemonLogger.Info("running work function on start")
		d.work()
	}

	cronScheduler := cron.New()
	workFunctionId, err := cronScheduler.AddFunc(viper.GetString("daemon.interval"), d.work)
	if err != nil {
		return errors.CannotAddWorkFunctionToCronSchedulerError.Wrap(err, "cannot add work function to cron scheduler")
	}

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
			case os.Kill:
				log.DaemonLogger.Info("received SIGKILL, stopping daemon")
				errors.CheckErr(writeStatus(DownStatus))
				errors.CheckErr(writePid(-1))
				cancel()
				os.Exit(0)
			}
		case <-ctx.Done():
			log.DaemonLogger.Infof("context done, exiting signal handler")
			errors.CheckErr(writeStatus(DownStatus))
			errors.CheckErr(writePid(-1))
			os.Exit(0)
		}
	}
}

// endRun runs at the end of the work function and cancels the context via the
// given cancel function and writes the daemon status
func endRun(signalChan chan os.Signal, cancel context.CancelFunc) {
	log.DaemonLogger.Info("stopping daemon")
	errors.CheckErr(writeStatus(DownStatus))
	errors.CheckErr(writePid(-1))
	signal.Stop(signalChan)
	cancel()
}

// work runs the validation pipeline and logs the results
func (d daemon) work() {
	err := d.Pipeline.Reload()
	errors.CheckErr(err)

	result := d.Pipeline.Validate()

	d.logResults(result)

	reportData, err := conversion.Json().Convert(result)
	if err != nil {
		errors.CheckErr(errors.CannotSerializeItemError.Wrap(err, "cannot serialize report"))
	}

	if len(result.Endpoints) > 0 {
		reportPath := filepath.FromSlash(directories.ReportsDirectory() + "/" + time.Now().Format(pipeline.ReportTimeFormat) + ".report.json")
		err = os.WriteFile(reportPath, reportData, 0644)
		if err != nil {
			errors.CheckErr(errors.CannotWriteFileError.Wrap(err, "cannot write report to file"))
		}

		log.DaemonLogger.Infof("generated report (%s)", reportPath)
	} else {
		log.DaemonLogger.Info("no endpoints to validate")
	}

	errors.CheckErr(d.cleanupReports())
}

func (d daemon) cleanupReports() error {
	if !viper.GetBool("daemon.discard.enabled") {
		return nil
	}

	log.DaemonLogger.Info("cleaning up reports")
	files, err := os.ReadDir(filepath.FromSlash(directories.ReportsDirectory()))
	errors.CheckErr(err)
	for _, file := range files {
		fName := file.Name()
		if !strings.HasSuffix(fName, ".report.json") {
			continue
		}
		fName = strings.TrimSuffix(fName, ".report.json")
		fTime, err := time.Parse(pipeline.ReportTimeFormat, fName)
		if err != nil {
			log.DaemonLogger.Warnf("cannot parse report name %s, skipping", file.Name())
			return nil
		}
		maxTime := viper.GetDuration("daemon.discard.max_lifetime")
		if time.Since(fTime) > maxTime {
			err = os.Remove(filepath.FromSlash(directories.ReportsDirectory() + "/" + file.Name()))
			if err != nil {
				return errors.CannotRemoveFileError.Wrap(err, "cannot remove report file")
			}
			log.DaemonLogger.Infof("removed report %s because it was too old", file.Name())
		}
	}
	return nil
}

// logResults logs the results of the validation pipeline to the output file or stdout using the log.DaemonLogger
func (d daemon) logResults(report pipeline.Report) {
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
