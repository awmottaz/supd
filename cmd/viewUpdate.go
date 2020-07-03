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
	"strings"

	"github.com/awmottaz/supd/internal/update"
	"github.com/spf13/cobra"
)

// viewUpdateCmd represents the viewUpdate command
var viewUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "View your update",
	Long:  `View your update`,
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
		}

		prevUpd, err := collection.FindPrev(update.Today())
		if err != nil && err != update.NotFound {
			cmd.PrintErrln("failed to find the previous update:", err)
			os.Exit(1)
		}

		var out strings.Builder

		if prevUpd.Date.Day == 0 {
			out.WriteString("No previous update")
		} else if prevUpd.Done == nil {
			out.WriteString(fmt.Sprintf("No completed tasks recorded for %s", prevUpd.Date))
		} else {
			out.WriteString(fmt.Sprintf("DID on %s:\n%s", prevUpd.Date, prevUpd.Done.PrefixedString("  ")))
		}

		out.WriteString(fmt.Sprintf("\n\nPLAN for today, %s:\n  %v", upd.Date, upd.Plan))

		cmd.Println(out.String())
	},
}

func init() {
	viewCmd.AddCommand(viewUpdateCmd)
}
