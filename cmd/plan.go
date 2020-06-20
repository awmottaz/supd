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

	collection, err := update.LoadUpdates(filename)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	if len(args) > 1 {
		fmt.Println("error: too many arguments")
		plan.Usage()
		return 1
	}

	if len(args) == 1 {
		today := update.Today()

		tu, err := collection.FindByDate(today)
		if err != nil && err != update.NotFound {
			fmt.Println(err.Error())
			return 1
		}

		if err != update.NotFound {
			fmt.Printf("Today's plan already set to\n\"%s\"\n", tu.Plan)
			if !proceed("Overwrite?") {
				fmt.Println("ignoring new plan")
				return 0
			}
		}

		u := update.Update{Date: today, Plan: args[0]}

		collection.Add(u)
		err = collection.Commit(filename)
		if err != nil {
			fmt.Println(err.Error())
			return 1
		}
		fmt.Printf("plan written to %s\n", filename)
		return 0
	}

	return showTodayPlan(collection)
}

func showTodayPlan(c update.Collection) int {
	u, err := c.FindByDate(update.Today())
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
