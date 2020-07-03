package update

import (
	"fmt"
	"time"
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
	s := fmt.Sprintf(`"%s"`, d)
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

// LessThan returns true if the date is less than the cmpDate.
func (d Date) LessThan(cmpDate Date) bool {
	if d.Year < cmpDate.Year {
		return true
	}
	if d.Month < cmpDate.Month {
		return true
	}
	if d.Day < cmpDate.Day {
		return true
	}
	return false
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
