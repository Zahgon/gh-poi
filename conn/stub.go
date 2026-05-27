package conn

import (
	"github.com/seachicken/gh-poi/mocks"
	"go.uber.org/mock/gomock"
)

type (
	Stub struct {
		Conn *mocks.MockConnection
		T    gomock.TestHelper
	}

	Times struct {
		N int
	}

	Conf struct {
		Times *Times
	}

	RepoNamesStub struct {
		RepoName string
		Filename string
	}

	RemoteHeadStub struct {
		BranchName string
		Filename   string
	}

	UpstreamOidStub struct {
		BranchName string
		Filename   string
	}

	LsRemoteHeadStub struct {
		BranchName string
		Filename   string
	}

	AssociatedBranchNamesStub struct {
		Oid      string
		Filename string
	}

	UncommittedChangeStub struct {
		Path   string
		Output string
	}

	LogStub struct {
		BranchName string
		Filename   string
	}

	WorktreeStub struct {
		Filename string
	}

	ConfigStub struct {
		Key      string
		Filename string
	}
)

var (
	fixturePath = "fixtures"
)

func Setup(ctrl *gomock.Controller) *Stub { _ = "STUB: not implemented"; return nil }

func NewConf(times *Times) *Conf { _ = "STUB: not implemented"; return nil }

func (s *Stub) GetRemoteNames(filename string, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetSshConfig(filename string, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetRepoNames(stubs []RepoNamesStub, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetBranchNames(filename string, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetMergedBranchNames(filename string, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetAssociatedRefNames(stubs []AssociatedBranchNamesStub, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetLog(stubs []LogStub, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetPullRequests(filename string, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetUncommittedChanges(stubs []UncommittedChangeStub, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) GetConfig(stubs []ConfigStub, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) CheckoutBranch(err error, conf *Conf) *Stub { _ = "STUB: not implemented"; return nil }

func (s *Stub) DeleteBranches(err error, conf *Conf) *Stub { _ = "STUB: not implemented"; return nil }

func (s *Stub) GetWorktrees(filename string, err error, conf *Conf) *Stub {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stub) RemoveWorktree(err error, conf *Conf) *Stub { _ = "STUB: not implemented"; return nil }

func configure(call *gomock.Call, conf *Conf) { _ = "STUB: not implemented"; return }

func (s *Stub) ReadFile(command string, category string, name string) string {
	_ = "STUB: not implemented"
	return ""
}
