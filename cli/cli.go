package cli

import (
	"log"
	"os"
	"path"
	"time"

	flag "github.com/spf13/pflag"
)

// Update represents an update for a given day.
type Update struct {
	Date      time.Time `json:"date"`
	Plan      string    `json:"plan"`
	Completed []string  `json:"completed"`
	Notes     []string  `json:"notes"`
}

type appEnv struct {
	updateFile string
}

func (a *appEnv) fromArgs(args []string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	flag.StringVarP(&a.updateFile, "updateFile", "f", path.Join(home, "supd.json"), "the location of the updates JSON file")
	flag.Parse()
	return nil
}

// Run runs the supd command line app and returns its exit status.
func Run(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	log.Printf("Update file: %s\n", app.updateFile)
	return 0
}
