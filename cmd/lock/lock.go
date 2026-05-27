package lock

import (
	"context"

	"github.com/seachicken/gh-poi/shared"
)

func LockBranches(ctx context.Context, targetBranchNames []string, connection shared.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Remove after deprecated commands are removed

func UnlockBranches(ctx context.Context, targetBranchNames []string, connection shared.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Remove after deprecated commands are removed
