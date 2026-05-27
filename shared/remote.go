package shared

type Remote struct {
	Name       string
	Hostname   string
	RepoName   string
	GhResolved string
}

func (r Remote) ResolvedRepoName() string { _ = "STUB: not implemented"; return "" }
