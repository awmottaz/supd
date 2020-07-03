package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// A DoneList is a list of completed tasks.
type DoneList []string

func (dl DoneList) String() string {
	var out strings.Builder
	for i, did := range dl {
		fmt.Fprintf(&out, "%d: %s\n", i+1, did)
	}
	return strings.TrimSuffix(out.String(), "\n")
}

// An Update is a record of what you plan to do on a given day. The Date must be
// a string in yyyy-mm-dd format.
type Update struct {
	Date Date     `json:"date"`
	Plan string   `json:"plan"`
	Done DoneList `json:"done"`
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

// LoadFrom loads the collection of updates from a JSON file located at filename.
func (c *Collection) LoadFrom(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)

	return err
}

// FindByDate finds the first update whose date matches date. A NotFound error
// is returned if such an update is not present in the updateList.
func (c Collection) FindByDate(date Date) (Update, error) {
	sort.Sort(c)
	for _, update := range c {
		if update.Date == date {
			return update, nil
		}
	}
	return Update{}, NotFound
}

// Commit writes the collection to filename as formatted JSON, sorted by Date.
func (c Collection) Commit(filename string) error {
	sort.Sort(c)
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "\t")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = buf.WriteTo(f)
	if err != nil {
		return err
	}

	return nil
}

// Add adds the update to a collection. If the collection already includes an
// update for the same date as u, it is overwritten with the new update.
func (c *Collection) Add(update Update) {
	n := -1
	*c = append(*c, update)

	for i, u := range *c {
		if u.Date == update.Date {
			n = i
			break
		}
	}

	if n >= 0 && n+1 < len(*c) {
		*c = append((*c)[:n], (*c)[n+1:]...)
	}
}
