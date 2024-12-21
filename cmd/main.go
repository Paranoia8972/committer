package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Paranoia8972/committer/internal"
)

func main() {
	changes, err := internal.FetchStagedChanges()
	if err != nil {
		log.Fatalf("Error fetching staged changes: %v", err)
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable is not set")
	}

	commitMessage, err := internal.GenerateCommitMessage(changes, githubToken)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	fmt.Println(commitMessage)
}
