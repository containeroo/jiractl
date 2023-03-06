package main

import (
	"github.com/containeroo/jiractl/cmd"
	_ "github.com/containeroo/jiractl/cmd/blockers"
	_ "github.com/containeroo/jiractl/cmd/issues"
	_ "github.com/containeroo/jiractl/cmd/worklogs"
)

func main() {
	cmd.Execute()
}
