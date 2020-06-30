/*
Copyright © 2020 Tony Mottaz <tony@mottaz.dev>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/awmottaz/supd/internal/prompt"
	"github.com/awmottaz/supd/internal/update"
	"github.com/spf13/cobra"
)

var listPlan bool
var removePlan bool

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "set your plan for today",
	Long:  `Use 'plan' to set or view your plan for today.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !listPlan && !removePlan && len(args) != 1 {
			return fmt.Errorf("exactly 1 argument expected, received %d", len(args))
		}
		return nil
	},
	Run: runPlan,
}

func init() {
	rootCmd.AddCommand(planCmd)

	planCmd.Flags().BoolVarP(&listPlan, "list", "l", false, "list your plan for today")
	planCmd.Flags().BoolVarP(&removePlan, "remove", "r", false, "remove your plan for today")
}

func runPlan(cmd *cobra.Command, args []string) {
	filename, err := update.GetUpdatesFile()
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(1)
	}

	collection := &update.Collection{}

	err = collection.LoadFrom(filename)
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(1)
	}

	upd, err := collection.FindByDate(update.Today())
	if err != nil && err != update.NotFound {
		cmd.PrintErrln("failed to read today's plan:", err)
		os.Exit(1)
	}
	updFound := err != update.NotFound

	if listPlan {
		if !updFound {
			cmd.Println()
			return
		}

		cmd.Println(upd.Plan)
		return
	}

	if removePlan {
		cmd.Println("I cannot remove the plan, yet :(")
		return
	}

	// If an update already exists for today, get user's permission to overwrite.
	if updFound {
		conf, err := prompt.Confirm(fmt.Sprintf("Plan for today already set to:\n\"%s\"\nOverwrite?", upd.Plan), false)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		if !conf {
			os.Exit(1)
		}
	}

	collection.Add(update.Update{Date: update.Today(), Plan: args[0]})
	err = collection.Commit(filename)
	if err != nil {
		cmd.PrintErrln("failed to commit update:", err)
		os.Exit(1)
	}

	cmd.Println("plan written to", filename)
}
