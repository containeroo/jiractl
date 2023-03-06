package issues

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/containeroo/jiractl/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var issueGetCmd = &cobra.Command{
	Use:     "get [ISSUE-KEY]",
	Aliases: []string{"g"},
	Short:   "Get an issue",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := internal.NewJiraClient()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		issue, _, _ := jiraClient.Issue.Get(context.Background(), args[0], nil)

		if issue.Fields.Assignee == nil {
			issue.Fields.Assignee = &onpremise.User{DisplayName: "Unassigned"}
		}

		blue := color.New(color.FgBlue).Add(color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Printf("%s\t%s\n", blue("Summary:"), yellow(issue.Fields.Summary))
		fmt.Printf("%s\t\t%s\n", blue("Status:"), yellow(issue.Fields.Status.Name))
		fmt.Printf("%s\t%s\n", blue("Assignee:"), yellow(issue.Fields.Assignee.DisplayName))
		fmt.Printf("%s\t\t%s\n", blue("Link:"), yellow(fmt.Sprintf("%sbrowse/%s", jiraClient.BaseURL, issue.Key)))
		fmt.Printf("%s\n%s\n", blue("Description:"), issue.Fields.Description)
	},
}

func init() {
	issuesRootCmd.AddCommand(issueGetCmd)
}
