package conn

import (
	"context"
	"regexp"

	"github.com/seachicken/gh-poi/shared"
)

type (
	Connection struct {
		Debug bool
	}

	DebugMask int
)

const (
	None DebugMask = iota
	Output
)

var (
	hasSchemePattern  = regexp.MustCompile("^[^:]+://")
	scpLikeURLPattern = regexp.MustCompile("^([^@]+@)?([^:]+):(/?.+)$")
)

func GetRemoteNames(ctx context.Context, conn shared.Connection) ([]shared.Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (conn *Connection) GetRemoteNames(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// acceptable url formats:
//
//	ssh://[user@]host.xz[:port]/path/to/repo.git/
//	git://host.xz[:port]/path/to/repo.git/
//	http[s]://host.xz[:port]/path/to/repo.git/
//	ftp[s]://host.xz[:port]/path/to/repo.git/
//
// An alternative scp-like syntax may also be used with the ssh protocol:
//
//	[user@]host.xz:path/to/repo.git/
//
// ref. http://git-scm.com/docs/git-fetch#_git_urls
// the code is heavily inspired by https://github.com/x-motemen/ghq/blob/7163e61e2309a039241ad40b4a25bea35671ea6f/url.go
func parseRemotes(output string) []shared.Remote { _ = "STUB: not implemented"; return nil }

func (conn *Connection) GetSshConfig(ctx context.Context, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetRepoNames(ctx context.Context, hostname string, repoName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetBranchNames(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetMergedBranchNames(ctx context.Context, remoteName string, branchName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetRemoteHeadOid(ctx context.Context, remoteName string, branchName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetUpstreamOid(ctx context.Context, branchName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetLog(ctx context.Context, branchName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) GetAssociatedRefNames(ctx context.Context, oid string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// limitations:
// - https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests#search-within-a-users-or-organizations-repositories
// - https://docs.github.com/en/graphql/overview/resource-limitations
func (conn *Connection) GetPullRequests(
	ctx context.Context,
	hostname string, orgs string, repos string, queryHashes string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetUncommittedChanges(ctx context.Context, conn shared.Connection, opts ...string) ([]shared.UncommittedChange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (conn *Connection) GetUncommittedChanges(ctx context.Context, opts ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseUncommittedChanges(output string) []shared.UncommittedChange {
	_ = "STUB: not implemented"
	return nil
}

func (conn *Connection) GetConfig(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) AddConfig(ctx context.Context, key string, value string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) RemoveConfig(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) CheckoutBranch(ctx context.Context, branchName string, detach bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) DeleteBranches(ctx context.Context, branchNames []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) PruneRemoteBranches(ctx context.Context, remoteName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetWorktrees(ctx context.Context, conn shared.Connection) ([]shared.Worktree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (conn *Connection) GetWorktrees(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseWorktrees(output string) []shared.Worktree { _ = "STUB: not implemented"; return nil }

func (conn *Connection) RemoveWorktree(ctx context.Context, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (conn *Connection) run(ctx context.Context, name string, args []string, mask DebugMask) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func splitLines(text string) []string { _ = "STUB: not implemented"; return nil }
