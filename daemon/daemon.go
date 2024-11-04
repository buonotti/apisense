package daemon

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/buonotti/apisense/v2/alerting"
	"github.com/buonotti/apisense/v2/conversion"
	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/filesystem/locations/directories"
	"github.com/buonotti/apisense/v2/log"
	"github.com/buonotti/apisense/v2/util"
	"github.com/buonotti/apisense/v2/validation/pipeline"
	"github.com/buonotti/apisense/v2/validation/validators"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
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

	log.DaemonLogger().Info("Daemon started")

	ctx := context.Background()
	ctx, contextCancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer endRun(signalChan, contextCancel)

	go d.signalListener(signalChan, contextCancel, ctx)

	if runOnStart {
		log.DaemonLogger().Info("Running validation on startup")
		d.work()
	}

	cronScheduler := cron.New()
	log.DaemonLogger().Debug("Starting with cron expr", "cron", viper.GetString("daemon.interval"))
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
			case os.Interrupt:
				log.DaemonLogger().Info("Received interrupt, stopping daemon")
				err := writeStatus(DownStatus)
				if err != nil {
					log.DaemonLogger().Fatal(err)
				}
				err = writePid(-1)
				if err != nil {
					log.DaemonLogger().Fatal(err)
				}
				cancel()
				os.Exit(0)
			}
		case <-ctx.Done():
			log.DaemonLogger().Info("Context done, exiting signal handler")
			err := writeStatus(DownStatus)
			if err != nil {
				log.DaemonLogger().Fatal(err)
			}
			err = writePid(-1)
			if err != nil {
				log.DaemonLogger().Fatal(err)
			}
			os.Exit(0)
		}
	}
}

// endRun runs at the end of the work function and cancels the context via the
// given cancel function and writes the daemon status
func endRun(signalChan chan os.Signal, cancel context.CancelFunc) {
	log.DaemonLogger().Info("Stopping daemon")
	err := writeStatus(DownStatus)
	if err != nil {
		log.DaemonLogger().Fatal(err)
	}
	err = writePid(-1)
	if err != nil {
		log.DaemonLogger().Fatal(err)
	}
	signal.Stop(signalChan)
	cancel()
}

// work runs the validation pipeline and logs the results
func (d daemon) work() {
	result, err := d.Pipeline.Validate()
	if err != nil {
		log.DaemonLogger().Fatal(err)
	}

	d.logResults(result)

	reportData, err := conversion.Json().Convert(result)
	if err != nil {
		log.DaemonLogger().Fatal(errors.CannotSerializeItemError.Wrap(err, "cannot serialize report"))
	}

	if len(result.Endpoints) > 0 {
		reportPath := filepath.FromSlash(directories.ReportsDirectory() + "/" + time.Time(result.Time).Format(util.ApisenseTimeFormat) + ".report.json")
		err = os.WriteFile(reportPath, reportData, 0o644)
		if err != nil {
			log.DaemonLogger().Fatal(errors.CannotWriteFileError.Wrap(err, "cannot write report to file"))
		}

		log.DaemonLogger().Info("Generated report", "path", reportPath)

		if viper.GetBool("daemon.notification.enabled") {
			alertData := alerting.AlertData{
				Time:        result.Time,
				ErrorAmount: countErrors(result),
			}
			err = alerting.SendAlert(alertData)
			if err != nil {
				log.DaemonLogger().Error(err)
			}
		}
	} else {
		log.DaemonLogger().Info("No endpoints to validate")
	}

	err = d.cleanupReports()
	if err != nil {
		log.DaemonLogger().Error(err)
	}
}

// countErrors counts the amount of errors in the report
func countErrors(report pipeline.Report) uint {
	count := 0
	for _, endpoint := range report.Endpoints {
		for _, testCase := range endpoint.TestCaseResults {
			for _, validatorRes := range testCase.ValidatorResults {
				if validatorRes.Status == validators.ValidatorStatusFail {
					count += 1
				}
			}
		}
	}
	return uint(count)
}

// cleanupReports cleans up old reports in the report directory
func (d daemon) cleanupReports() error {
	if !viper.GetBool("daemon.discard.enabled") {
		return nil
	}

	log.DaemonLogger().Info("Cleaning up reports")
	files, err := os.ReadDir(filepath.FromSlash(directories.ReportsDirectory()))
	if err != nil {
		return err
	}
	for _, file := range files {
		fName := file.Name()
		if !strings.HasSuffix(fName, ".report.json") {
			continue
		}
		fName = strings.TrimSuffix(fName, ".report.json")
		fTime, err := time.Parse(util.ApisenseTimeFormat, fName)
		if err != nil {
			log.DaemonLogger().Warn("Cannot parse report name, skipping", "report", file.Name())
			continue
		}
		maxTime := viper.GetDuration("daemon.discard.max_lifetime")
		if time.Since(fTime) > maxTime {
			err = os.Remove(filepath.FromSlash(directories.ReportsDirectory() + "/" + file.Name()))
			if err != nil {
				return errors.CannotRemoveFileError.Wrap(err, "cannot remove report file")
			}
			log.DaemonLogger().Info("Removed report because it was too old", "filename", file.Name())
		}
	}
	return nil
}

// logResults logs the results of the validation pipeline to the output file or stdout using the log.DaemonLogger()
func (d daemon) logResults(report pipeline.Report) {
	for _, validatedEndpoint := range report.Endpoints {
		for _, result := range validatedEndpoint.TestCaseResults {
			for _, validatorResult := range result.ValidatorResults {
				if validatorResult.Status == validators.ValidatorStatusFail {
					log.DaemonLogger().Error("Validation failed",
						"validator", validatorResult.Name,
						"endpoint", validatedEndpoint.EndpointName,
						"url", result.Url,
						"message", validatorResult.Message)
				} else {
					log.DaemonLogger().Info("Validation succeeded",
						"validator", validatorResult.Name,
						"endpoint", validatedEndpoint.EndpointName,
						"url", result.Url)
				}
			}
		}
	}
}
