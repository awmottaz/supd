package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// An Update is a record of a given day.
type Update struct {
	Plan     `json:"plan"`
	DoneList `json:"done"`
}

func (u Update) String() string {
	return fmt.Sprintf("PLAN:\n%v\n\nDONE:\n%v", u.Plan, u.DoneList)
}

// WithDate is an Update plus the Date.
type WithDate struct {
	Date
	Plan
	DoneList
}

// Agenda is a mapping from Dates to Updates.
type Agenda map[Date]Update

// WithDateList is a list of updates with their dates included.
type WithDateList []WithDate

// FromAgenda converts an Agenda to a List.
func FromAgenda(a Agenda) WithDateList {
	var l WithDateList
	for date, update := range a {
		l = append(l, WithDate{Plan: update.Plan, DoneList: update.DoneList, Date: date})
	}
	return l
}

// ByDateChronological implements sort.Sort on a WithDateList, sorting the entries
// chronologically by Date.
type ByDateChronological WithDateList

func (b ByDateChronological) Len() int {
	return len(b)
}

func (b ByDateChronological) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDateChronological) Less(i, j int) bool {
	d1, d2 := b[i].Date, b[j].Date
	return d1.Before(d2)
}

// LoadFrom loads the Agenda from a JSON file located at filename.
func (a *Agenda) LoadFrom(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, a)

	return err
}

// FindByDate finds the update whose date matches date, or nil if not found.
func (a Agenda) FindByDate(date Date) *Update {
	ul := FromAgenda(a)
	sort.Sort(ByDateChronological(ul))
	for _, u := range ul {
		if u.Date == date {
			return &Update{Plan: u.Plan, DoneList: u.DoneList}
		}
	}
	return nil
}

// FindPrev finds the most recent update whose date is before the given date,
// returning nil if none exists.
func (a Agenda) FindPrev(date Date) *WithDate {
	ul := FromAgenda(a)
	sort.Sort(ByDateChronological(ul))

	for _, update := range ul {
		if update.Date.Before(date) {
			return &update
		}
	}

	return nil
}

// FindLastN returns the most recent n updates from the collection, up to the
// entire collection.
func (a Agenda) FindLastN(n int) []WithDate {
	ul := FromAgenda(a)
	sort.Sort(ByDateChronological(ul))
	num := minInt(len(ul), n)
	return ul[len(ul)-num:]
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

// WriteTo writes the Agenda to filename as formatted JSON.
func (a *Agenda) WriteTo(filename string) error {
	data, err := json.Marshal(a)
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

// AddUpdate adds an update for a given date. If overwrite is true, then it will
// overwrite an existing update if present. Returns an error if an update for
// this date exists and overwrite is false.
func (a *Agenda) AddUpdate(date Date, update Update, overwrite bool) error {
	_, ok := (*a)[date]

	if !ok && !overwrite {
		return fmt.Errorf("cannot overwrite existing update for date %v", date)
	}

	(*a)[date] = update
	return nil
}
