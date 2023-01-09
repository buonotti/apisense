package tui

import (
	"os"
	"time"

	"github.com/buonotti/odh-data-monitor/errors"
)

type fileWatcher struct {
	content string
	file    string
	Events  chan bool
	test    int
}

func NewFileWatcher() fileWatcher {
	return fileWatcher{
		file:    "",
		content: "",
		test:    0,
		Events:  make(chan bool),
	}
}

func (f *fileWatcher) Start() error {
	for {
		if f.file == "" {
			return errors.WatcherError.New("No file specified")
		}
		dat, err := os.ReadFile(f.file)
		if err != nil {
			return errors.WatcherError.Wrap(err, "Error in watcher file reading")
		}
		if f.content != string(dat) && f.content != "" {
			f.Events <- true
			time.Sleep(5 * time.Second)
			f.test++
		}
		f.content = string(dat)
		time.Sleep(500 * time.Millisecond)
	}
}

func (f *fileWatcher) AddFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		return errors.WatcherError.Wrap(err, "Cannot resolve file")
	}
	f.file = file
	return nil
}
