package update

import (
	"os"
	"path/filepath"
)

// An Update is a record of what you plan to do on a given day. The Date must be
// a string in yyyy-mm-dd format.
type Update struct {
	Date string `json:"date"`
	Plan string `json:"plan"`
}

// GetUpdatesFile returns the resolved absolute path to the user's updates file.
// If the SUPD_FILE environment variable is set, then this path will be used.
// Otherwise, this defaults to $HOME/supd.json.
func GetUpdatesFile() (string, error) {
	envPath := os.Getenv("SUPD_FILE")
	if envPath != "" {
		return envPath, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "supd.json"), nil
}
