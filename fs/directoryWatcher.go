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
	directoryEntries, err := os.ReadDir(directory)
	if err != nil {
		return nil, errors.CannotReadDirectoryError.Wrap(err, "cannot read directory: "+directory)
	}

	files := make([]string, 0)
	for _, file := range directoryEntries {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	return files, nil
}

func (f *DirectoryWatcher) Start() error {
	for {
		if f.directory == "" {
			return errors.WatcherError.New("no directory specified")
		}

		dat, err := os.ReadDir(f.directory)
		if err != nil {
			return errors.CannotReadDirectoryError.Wrap(err, "cannot read directory: "+f.directory)
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

func (f *DirectoryWatcher) SetDirectory(directory string) error {
	if _, err := os.Stat(directory); err != nil {
		return errors.CannotReadDirectoryError.Wrap(err, "cannot resolve directory: "+directory)
	}

	f.directory = directory
	return nil
}
