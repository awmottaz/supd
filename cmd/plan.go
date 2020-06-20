package cmd

import (
	"fmt"

	"github.com/awmottaz/supd/update"
)

// PlanCmd is a Command for interacting with today's plan.
type PlanCmd struct{}

// Summary returns the summary for the plan command.
func (plan *PlanCmd) Summary() string {
	return "view your plan for today"
}

// Usage prints the instructions for using the plan command.
func (plan *PlanCmd) Usage() {
	fmt.Println(`Usage:

	supd plan

Summary:
	Use the plan command to see what your plan is for today. Prints
	"<none>" if you have not yet set your plan for today.`)
}

// Run executes the plan command and returns the exit code.
func (plan *PlanCmd) Run(args []string) int {
	filename, err := update.GetUpdatesFile()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	updates, err := update.LoadUpdates(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	u, err := update.FindByDate(updates, update.Today())
	if err == update.NotFound {
		fmt.Println("<none>")
		return 0
	} else if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	fmt.Println(u)
	return 0
}
