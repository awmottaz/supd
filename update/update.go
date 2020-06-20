package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// An Error describes known or expected errors that may arise while using
// this package.
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// NotFound error occurs when an update cannot be found.
	NotFound = Error("Update not found")
)

// A Date consists of a year, month, and day.
type Date struct {
	Year  int
	Month int
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

// MarshalJSON converts a date to a JSON string.
func (d Date) MarshalJSON() ([]byte, error) {
	s := fmt.Sprint(d)
	return []byte(s), nil
}

// UnmarshalJSON converts JSON data to a date.
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}

	year, month, day := t.Date()
	d.Year = year
	d.Month = int(month)
	d.Day = day

	return nil
}

func timeToDate(t time.Time) Date {
	var d Date

	year, month, day := t.Date()

	d.Year = year
	d.Month = int(month)
	d.Day = day

	return d
}

// Today returns today's date.
func Today() Date {
	now := time.Now().Local()
	return timeToDate(now)
}

// An Update is a record of what you plan to do on a given day. The Date must be
// a string in yyyy-mm-dd format.
type Update struct {
	Date Date   `json:"date"`
	Plan string `json:"plan"`
}

func (u Update) String() string {
	return fmt.Sprintf("* Update %s *\nPLAN\n    %s", u.Date, u.Plan)
}

// Collection is a list of updates. It implements sort.Interface for []Update
// based on the Date.
type Collection []Update

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Collection) Less(i, j int) bool {
	d1, d2 := c[i].Date, c[j].Date

	if d1.Year < d2.Year {
		return true
	}
	if d1.Month < d2.Month {
		return true
	}
	if d1.Day < d2.Day {
		return true
	}
	return false
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

// LoadUpdates parses the list of updates in filename.
func LoadUpdates(filename string) (Collection, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	update := make([]Update, 0)
	err = json.Unmarshal(data, &update)

	return update, err
}

// FindByDate finds the first update whose date matches date. A NotFound error
// is returned if such an update is not present in the updateList.
func FindByDate(collection Collection, date Date) (Update, error) {
	sort.Sort(collection)
	for _, update := range collection {
		if update.Date == date {
			return update, nil
		}
	}
	return Update{}, NotFound
}
