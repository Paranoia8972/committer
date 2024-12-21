package internal

import (
	"os/exec"
	"strings"
)

func FetchStagedChanges() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	changes := strings.Split(string(output), "\n")
	var stagedChanges []string
	for _, change := range changes {
		if change != "" {
			stagedChanges = append(stagedChanges, change)
		}
	}

	return stagedChanges, nil
}
