package main

import (
	"fmt"
	"os"

	"github.com/awmottaz/supd/cmd"
)

// Version is the current version of the app.
var Version = "<dev>"

func main() {
	if len(os.Args) < 2 {
		summary()
		os.Exit(0)
	}

	edit := &cmd.EditCmd{}
	plan := &cmd.PlanCmd{}

	help := cmd.NewHelp()
	help.Register("edit", edit)
	help.Register("plan", plan)

	switch os.Args[1] {
	case "help":
		os.Exit(help.Run(os.Args[2:]))
	case "edit":
		os.Exit(edit.Run(os.Args[2:]))
	case "plan":
		os.Exit(plan.Run(os.Args[2:]))
	default:
		fmt.Printf("unknown command \"%s\"\n", os.Args[1])
		help.Usage()
		os.Exit(1)
	}
}

func summary() {
	fmt.Printf("supd - version %s\nRun \"supd help\" for usage instructions", Version)
}
