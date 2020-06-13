package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
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
	case "edit":
		os.Exit(edit())
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
	supd <command>

Options:

	-h    display these help instructions

Commands:

	edit    open the updates file for editing`)
}

func getUpdatesFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "supd.json"), nil
}

func edit() int {
	file, err := getUpdatesFile()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	editorEx, err := exec.LookPath(editor)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	cmd := exec.Command(editorEx, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	return 0
}
