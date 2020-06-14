package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

// EditCmd is a Command that can open the updates file in an editor.
type EditCmd struct{}

// Summary provides a brief description of what edit does.
func (edit *EditCmd) Summary() string {
	return "open the updates file for editing"
}

// Usage prints instructions for using the edit command.
func (edit *EditCmd) Usage() {
	fmt.Println(`Usage:

	supd edit

Summary:

	Opens the updates file for editing. The editor used will be inferred
	from the EDITOR environment variable if it is set, falling back to vim.`)
}

// Run executes the edit command.
func (edit *EditCmd) Run(args []string) int {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	filepath := path.Join(home, "supd.json")

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	fmt.Printf("Opening %s with %s...\n", filepath, editor)
	editorEx, err := exec.LookPath(editor)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	cmd := exec.Command(editorEx, filepath)
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
