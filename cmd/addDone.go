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
	"os"

	"github.com/awmottaz/supd/internal/update"
	"github.com/spf13/cobra"
)

// addDoneCmd represents the done command
var addDoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Add a completed task to your done list",
	Long:  `Add a completed task to your done list`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		} else if err == update.NotFound {
			upd = update.Update{Date: update.Today()}
		}

		if upd.Done == nil {
			upd.Done = update.DoneList{}
		}

		upd.Done = append(upd.Done, args[0])

		collection.Add(upd)

		err = collection.Commit(filename)
		if err != nil {
			cmd.PrintErrln("failed to commit update:", err)
			os.Exit(1)
		}

		cmd.Println("completed task saved for", update.Today())
	},
}

func init() {
	addCmd.AddCommand(addDoneCmd)
}
