package internal

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func NewWorklogGetSingleIssueTable() table.Table {
	worklogGetTable := table.New("Date", "Time Spent", "Comment")
	colorTable(worklogGetTable)
	return worklogGetTable
}

func NewWorklogGetMultiIssuesTable() table.Table {
	worklogGetTable := table.New("Issue", "Date", "Time Spent", "Comment")
	colorTable(worklogGetTable)
	return worklogGetTable
}

func NewIssueListTable() table.Table {
	issueGetTable := table.New("Issue Key", "Summary", "Status", "Assignee")
	colorTable(issueGetTable)
	return issueGetTable
}

func NewBlockerGetTable() table.Table {
	blockerGetTable := table.New("Issue Key", "Summary")
	colorTable(blockerGetTable)
	return blockerGetTable
}

func colorTable(tbl table.Table) {
	headerFmt := color.New(color.FgHiBlue, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
}
