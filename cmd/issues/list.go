package issues

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/containeroo/jiractl/internal"
	"github.com/spf13/cobra"
)

var issuesListCmd = &cobra.Command{
	Use:     "list [JQL-QUERY]",
	Aliases: []string{"l"},
	Short:   "List issues",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := internal.NewJiraClient()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		issues, _, _ := jiraClient.Issue.Search(context.Background(), args[0], nil)
		tbl := internal.NewIssueListTable()

		for _, issue := range issues {
			if issue.Fields.Assignee == nil {
				issue.Fields.Assignee = &onpremise.User{DisplayName: "Unassigned"}
			}
			tbl.AddRow(issue.Key, issue.Fields.Summary, issue.Fields.Status.Name, issue.Fields.Assignee.DisplayName)
		}

		tbl.Print()
	},
}

func init() {
	issuesRootCmd.AddCommand(issuesListCmd)
}
