package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// HelpCmd is a Command that provides help to the user.
type HelpCmd struct {
	commands map[string]Command
}

// NewHelp constructs a new HelpCmd that can register other Commands.
func NewHelp() *HelpCmd {
	return &HelpCmd{make(map[string]Command)}
}

// Usage prints the usage instructions for this help command.
func (help *HelpCmd) Usage() {
	fmt.Println(`Usage:

	supd [options]
	supd <command>

Options:

	-h    display these help instructions

Commands:`)
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	for name, cmd := range help.commands {
		fmt.Fprintf(w, "\t%s\t%s\n", name, cmd.Summary())
	}
	w.Flush()
	fmt.Println("\nUse \"supd help <command>\" for more information about a command.")
}

// Register adds cmd to the commands that help can help with.
func (help *HelpCmd) Register(name string, cmd Command) {
	help.commands[name] = cmd
}

// Run shows the usage instructions for the given command. If no command is given, the usage
// instructions for the entire app is shown.
func (help *HelpCmd) Run(args []string) int {
	if len(args) < 1 {
		help.Usage()
		return 0
	}

	if len(args) > 1 {
		fmt.Println("error: too many arguments")
		help.Usage()
		return 1
	}

	for name, cmd := range help.commands {
		if args[0] == name {
			cmd.Usage()
			return 0
		}
	}

	fmt.Printf("supd help %s: unknown command \"%s\"\n", args[0], args[0])
	help.Usage()
	return 1
}
