package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/seachicken/gh-poi/shared"
)

var (
	bold    = color.New(color.Bold).SprintFunc()
	hiBlack = color.New(color.FgHiBlack).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
)

type StateFlag string

const (
	Closed StateFlag = "closed"
	Merged StateFlag = "merged"
)

func (s *StateFlag) String() string { _ = "STUB: not implemented"; return "" }

func (s *StateFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (s StateFlag) toModel() shared.PullRequestState {
	_ = "STUB: not implemented"
	return *new(shared.PullRequestState)
}

type ScanFlag string

const (
	Quick ScanFlag = "quick"
	Deep  ScanFlag = "deep"
)

func (s *ScanFlag) String() string { _ = "STUB: not implemented"; return "" }

func (s *ScanFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (s ScanFlag) toModel() shared.ScanMode {
	_ = "STUB: not implemented"
	return *new(shared.ScanMode)
}

func main() {
	state := Merged
	scan := Quick
	var dryRun bool
	var debug bool
	flag.Var(&state, "state", "Specify the PR state to delete by {closed|merged}")
	flag.Var(&scan, "scan", "Specify the scan mode by {quick|deep}")
	flag.BoolVar(&dryRun, "dry-run", false, "Show branches to delete without actually deleting it")
	flag.BoolVar(&debug, "debug", false, "Enable debug logs")
	flag.Usage = func() {
		fmt.Fprintf(color.Output, "%s\n\n", "Delete the merged local branches.")
		fmt.Fprintf(color.Output, "%s\n", bold("USAGE"))
		fmt.Fprintf(color.Output, "  %s\n\n", "gh poi <command> [flags]")
		fmt.Fprintf(color.Output, "%s", bold("COMMANDS"))
		fmt.Fprintf(color.Output, "%s\n", `
  lock:      Lock branches to prevent them from being deleted
  unlock:    Unlock branches to allow them to be deleted
  protect:   (Deprecated) use 'lock' instead
  unprotect: (Deprecated) use 'unlock' instead
  `)
		fmt.Fprintf(color.Output, "%s\n", bold("FLAGS"))
		maxLen := 0
		flag.VisitAll(func(f *flag.Flag) {
			if len(f.Name) > maxLen {
				maxLen = len(f.Name)
			}
		})
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(color.Output, "  --%-*s %s\n", maxLen+2, f.Name, f.Usage)
		})
		fmt.Println()
	}
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		runMain(state, scan, dryRun, debug)
	} else {
		subcmd, args := args[0], args[1:]
		switch subcmd {
		case "lock", "protect":
			lockCmd := flag.NewFlagSet("lock", flag.ExitOnError)
			lockCmd.Usage = func() {
				fmt.Fprintf(color.Output, "%s\n\n", "Lock branches to prevent them from being deleted")
				fmt.Fprintf(color.Output, "%s\n", bold("USAGE"))
				fmt.Fprintf(color.Output, "  %s\n\n", "gh poi lock <branchname>...")
			}
			lockCmd.Parse(args)

			// TODO: Remove after deprecated commands are removed
			if subcmd == "protect" {
				fmt.Fprintln(os.Stderr, "warning: 'protect' is deprecated, please use 'lock' instead")
			}
			runLock(args, debug)
		case "unlock", "unprotect":
			unlockCmd := flag.NewFlagSet("unlock", flag.ExitOnError)
			unlockCmd.Usage = func() {
				fmt.Fprintf(color.Output, "%s\n\n", "Unlock branches to allow them to be deleted")
				fmt.Fprintf(color.Output, "%s\n", bold("USAGE"))
				fmt.Fprintf(color.Output, "  %s\n\n", "gh poi unlock <branchname>...")
			}
			unlockCmd.Parse(args)

			// TODO: Remove after deprecated commands are removed
			if subcmd == "unprotect" {
				fmt.Fprintln(os.Stderr, "warning: 'unprotect' is deprecated, please use 'unlock' instead")
			}
			runUnlock(args, debug)
		default:
			fmt.Fprintf(os.Stderr, "unknown command %q for poi\n", subcmd)
		}
	}
}

func runMain(state StateFlag, scan ScanFlag, dryRun bool, debug bool) {
	_ = "STUB: not implemented"
	return
}

func runLock(branchNames []string, debug bool) { _ = "STUB: not implemented"; return }

func runUnlock(branchNames []string, debug bool) { _ = "STUB: not implemented"; return }

func printBranches(branches []shared.Branch) { _ = "STUB: not implemented"; return }

// Show worktree info for any branch with an associated worktree

func getIssueNoColor(state shared.PullRequestState, isDraft bool) color.Attribute {
	_ = "STUB: not implemented"
	return *new(color.Attribute)
}

func getBranches(branches []shared.Branch, states []shared.BranchState) []shared.Branch {
	_ = "STUB: not implemented"
	return nil
}
