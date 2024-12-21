package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
	TopP        float64   `json:"top_p"`
}

type Response struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func GenerateCommitMessage(stagedChanges []string, githubToken string) (string, error) {
	messages := []Message{
		{
			Role:    "system",
			Content: "",
		},
		{
			Role: "user",
			Content: `
Create well formed git commit message based of off the currently staged file
contents. The message should be formatted as follows: <type>(<scope>): <subject>

Type (The type of commit being made): chore, docs, feat, fix, refactor, test, etc
Scope (The scope of the change): core, cli, api, auth, etc 
Subject (A very short description of the change): 

Only include changes to source files for the programming languages, shell configurations files, documentation such as readme and other .mds, and any changes to package management file. Exclude
any lock or sum files.

Do not use markdown format for the output.

Do not add more lines explaining the changes for each file that the single message formatted like <type>(<scope>): <subject>.

If there are no changes abort.
` + formatChanges(stagedChanges),
		},
	}

	requestBody := RequestBody{
		Messages:    messages,
		Model:       "gpt-4o",
		Temperature: 1,
		MaxTokens:   4096,
		TopP:        1,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://models.inference.ai.azure.com/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+githubToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", nil
}

func formatChanges(changes []string) string {
	return "[" + strings.Join(changes, ", ") + "]"
}
