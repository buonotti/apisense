package repo

import (
	"os"
	"path/filepath"
	"time"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/filesystem/locations/files"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/goccy/go-json"
	"github.com/ldez/go-git-cmd-wrapper/v2/add"
	"github.com/ldez/go-git-cmd-wrapper/v2/clone"
	"github.com/ldez/go-git-cmd-wrapper/v2/commit"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	"github.com/plus3it/gorecurcopy"
)

// Lockfile
type Lockfile struct {
	LastUpdate util.ApisenseTime
	Templates  map[string]string
}

// loadLockfile loads the lockfile from disk
func loadLockfile() (Lockfile, error) {
	if !util.Exists(files.PkgLockFile()) {
		defaultLockfile := Lockfile{LastUpdate: util.ApisenseTime(time.Now().UTC())}
		err := saveLockfile(defaultLockfile)
		if err != nil {
			return Lockfile{}, err
		}
		return defaultLockfile, nil
	}
	content, err := os.ReadFile(files.PkgLockFile())
	if err != nil {
		return Lockfile{}, errors.CannotReadFileError.Wrap(err, "cannot read lockfile")
	}
	var lockfile Lockfile
	err = json.Unmarshal(content, &lockfile)
	if err != nil {
		return Lockfile{}, errors.CannotUmarshalError.Wrap(err, "cannot umarshal lockfile")
	}
	return lockfile, err
}

// saveLockfile saves the lockfile to disk
func saveLockfile(lockfile Lockfile) error {
	marshalled, _ := json.MarshalIndent(lockfile, "", "  ")
	err := os.WriteFile(files.PkgLockFile(), marshalled, os.ModePerm)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot save lockfile")
	}
	return nil
}

// Update updates all the templates
func Update(force bool) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	remoteTemplates, err := DiscoverTemplates()
	if err != nil {
		return err
	}

	for lang, url := range remoteTemplates {
		_, ok := lockfile.Templates[lang]
		if !ok || force {
			log.DefaultLogger().Info("Found new language. Adding to templates", "lang", lang)
			if lockfile.Templates == nil {
				lockfile.Templates = make(map[string]string)
			}
			lockfile.Templates[lang] = url
		}
	}

	for lang, url := range lockfile.Templates {
		langPath := filepath.FromSlash(directories.ValidatorRepoDirectory() + "/" + lang)
		if !util.Exists(langPath) {
			log.DefaultLogger().Info("Template dir does not exist. Cloning template", "lang", lang, "dir", langPath, "url", url)
			out, err := git.Clone(clone.Repository(url), clone.Directory(langPath))
			if err != nil {
				return errors.CannotCloneRepoError.Wrap(err, "cannot clone git repo. is git installed? error: %s", out)
			}
		} else {
			err := os.Chdir(langPath)
			if err != nil {
				return errors.CannotChangeDirectoryError.WrapWithNoMessage(err)
			}
			out, err := git.Fetch()
			if err != nil {
				return errors.CannotFetchRepoError.Wrap(err, "cannot fetch git repo. error: %s", out)
			}
			out, err = git.Pull()
			if err != nil {
				return errors.CannotPullRepoError.Wrap(err, "cannot pull git repo. error: %s", out)
			}
			log.DefaultLogger().Info("Template dir does exist. Updating", "lang", lang, "dir", langPath, "url", url)
		}
	}

	lockfile.LastUpdate = util.ApisenseTime(time.Now().UTC())
	return saveLockfile(lockfile)
}

// Create creates a new validator from the given template
func Create(lang string, name string, force bool) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	_, ok := lockfile.Templates[lang]
	if !ok {
		return errors.LanguageNotAvailableError.New("language '%s' is not installed. run 'apisense templates update' to update local templates", lang)
	}
	langPath := filepath.FromSlash(directories.ValidatorRepoDirectory() + "/" + lang)
	destPath := filepath.FromSlash(directories.ValidatorCustomDirectory() + "/" + name)
	if util.Exists(destPath) {
		if force {
			err := os.RemoveAll(destPath)
			if err != nil {
				return errors.CannotRemoveFileError.WrapWithNoMessage(err)
			}
		} else {
			log.DefaultLogger().Warn("Destination path already exists. Aborting", "path", destPath)
			return nil
		}
	}
	err = gorecurcopy.CopyDirectory(langPath, destPath)
	if err != nil {
		return errors.CannotCopyDirectoryError.WrapWithNoMessage(err)
	}
	err = os.RemoveAll(destPath + "/.git")
	if err != nil {
		return errors.CannotRemoveFileError.WrapWithNoMessage(err)
	}

	err = os.Chdir(destPath)
	if err != nil {
		return errors.CannotChangeDirectoryError.WrapWithNoMessage(err)
	}

	out, err := git.Init()
	if err != nil {
		return errors.CannotInitRepoError.Wrap(err, "cannot init repo. error: %s", out)
	}

	out, err = git.Add(add.PathSpec("."))
	if err != nil {
		return errors.CannotAddError.Wrap(err, "cannot add to working tree. error: %s", out)
	}

	out, err = git.Commit(commit.Message("Initial commit"))
	if err != nil {
		return errors.CannotCommitError.Wrap(err, "cannot commit changes. error: %s", out)
	}

	log.DefaultLogger().Info("Created new validator from template", "lang", lang, "path", destPath)
	return nil
}

// AddCustomRepo adds a language with a custom git repo to the lockfile
func AddCustomRepo(lang string, url string) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	_, ok := lockfile.Templates[lang]
	if ok {
		return errors.LanguageAlreadyExistsError.New("language '%s' already has a template", lang)
	}
	if lockfile.Templates == nil {
		lockfile.Templates = make(map[string]string)
	}
	lockfile.Templates[lang] = url
	return saveLockfile(lockfile)
}
