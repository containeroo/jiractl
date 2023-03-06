package internal

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func NewJiraClient() (*jira.Client, error) {
	jiraCredentials, err := getJiraCredentials()
	if err != nil {
		return nil, err
	}

	jiraTransport := jira.BearerAuthTransport{
		Token: jiraCredentials.Token,
	}

	jiraClient, err := jira.NewClient(jiraCredentials.Server, jiraTransport.Client())
	if err != nil {
		fmt.Printf("Error creating Jira client: %s", err)
		return nil, err
	}
	return jiraClient, nil
}
