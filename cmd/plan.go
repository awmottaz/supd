package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	supd plan [STRING]

Summary:
	Use the plan command to set or view your plan for today.

	If called with no arguments, your plan for today will be printed to the
	console. If there is no plan set for today, the string "<none>" will be
	printed.

	If called with STRING, your plan for today will be set to that value. If
	there is already a plan set for today, you will be prompted to overwite
	it.`)
}

// Run executes the plan command and returns the exit code.
func (plan *PlanCmd) Run(args []string) int {
	filename, err := update.GetUpdatesFile()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	collection := &update.Collection{}

	err = collection.LoadFrom(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	switch len(args) {
	// no args, print today's plan
	case 0:
		u, err := collection.FindByDate(update.Today())
		if err == update.NotFound {
			fmt.Println("<none>")
			return 0
		} else if err != nil {
			fmt.Printf("getting today's plan: %s\n", err)
			return 1
		}
		fmt.Println(u.Plan)
		return 0

	// one arg, write it to today's plan
	case 1:
		u, err := collection.FindByDate(update.Today())
		if err != nil && err != update.NotFound {
			fmt.Println("getting today's plan:", err)
			return 1
		}

		// If an update already exists for today, get user's confirmation
		// before overwriting it with a new plan.
		if err != update.NotFound && !proceed(fmt.Sprintf("Plan for today already set to:\n\"%s\"\nOverwrite?", u.Plan)) {
			fmt.Println("ignoring new plan")
			return 0
		}

		collection.Add(update.Update{Date: update.Today(), Plan: args[0]})
		err = collection.Commit(filename)
		if err != nil {
			fmt.Println("commiting updates:", err)
			return 1
		}

		fmt.Println("plan written to", filename)
		return 0

	default:
		fmt.Println("error: too many arguments")
		plan.Usage()
		return 1
	}
}

// proceed asks for user confirmation before proceeding. Anything other than
// explicit approval results in disapproval. The approval status is returned.
func proceed(q string) bool {
	fmt.Printf("%s (y/N) ", q)

	scanner := bufio.NewReader(os.Stdin)
	out, err := scanner.ReadString('\n')
	if err != nil {
		return false
	}

	out = strings.Replace(out, "\n", "", -1)

	switch strings.ToLower(out) {
	case "y", "yes":
		return true
	default:
		return false
	}
}
