package filesystem

import (
	"os"
	"time"

	"github.com/buonotti/apisense/v2/errors"
)

type FileWatcher struct {
	content string
	file    string
	Events  chan bool
}

func NewFileWatcher() FileWatcher {
	return FileWatcher{
		file:    "",
		content: "",
		Events:  make(chan bool),
	}
}

func (f *FileWatcher) Start() error {
	for {
		if f.file == "" {
			return errors.WatcherError.New("no file specified")
		}

		dat, err := os.ReadFile(f.file)
		if err != nil {
			return errors.CannotReadFileError.Wrap(err, "cannot read file: "+f.file)
		}

		if f.content != string(dat) && f.content != "" {
			f.Events <- true
			time.Sleep(1 * time.Second)
		}

		f.content = string(dat)
		time.Sleep(500 * time.Millisecond)
	}
}

func (f *FileWatcher) AddFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		return errors.CannotReadFileError.Wrap(err, "cannot resolve file: "+file)
	}

	f.file = file
	return nil
}
