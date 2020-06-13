package cli

import (
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

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
