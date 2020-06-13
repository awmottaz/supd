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

	upd := New()

	f, err := os.Open(app.updateFile)
	if err != nil {
		log.Println(err.Error())
		return 1
	}
	defer f.Close()

	err = upd.LoadToday(f)
	if err != nil {
		log.Println(err.Error())
		return 1
	}

	err = prompt(&upd.Plan, "Plan for today")
	if err != nil {
		log.Println(err.Error())
		return 1
	}

	fmt.Println(upd)

	return 0
}

func today() string {
	return time.Now().Format("2006-01-02")
}

// prompt shows the message to the user and saves their response to target.
func prompt(target *string, message string) error {
	reader := bufio.NewReader(os.Stdin)
	var err error

	fmt.Printf("%s> ", message)
	*target, err = reader.ReadString('\n')

	return err
}
