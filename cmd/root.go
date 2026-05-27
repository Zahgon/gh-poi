package cmd

import (
	"context"
	"errors"

	"github.com/seachicken/gh-poi/shared"
)

const (
	github    = "github.com"
	localhost = "github.localhost"
)

var ErrNotFound = errors.New("not found")

// Returns a list of remotes prioritized for PR discovery.
// Both modes prioritize "origin," and when searching for pull requests,
// the parent (a.k.a upstream) repository is also included in the search.
//
// quick:
//   - Focuses on the most likely PR sources to minimize API calls.
//   - Returns only "origin" and the remote configured via `gh repo set-default`.
//
// deep:
//   - Scans all registered remotes to ensure comprehensive PR discovery.
//   - Useful for complex setups where PRs may span multiple forks or parents.
func GetPreferredRemotes(ctx context.Context, connection shared.Connection, scan shared.ScanMode) ([]shared.Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://github.com/cli/cli/blob/8f28d1f9d5b112b222f96eb793682ff0b5a7927d/internal/ghinstance/host.go#L26
func normalizeHostname(host string) string { _ = "STUB: not implemented"; return "" }

func findHostname(params []string, defaultName string) string { _ = "STUB: not implemented"; return "" }

func GetBranches(ctx context.Context, remotes []shared.Remote, connection shared.Connection, state shared.PullRequestState, scan shared.ScanMode, dryRun bool) ([]shared.
	Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadBranches(ctx context.Context, remote shared.Remote, defaultBranchName string, repoNames []string, connection shared.Connection, scan shared.ScanMode) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractMergedBranchNames(mergedNames []string) []string { _ = "STUB: not implemented"; return nil }

func applyDefault(branches []shared.Branch, defaultBranchName string) []shared.Branch {
	_ = "STUB: not implemented"
	return nil
}

func applyMerged(branches []shared.Branch, mergedNames []string) []shared.Branch {
	_ = "STUB: not implemented"
	return nil
}

func applyLocked(ctx context.Context, branches []shared.Branch, connection shared.Connection) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Remove after deprecated commands are removed

func applyCommits(ctx context.Context, branches []shared.Branch, defaultBranchName string, connection shared.Connection, scan shared.ScanMode) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyTrackedChanges(ctx context.Context, branches []shared.Branch, connection shared.Connection) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyWorktrees(ctx context.Context, branches []shared.Branch, connection shared.Connection) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Worktrees might not be supported or available, continue gracefully

// Create a map for quick branch-to-worktree lookup

func trimBranch(ctx context.Context, oids []string, branch shared.Branch, defaultBranchName string, connection shared.Connection) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractBranchNames(refNames []string) []string { _ = "STUB: not implemented"; return nil }

func applyPullRequest(ctx context.Context, branches []shared.Branch, prs []shared.PullRequest, connection shared.Connection) []shared.Branch {
	_ = "STUB: not implemented"
	return nil
}

func getPRNumber(mergeConfig string) int { _ = "STUB: not implemented"; return 0 }

func findMatchedPullRequest(branchName string, prs []shared.PullRequest, prNumbers map[string]int) []shared.PullRequest {
	_ = "STUB: not implemented"
	return nil
}

func checkDeletion(branches []shared.Branch, state shared.PullRequestState) []shared.Branch {
	_ = "STUB: not implemented"
	return nil
}

func getDeleteStatus(branch shared.Branch, state shared.PullRequestState) shared.BranchState {
	_ = "STUB: not implemented"
	return *new(shared.BranchState)
}

func isFullyMerged(branch shared.Branch, pr shared.PullRequest, state shared.PullRequestState) bool {
	_ = "STUB: not implemented"
	return false
}

// In the GitHub interface, closed status includes merged status, so we make it behave the same way.
// https://github.com/cli/cli/issues/8102

func switchToDefaultBranchIfDeleted(ctx context.Context, remotes []shared.Remote, branches []shared.Branch, defaultBranchName string, connection shared.Connection, dryRun bool) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToBranch(branchNames []string) []shared.Branch { _ = "STUB: not implemented"; return nil }

func getRepo(jsonResp string) ([]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func toPullRequests(jsonResp string) ([]shared.PullRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toPullRequestState(state string) (shared.PullRequestState, error) {
	_ = "STUB: not implemented"
	return *new(shared.PullRequestState), nil
}

func DeleteBranches(ctx context.Context, branches []shared.Branch, connection shared.Connection) ([]shared.Branch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBranchNames(branches []shared.Branch, state shared.BranchState) []string {
	_ = "STUB: not implemented"
	return nil
}

func deleteWorktrees(ctx context.Context, branches []shared.Branch, connection shared.Connection) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkDeleted(branchesBefore []shared.Branch, branchesAfter []shared.Branch) []shared.Branch {
	_ = "STUB: not implemented"
	return nil
}

func BranchNameExists(branchName string, branches []shared.Branch) bool {
	_ = "STUB: not implemented"
	return false
}

func SplitLines(text string) []string { _ = "STUB: not implemented"; return nil }
