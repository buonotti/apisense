package errors

import "github.com/joomcode/errorx"

var (
	RepoErrors = errorx.NewNamespace("repo")
	GitErrors  = RepoErrors.NewSubNamespace("git")
)

var (
	InvalidRepoTypeError       = RepoErrors.NewType("invalid_type")
	GithubUnreachableError     = RepoErrors.NewType("github_unreachable")
	InvalidGithubResponseError = RepoErrors.NewType("invalid_response")
	RepoExistsError            = RepoErrors.NewType("repo_exists")
	CannotCloneRepoError       = GitErrors.NewType("cannot_clone")
	CannotFetchRepoError       = GitErrors.NewType("cannot_fetch")
	CannotPullRepoError        = GitErrors.NewType("cannot_pull")
	CannotInitRepoError        = GitErrors.NewType("cannot_init")
	CannotAddError             = GitErrors.NewType("cannot_add")
	CannotCommitError          = GitErrors.NewType("cannot_commit")
	CannotCheckoutError        = GitErrors.NewType("cannot_checkout")
	LanguageNotAvailableError  = RepoErrors.NewType("language_not_available")
	LanguageAlreadyExistsError = RepoErrors.NewType("language_already_exists")
)
