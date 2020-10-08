package agenda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// DateFormat is used to stringify and parse Date values to/from strings.
const DateFormat = "2006-01-02"

// Date specifies a given day with no time information.
type Date struct {
	Day   int
	Month time.Month
	Year  int
}

func (d Date) toTime() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 12, 0, 0, 0, time.Local)
}

func fromTime(t time.Time) Date {
	return Date{t.Day(), t.Month(), t.Year()}
}

func (d Date) String() string {
	return d.toTime().Format(DateFormat)
}

// MarshalJSON converts a Date to a JSON string.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.toTime().Format(DateFormat))
}

// UnmarshalJSON converts JSON data to a date.
func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("parsing date field: %w", err)
	}

	if s == "null" {
		return nil
	}

	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return fmt.Errorf("parsing date field: %w", err)
	}

	*d = fromTime(t)
	return nil
}

// Today returns a Date from the current local time.
func Today() Date {
	return fromTime(time.Now())
}

// A Plan is what you intend to accomplish.
type Plan string

// DoneList is a list of tasks you have completed.
type DoneList []string

// An Update describes what happened in a day.
type Update struct {
	Plan     `json:"plan"`
	DoneList `json:"done"`
}

// An Agenda is a schedule of updates.
type Agenda map[Date]Update

// MarshalJSON encodes an Agenda into JSON.
func (a Agenda) MarshalJSON() ([]byte, error) {
	s := make(map[string]Update)
	for date, update := range a {
		s[date.String()] = update
	}
	return json.Marshal(s)
}

// UnmarshalJSON decodes a JSON object into an Agenda.
func (a *Agenda) UnmarshalJSON(data []byte) error {
	out := make(map[string]Update)
	err := json.Unmarshal(data, &out)
	if err != nil {
		return fmt.Errorf("parsing agenda: %w", err)
	}

	if a == nil {
		*a = make(map[Date]Update)
	}

	for d, update := range out {
		var date Date
		err = json.Unmarshal([]byte(strconv.Quote(d)), &date)
		if err != nil {
			return fmt.Errorf("parsing date in agenda: %v\n %w", d, err)
		}
		(*a)[date] = update
	}

	return nil
}

// AddUpdate will add a new update to the Agenda. This operation fails if an Update already exists for the given date.
func (a *Agenda) AddUpdate(date Date, plan Plan, doneList DoneList) error {
	_, ok := (*a)[date]
	if ok {
		return fmt.Errorf("update for date already exists: %v", date)
	}

	(*a)[date] = Update{plan, doneList}
	return nil
}

// AddDone adds done to the DoneList for the update on the given date. A new update will be created if none yet exists.
func (a *Agenda) AddDone(date Date, done string) {
	upd := (*a)[date]
	upd.DoneList = append(upd.DoneList, done)
	(*a)[date] = upd
}

// SetPlan sets the plan for the update on the given date. A new update will be created if none yet exists.
func (a *Agenda) SetPlan(date Date, plan Plan) {
	upd := (*a)[date]
	upd.Plan = plan
	(*a)[date] = upd
}

// LoadFile attempts to load an Agenda from a file located at fpath.
func LoadFile(fpath string) (*Agenda, error) {
	a := &Agenda{}

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return a, err
	}

	err = json.Unmarshal(data, a)
	return a, err
}

// WriteFile writes the agenda to fpath in JSON format.
func (a *Agenda) WriteFile(fpath string) error {
	data, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fpath, data, 0644)
}
