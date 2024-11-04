package pkg

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/filesystem/locations/files"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/ldez/go-git-cmd-wrapper/v2/add"
	"github.com/ldez/go-git-cmd-wrapper/v2/checkout"
	"github.com/ldez/go-git-cmd-wrapper/v2/clone"
	"github.com/ldez/go-git-cmd-wrapper/v2/commit"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	"github.com/plus3it/gorecurcopy"
)

// RepoType is the type of the remote repo
type RepoType string

const (
	GitHub RepoType = "github" // GitHub defines the remote is a github profile/org
)

// TemplateEntry holds the information for a single template
type TemplateEntry struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
	Commit string `json:"commit"`
}

// RepoEntry holds the information of a remote template
type RepoEntry struct {
	Url  string   `json:"url"`
	Type RepoType `json:"type"`
}

type Lockfile struct {
	LastUpdate util.ApisenseTime        `json:"lastUpdate"` // LastUpdate is when the last update was run
	Repos      map[string]RepoEntry     `json:"repos"`      // Repos are the remote repos to look for new templates
	Templates  map[string]TemplateEntry `json:"templates"`  // Templates are the available templates
}

// BUONOTTI_ENDPOINT is the official repo
const BUONOTTI_ENDPOINT string = "https://api.github.com/orgs/buonotti/repos"

// DiscoverTemplates loads all templates from a repo
func DiscoverTemplates(repo RepoEntry) (map[string]TemplateEntry, error) {
	if repo.Type != GitHub {
		return nil, errors.InvalidRepoTypeError.New("invalid repo type: %s", repo.Type)
	}
	resp, err := resty.New().R().Get(repo.Url)
	if err != nil {
		return nil, errors.GithubUnreachableError.Wrap(err, "failed to request endpoints")
	}
	var repos []map[string]any
	err = json.Unmarshal(resp.Body(), &repos)
	if err != nil {
		return nil, errors.CannotUmarshalError.Wrap(err, "failed to umarshal github response")
	}

	res := make(map[string]TemplateEntry)

	for _, repo := range repos {
		if name, ok := repo["name"].(string); ok {
			if strings.HasPrefix(name, "validator-template-") {
				lang, _ := strings.CutPrefix(name, "validator-template-")
				res[lang] = TemplateEntry{Repo: repo["html_url"].(string), Branch: "main", Commit: "*"}
			}
		} else {
			return nil, errors.InvalidGithubResponseError.New("no field 'name' of type string")
		}
	}
	return res, nil
}

// loadLockfile loads the lockfile from disk
func loadLockfile() (Lockfile, error) {
	if !util.Exists(files.PkgLockFile()) {
		defaultLockfile := Lockfile{
			LastUpdate: util.ApisenseTime(time.Now().UTC()),
			Repos: map[string]RepoEntry{
				"buonotti": {
					Type: GitHub,
					Url:  BUONOTTI_ENDPOINT,
				},
			},
		}
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

// Update updates the lockfile with the templates found in the repos
func Update(override bool) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	for repoName, repo := range lockfile.Repos {
		log.PkgLogger().Info("Updating repo", "repo", repoName)
		remoteTemplates, err := DiscoverTemplates(repo)
		if err != nil {
			return err
		}
		for lang, temp := range remoteTemplates {
			if _, ok := lockfile.Templates[lang]; !ok || override {
				if lockfile.Templates == nil {
					lockfile.Templates = make(map[string]TemplateEntry)
				}
				log.PkgLogger().Info("Discovered template", "repo", repoName, "lang", lang, "url", temp.Repo)
				lockfile.Templates[lang] = temp
			}
		}
	}
	lockfile.LastUpdate = util.ApisenseTime(time.Now().UTC())
	return saveLockfile(lockfile)
}

// Create creates a new validator from the given template
func Create(lang string, name string, force bool, noCache bool) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	temp, ok := lockfile.Templates[lang]
	if !ok {
		return errors.LanguageNotAvailableError.New("language '%s' is not installed. run 'apisense templates update' to update local templates", lang)
	}
	cachePath := filepath.FromSlash(directories.ValidatorsCacheDirectory() + "/" + lang)
	destPath := filepath.FromSlash(directories.ValidatorCustomDirectory() + "/" + name)
	if !util.Exists(cachePath) {
		out, err := git.Clone(clone.Repository(temp.Repo), clone.Branch(temp.Branch), clone.Directory(cachePath))
		if err != nil {
			return errors.CannotCloneRepoError.Wrap(err, out)
		}
		if temp.Commit != "*" {
			err = os.Chdir(cachePath)
			if err != nil {
				return errors.CannotChangeDirectoryError.WrapWithNoMessage(err)
			}
			out, err = git.Checkout(checkout.Branch(temp.Commit))
			if err != nil {
				return errors.CannotCheckoutError.Wrap(err, out)
			}
		}
	}
	if util.Exists(destPath) && !force {
		return errors.DirectoryExistsError.New("destination directory already exists: %s", destPath)
	}
	err = gorecurcopy.CopyDirectory(cachePath, destPath)
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

	log.PkgLogger().Info("Created new validator from template", "lang", lang, "path", destPath)

	if noCache {
		err = os.RemoveAll(cachePath)
		if err != nil {
			return errors.CannotRemoveFileError.WrapWithNoMessage(err)
		}
	}

	return nil
}

func UpgradeAll() error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	for name := range lockfile.Templates {
		err = Upgrade(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func Upgrade(name string) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	if temp, ok := lockfile.Templates[name]; ok {
		if temp.Commit != "*" {
			log.PkgLogger().Warn("Template is locked to commit. Ignoring", "name", name, "commit", temp.Commit)
		} else {
			templDir := filepath.FromSlash(directories.ValidatorsCacheDirectory() + "/" + name)
			if !util.Exists(templDir) {
				log.PkgLogger().Warn("Template is not cached locally. Cannot upgrade", "name", name)
				return nil
			}
			err = os.Chdir(templDir)
			if err != nil {
				return errors.CannotChangeDirectoryError.WrapWithNoMessage(err)
			}
			out, err := git.Fetch()
			if err != nil {
				return errors.CannotFetchRepoError.Wrap(err, out)
			}
			out, err = git.Pull()
			if err != nil {
				return errors.CannotPullRepoError.Wrap(err, out)
			}
			log.PkgLogger().Info("Upgraded template", "name", name)
		}
	} else {
		return errors.TemplateNotFoundError.New("cannot find template with name: %s", name)
	}
	return nil
}

// AddRepo adds a new remote repo
func AddRepo(name string, url string, ty RepoType) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	_, ok := lockfile.Repos[name]
	if ok {
		return errors.RepoExistsError.New("repo %s already exists", name)
	}
	lockfile.Repos[name] = RepoEntry{
		Url:  url,
		Type: ty,
	}
	return saveLockfile(lockfile)
}

// AddTemplateSource adds a language with a custom git repo to the lockfile
func AddTemplateSource(lang string, url string, branch string, commit string) error {
	lockfile, err := loadLockfile()
	if err != nil {
		return err
	}
	_, ok := lockfile.Templates[lang]
	if ok {
		return errors.LanguageAlreadyExistsError.New("language '%s' already has a template", lang)
	}
	if lockfile.Templates == nil {
		lockfile.Templates = make(map[string]TemplateEntry)
	}
	lockfile.Templates[lang] = TemplateEntry{Repo: url, Branch: branch, Commit: commit}
	return saveLockfile(lockfile)
}
