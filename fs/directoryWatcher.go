package fs

import (
	"os"
	"time"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
)

type DirectoryWatcher struct {
	directory string
	files     []string
	Events    chan string
}

func NewDirectoryWatcher() DirectoryWatcher {
	return DirectoryWatcher{
		directory: "",
		files:     make([]string, 0),
		Events:    make(chan string),
	}
}

func NewDirectoryWatcherWithFiles(directory string) (DirectoryWatcher, error) {
	files, err := loadFiles(directory)
	if err != nil {
		return DirectoryWatcher{}, err
	}
	return DirectoryWatcher{
		directory: directory,
		files:     files,
		Events:    make(chan string),
	}, nil
}

func loadFiles(directory string) ([]string, error) {
	dat, err := os.ReadDir(directory)
	if err != nil {
		return nil, errors.WatcherError.Wrap(err, "Error in watcher directory reading")
	}
	files := make([]string, 0)
	for _, file := range dat {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

func (f *DirectoryWatcher) Start() error {
	for {
		if f.directory == "" {
			return errors.WatcherError.New("No directory specified")
		}
		dat, err := os.ReadDir(f.directory)
		if err != nil {
			return errors.WatcherError.Wrap(err, "Error in watcher directory reading")
		}
		for _, file := range dat {
			if !file.IsDir() {
				if !util.Contains(f.files, file.Name()) {
					f.Events <- file.Name()
					f.files = append(f.files, file.Name())
					time.Sleep(1 * time.Second)
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (f *DirectoryWatcher) Watch(directory string) error {
	if _, err := os.Stat(directory); err != nil {
		return errors.WatcherError.Wrap(err, "Cannot resolve directory")
	}
	f.directory = directory
	return nil
}
