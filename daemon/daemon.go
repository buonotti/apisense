package daemon

import (
	"context"
	"os"
	"os/signal"
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
					err := d.Pipeline.RefreshItems()
					errors.HandleError(err)
				case SIGINT:
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
			for _, item := range result {
				for _, set := range item.ValidatorResults {
					log.DaemonLogger.Infof("Validation result for validator '%s' on endpoint %s (%s)", set.Validator, item.EndpointName, set.Url)
					if set.Error != nil {
						log.DaemonLogger.Errorf("Error: %s", set.Error)
					} else {
						log.DaemonLogger.Infof("Nothing to report")
					}
				}
			}
			time.Sleep(1 * time.Minute)
		}
	}
}
