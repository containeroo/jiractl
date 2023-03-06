package issues

import (
	"github.com/containeroo/jiractl/cmd"
	"github.com/spf13/cobra"
)

var blockersRootCmd = &cobra.Command{
	Use:     "blockers",
	Aliases: []string{"b"},
	Short:   "Find blocker issues",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(blockersRootCmd)
}
