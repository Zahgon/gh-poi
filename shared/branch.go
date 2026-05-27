package shared

import (
	"regexp"
)

type (
	BranchState int

	Branch struct {
		Head              bool
		Name              string
		IsDefault         bool
		IsMerged          bool
		IsLocked          bool
		HasTrackedChanges bool
		HasUntrackedFiles bool
		Commits           []string
		PullRequests      []PullRequest
		State             BranchState
		Worktree          *Worktree
	}

	UncommittedChange struct {
		X    string
		Y    string
		Path string
	}
)

const (
	Unknown BranchState = iota
	NotDeletable
	Deletable
	Deleted
)

var detachedBranchNameRegex = regexp.MustCompile(`^\(.+\)`)

func (b Branch) IsDetached() bool { _ = "STUB: not implemented"; return false }

func (uc *UncommittedChange) IsUntracked() bool { _ = "STUB: not implemented"; return false }
