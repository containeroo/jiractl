package worklogs

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/containeroo/jiractl/internal"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type TableContent struct {
	IssueKey       string
	StartTime      string
	TimeSpent      string
	WorklogComment string
}

func getIssueWorklogs(jiraClient *onpremise.Client, issue string) []onpremise.WorklogRecord {
	workLogs, _, _ := jiraClient.Issue.GetWorklogs(context.Background(), issue)
	return workLogs.Worklogs
}

func extractWorklogs(issueKey string, worklogs []onpremise.WorklogRecord, email, date string, all bool) []TableContent {
	var tableContent []TableContent
	for _, worklog := range worklogs {
		if worklog.Author.EmailAddress != email {
			continue
		}
		startTime := time.Time(*worklog.Started).Format("2006-01-02 15:04")
		if !all && !strings.Contains(startTime, date) {
			continue
		}
		tableContent = append(tableContent, TableContent{
			IssueKey:       issueKey,
			StartTime:      startTime,
			TimeSpent:      worklog.TimeSpent,
			WorklogComment: worklog.Comment,
		})
	}
	return tableContent
}

var worklogsGetCmd = &cobra.Command{
	Use:     "get [ISSUE-KEY]",
	Aliases: []string{"g"},
	Short:   "Get worklogs",
	Args:    cobra.RangeArgs(0, 1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && cmd.Flag("all").Value.String() == "true" {
			fmt.Println("You can't use the --all flag without an issue key")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := internal.NewJiraClient()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var tbl table.Table
		user := internal.GetUser(jiraClient)

		dateFilter := cmd.Flag("date").Value.String()
		if dateFilter == "today" {
			dateFilter = time.Now().Format("2006-01-02")
		}

		switch len(args) {
		case 0:
			tbl = internal.NewWorklogGetMultiIssuesTable()
			jql := fmt.Sprintf("worklogAuthor = currentUser() AND worklogDate = '%s'", dateFilter)
			issues, _, _ := jiraClient.Issue.Search(context.Background(), jql, nil)
			var extractedWorklogs []TableContent
			for _, issue := range issues {
				workLogs := getIssueWorklogs(jiraClient, issue.Key)
				extractedWorklogs = append(extractedWorklogs, extractWorklogs(issue.Key, workLogs, user.EmailAddress, dateFilter, false)...)
			}

			sort.Slice(extractedWorklogs, func(i, j int) bool {
				return extractedWorklogs[i].StartTime < extractedWorklogs[j].StartTime
			})

			for _, workLog := range extractedWorklogs {
				tbl.AddRow(workLog.IssueKey, workLog.StartTime, workLog.TimeSpent, workLog.WorklogComment)
			}
		case 1:
			tbl = internal.NewWorklogGetSingleIssueTable()
			workLogs := getIssueWorklogs(jiraClient, args[0])
			all := false
			if cmd.Flag("all").Value.String() == "true" {
				all = true
			}
			extractedWorklogs := extractWorklogs(args[0], workLogs, user.EmailAddress, dateFilter, all)

			for _, workLog := range extractedWorklogs {
				tbl.AddRow(workLog.StartTime, workLog.TimeSpent, workLog.WorklogComment)
			}
		}

		tbl.Print()
	},
}

func init() {
	var all bool
	worklogsRootCmd.AddCommand(worklogsGetCmd)
	worklogsGetCmd.PersistentFlags().String("date", "today", "Filter worklogs by date e.g. '2023-02-02'")
	worklogsGetCmd.PersistentFlags().BoolVar(&all, "all", false, "Show all worklogs from an issue")
}
