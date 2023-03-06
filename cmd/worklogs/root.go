package worklogs

import (
	"github.com/containeroo/jiractl/cmd"
	"github.com/spf13/cobra"
)

var worklogsRootCmd = &cobra.Command{
	Use:     "worklogs",
	Aliases: []string{"w"},
	Short:   "Worklogs related commands",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(worklogsRootCmd)
}
