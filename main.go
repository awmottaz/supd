package main

import (
	"fmt"
	"os"
)

// Version is the current version of the app.
var Version = "<dev>"

func main() {
	if len(os.Args) < 2 {
		summary()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "-h":
		usage()
		os.Exit(0)
	case "help":
		usage()
		os.Exit(0)
	default:
		fmt.Printf("unknown command \"%s\"\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func summary() {
	fmt.Printf(`supd - version %s
Run "supd help" for usage instructions
`, Version)
}

func usage() {
	fmt.Println(`Usage:

	supd [options]

Options:

	-h    display these help instructions`)
}
