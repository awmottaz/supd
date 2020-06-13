package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// Run runs the supd command line app and returns its exit status.
func Run(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}

	var upd Update

	upd.Date = time.Now().Format("2006-01-02")

	err = prompt(&upd.Plan, "Plan for today")
	if err != nil {
		log.Println(err.Error())
		return 1
	}

	fmt.Println(upd)

	return 0
}

// prompt shows the message to the user and saves their response to target.
func prompt(target *string, message string) error {
	reader := bufio.NewReader(os.Stdin)
	var err error

	fmt.Printf("%s> ", message)
	*target, err = reader.ReadString('\n')

	return err
}
