package main

// TODO: Use a Go package for git stuff, if posssible.

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

type git struct {
	Location string
	URL      string
	Branch   string
}

// checkout checks out a repo.
func (g git) checkout() error {
	cmd := exec.Command("git", "checkout", g.Branch)
	cmd.Dir = g.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git checkout %s: Something went wrong", g.Location)
	}

	return nil
}

// clean cleans a repo.
func (g git) clean() error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = g.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clean %s: Something went wrong", g.Location)
	}

	return nil
}

// clone clones a repo.
func (g git) clone() error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", g.Branch, g.URL,
		g.Location)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone %s: Something went wrong", g.URL)
	}

	return nil
}

// diff checks a repo for differences.
func (g git) diff() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "--diff-filter",
		"ACDMR", "origin/"+g.Branch)
	cmd.Dir = g.Location
	bb := new(bytes.Buffer)
	cmd.Stdout = bb

	if err := cmd.Run(); err != nil {
		return []string{}, fmt.Errorf("git diff %s: Something went wrong", g.
			Location)
	}

	d := bb.String()
	if len(d) < 1 {
		return []string{}, nil
	}

	// Make output pretty.
	// TODO: This prints Deleted when it should be Added.
	d = strings.Replace(d, "A\t", "Added ", -1)
	d = strings.Replace(d, "C\t", "Copied ", -1)
	d = strings.Replace(d, "D\t", "Deleted ", -1)
	d = strings.Replace(d, "M\t", "Modiefied ", -1)
	d = strings.Replace(d, "R\t", "Renamed ", -1)
	dl := strings.Split(d, "\n")
	sort.Strings(dl)

	return dl[1:], nil
}

// fetch fetches a repo.
func (g git) fetch() error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = g.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git fetch %s: Something went wrong", g.Location)
	}

	return nil
}

// reset resets a repo.
func (g git) reset() error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+g.Branch)
	cmd.Dir = g.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git reset %s: Something went wrong", g.Location)
	}

	return nil
}
