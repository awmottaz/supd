package cli

import "fmt"

// Update represents an update for a given day.
type Update struct {
	Date      string   `json:"date"`
	Plan      string   `json:"plan"`
	Completed []string `json:"completed"`
	Notes     []string `json:"notes"`
}

func (upd Update) String() string {
	return fmt.Sprintf("* %s\n\tPLAN\n\t%s", upd.Date, upd.Plan)
}
