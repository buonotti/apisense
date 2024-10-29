package tui

import (
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/files"
	"github.com/buonotti/apisense/log"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fsnotify/fsnotify"
)

func Run() error {

	p := tea.NewProgram(TuiModule(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.TuiLogger().Fatal(errors.UnknownError.Wrap(err, "unable to create watcher"))
	}

	go func() {
		for {
			select {
			// Watch for events
			case event, ok := <-watcher.Events:
				if !ok {
					log.TuiLogger().Warn("bad watcher event")
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					p.Send(watcherMsg{})
				}

			// Watch for errors
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.TuiLogger().Fatal(errors.WatcherError.Wrap(err, "unable to create watcher"))
			}
		}
	}()

	//err = watcher.Add(directories.DaemonDirectory())
	err = watcher.Add(files.DaemonPidFile())
	if err != nil {
		log.TuiLogger().Fatal(errors.WatcherError.Wrap(err, "unable to start watcher"))
	}

	defer watcher.Close()
	_, perr := p.Run()
	return perr
}
