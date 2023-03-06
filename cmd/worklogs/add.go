package worklogs

import (
	"context"
	"fmt"
	"time"

	"github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/containeroo/jiractl/internal"
	"github.com/spf13/cobra"
)

var worklogAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a worklog",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := internal.NewJiraClient()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		location, _ := time.LoadLocation(cmd.Flag("timezone").Value.String())
		startDate := cmd.Flag("date").Value.String()
		startTime := cmd.Flag("start-time").Value.String()
		if startDate == "today" {
			startDate = time.Now().In(location).Format("2006-01-02")
		}
		startDateTime := fmt.Sprintf("%s %s", startDate, startTime)
		parsedStartTime, err := time.ParseInLocation("2006-01-02 15:04", startDateTime, location)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		newTime := onpremise.Time(parsedStartTime)

		worklogRecord := onpremise.WorklogRecord{
			Comment:   cmd.Flag("comment").Value.String(),
			Started:   &newTime,
			TimeSpent: cmd.Flag("spent").Value.String(),
		}
		_, _, err = jiraClient.Issue.AddWorklogRecord(context.Background(), cmd.Flag("issue").Value.String(), &worklogRecord)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Worklog added successfully")
	},
}

func init() {
	worklogsRootCmd.AddCommand(worklogAddCmd)
	worklogAddCmd.PersistentFlags().String("issue", "", "Issue where to book time to")
	worklogAddCmd.PersistentFlags().String("date", "today", "When to add the worklog e.g. '2023-01-20'")
	worklogAddCmd.PersistentFlags().String("start-time", "", "When work was started e.g. '13:30'")
	worklogAddCmd.PersistentFlags().String("spent", "", "How much time was spent e.g. '2h 15m'")
	worklogAddCmd.PersistentFlags().String("comment", "", "Comment to the worklog")
	worklogAddCmd.PersistentFlags().String("timezone", "Europe/Zurich", "Timezone for the worklog")
	_ = worklogAddCmd.MarkPersistentFlagRequired("issue")
	_ = worklogAddCmd.MarkPersistentFlagRequired("start-time")
	_ = worklogAddCmd.MarkPersistentFlagRequired("spent")
	_ = worklogAddCmd.MarkPersistentFlagRequired("comment")
}
