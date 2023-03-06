package internal

import (
	"context"

	"github.com/andygrunwald/go-jira/v2/onpremise"
)

func GetUser(jiraClient *onpremise.Client) *onpremise.User {
	user, _, _ := jiraClient.User.GetSelf(context.Background())
	return user
}
