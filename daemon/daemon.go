package daemon

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
	"github.com/buonotti/odh-data-monitor/validation"
)

type daemon struct {
	Pipeline *validation.Pipeline
}

func (d daemon) run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, SIGINT, SIGHUP)

	defer func() {
		log.DaemonLogger.Info("Stopping daemon")
		errors.HandleError(writeStatus(DOWN))
		errors.HandleError(writePid(-1))
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case SIGHUP:
					log.DaemonLogger.Info("Received SIGHUP, reloading configuration")
					err := d.Pipeline.RefreshItems()
					errors.HandleError(err)
				case SIGINT:
					log.DaemonLogger.Info("Received SIGINT, stopping daemon")
					errors.HandleError(writeStatus(DOWN))
					errors.HandleError(writePid(-1))
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done():
				errors.HandleError(writeStatus(DOWN))
				errors.HandleError(writePid(-1))
				os.Exit(1)
			}
		}
	}()
	if err := d.work(ctx); err != nil {
		return err
	}
	return nil
}

func (d daemon) work(ctx context.Context) error {
	err := writeStatus(UP)
	if err != nil {
		return err
	}
	err = writePid(os.Getpid())
	if err != nil {
		return err
	}
	defer func() {
		err := writeStatus(DOWN)
		errors.HandleError(err)
		err = writePid(-1)
		errors.HandleError(err)
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := d.Pipeline.RefreshItems()
			if err != nil {
				return err
			}
			result := d.Pipeline.Validate()
			d.LogResults(result)
			reportData, err := json.Marshal(result)
			if err != nil {
				return errors.CannotSerializeItemError.Wrap(err, "Cannot serialize report")
			}
			reportPath := validation.ReportLocation() + string(filepath.Separator) + time.Now().Format("02-01-2006@15:04") + ".report.json"
			err = os.WriteFile(reportPath, reportData, 0644)
			if err != nil {
				return errors.CannotWriteFileError.Wrap(err, "Cannot write report to file")
			}
			log.DaemonLogger.Infof("Generated report (%s)", reportPath)
			time.Sleep(1 * time.Minute)
		}
	}
}

func (d daemon) LogResults(report validation.Report) {
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
