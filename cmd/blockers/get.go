package issues

import (
	"context"
	"fmt"

	"github.com/containeroo/jiractl/internal"
	"github.com/spf13/cobra"
)

var blockersGetCmd = &cobra.Command{
	Use:     "get [PROJECT-NAME]",
	Aliases: []string{"g"},
	Short:   "Get blocker issue for a project",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := internal.NewJiraClient()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		jql := fmt.Sprintf("project = %s AND type = Operations AND status = Running", args[0])
		issues, _, _ := jiraClient.Issue.Search(context.Background(), jql, nil)
		tbl := internal.NewBlockerGetTable()

		for _, issue := range issues {
			tbl.AddRow(issue.Key, issue.Fields.Summary)
		}

		tbl.Print()
	},
}

func init() {
	blockersRootCmd.AddCommand(blockersGetCmd)
}
