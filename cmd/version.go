package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.9"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of jiractl",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
