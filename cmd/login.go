package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/containeroo/jiractl/internal"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Jira",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(internal.ConfigDir); os.IsNotExist(err) {
			err := os.Mkdir(internal.ConfigDir, 0o755)
			if err != nil {
				fmt.Println("Could not create config directory:", err)
				return
			}
		}
		jiraLogin := internal.JiraLogin{
			Server: cmd.Flag("server").Value.String(),
			Token:  cmd.Flag("token").Value.String(),
		}

		fileContent, _ := json.Marshal(jiraLogin)
		_ = os.WriteFile(filepath.Join(internal.ConfigDir, "login.json"), fileContent, 0o644)

		jiraClient, _ := internal.NewJiraClient()
		user, _, _ := jiraClient.User.GetSelf(context.Background())
		if user == nil {
			fmt.Println("Could not login to Jira")
			return
		}

		fmt.Printf("Login to %s successful", jiraLogin.Server)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().String("server", "", "Jira server URL")
	loginCmd.PersistentFlags().String("token", "", "Personal access token")
	_ = loginCmd.MarkPersistentFlagRequired("server")
	_ = loginCmd.MarkPersistentFlagRequired("token")
}
