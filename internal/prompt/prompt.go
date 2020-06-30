// Copyright © 2020 Tony Mottaz <tony@mottaz.dev>

// Package prompt provides functions that prompt the user for input.
package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirm will ask the user to confirm by typing "y" or "n". Valid positive
// responses are
//
// 	y, Y, yes, Yes, YES
//
// Valid negative responses are
//
//	n, N, no, No, NO
//
// If no valid response is received, the user will be asked again. This function
// does not exit until a valid response is received. An error is returned in the
// case of an I/O issue.
func Confirm(question string, defaultVal bool) (bool, error) {
	var ynPrompt string

	if defaultVal {
		ynPrompt = "[Y/n]"
	} else {
		ynPrompt = "[y/N]"
	}

	fmt.Printf("%s %s ", question, ynPrompt)

	scanner := bufio.NewReader(os.Stdin)
	response, err := scanner.ReadString('\n')
	if err != nil {
		return false, err
	}

	response = strings.Replace(response, "\n", "", -1)

	switch strings.ToLower(response) {
	case "":
		return defaultVal, nil
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		fmt.Println("\n✘ Please type (y)es or (n)o and then press enter.")
		return Confirm(question, defaultVal)
	}
}
