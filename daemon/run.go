package daemon

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buonotti/odh-data-monitor/log"
)

func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					// TODO reload config
				case os.Interrupt:
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done():
				os.Exit(1)
			}
		}
	}()
	if err := work(ctx); err != nil {
		return err
	}
	return nil
}

func work(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			log.DaemonLogger.Info("Daemon is running")
			time.Sleep(time.Second)
		}
	}
}
