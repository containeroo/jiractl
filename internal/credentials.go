package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type JiraLogin struct {
	Server string `json:"server"`
	Token  string `json:"token"`
}

var ConfigDir = filepath.Join(os.Getenv("HOME"), "/.config/jiractl")

func getJiraCredentials() (JiraLogin, error) {
	jiraCredentials := JiraLogin{}
	file, err := os.ReadFile(filepath.Join(ConfigDir, "login.json"))
	if err != nil {
		return jiraCredentials, fmt.Errorf("Could not find Jira credentials file. Please run 'jiractl login' first.")
	}

	_ = json.Unmarshal([]byte(file), &jiraCredentials)

	return jiraCredentials, nil
}
