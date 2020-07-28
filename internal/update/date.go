package update

import (
	"time"
)

// A Date consists of a year, month, and day.
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) String() string {
	return time.Date(d.Year, d.Month, d.Day, 12, 0, 0, 0, time.Local).Format("2006-01-02")
}

// MarshalJSON converts a Date to a JSON string.
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
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

	d.FromTime(t)
	return nil
}

// FromTime sets the Date from a time.Time value.
func (d *Date) FromTime(t time.Time) {
	year, month, day := t.Date()
	d.Year = year
	d.Month = month
	d.Day = day
}

// Before returns true if the date occurred before compareDate.
func (d Date) Before(compareDate Date) bool {
	if d.Year < compareDate.Year {
		return true
	}
	if d.Month < compareDate.Month {
		return true
	}
	if d.Day < compareDate.Day {
		return true
	}
	return false
}

// ToDate converts a time.ToDate to a Date.
func ToDate(t time.Time) Date {
	var d Date
	d.FromTime(t)
	return d
}

// Today returns today's date.
func Today() Date {
	return ToDate(time.Now())
}
