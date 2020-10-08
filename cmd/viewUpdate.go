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

	"github.com/awmottaz/supd/internal/agenda"
	"github.com/spf13/cobra"
)

var lastN int

// viewUpdateCmd represents the viewUpdate command
var viewUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "View your update",
	Long:  `View your update`,
	Run: func(cmd *cobra.Command, args []string) {
		a, err := agenda.LoadFile(updatesFile)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		today := agenda.Today()

		upd := (*a)[today]

		cmd.Printf("Update for today (%v)\n\n%v\n", today, upd)
	},
}

func init() {
	viewCmd.AddCommand(viewUpdateCmd)

	viewUpdateCmd.Flags().IntVar(&lastN, "last", -1, "Shows this many recent updates. If set to a negative value, all updates are shown.")
}
