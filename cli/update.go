package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Update represents an update for a given day.
type Update struct {
	Date      string   `json:"date"`
	Plan      string   `json:"plan"`
	Completed []string `json:"completed"`
	Notes     []string `json:"notes"`
}

// New creates a new Update for today.
func New() *Update {
	return &Update{
		Date: today(),
	}
}

func (upd Update) String() string {
	return fmt.Sprintf("* %s\n\tPLAN\n\t%s\n\tCOMPLETED\n\t%s", upd.Date, upd.Plan, upd.Completed)
}

// LoadToday loads the data from today's update into upd. If it does not exist, then upd will be a
// new Update for today.
func (upd *Update) LoadToday(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	updates := make([]Update, 0)
	err = json.Unmarshal(b, &updates)
	if err != nil {
		return err
	}

	var ok bool
	*upd, ok = find(updates, today())
	if !ok {
		upd = New()
		return nil
	}

	return nil
}

func find(updates []Update, date string) (Update, bool) {
	for _, u := range updates {
		if u.Date == date {
			return u, true
		}
	}

	return Update{}, false
}

// func (upd *Update) Write(w io.ReadWriter) error {

// }
