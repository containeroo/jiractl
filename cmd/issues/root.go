package issues

import (
	"github.com/containeroo/jiractl/cmd"
	"github.com/spf13/cobra"
)

var issuesRootCmd = &cobra.Command{
	Use:     "issues",
	Aliases: []string{"i"},
	Short:   "Issues related commands",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(issuesRootCmd)
}
