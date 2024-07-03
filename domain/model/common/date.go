package common

import (
	"go-complaint/domain"
	"strconv"
	"time"
)

type Date struct {
	stringRepresentation string
	date                 time.Time
	director             domain.Director
}

func NewDateWithDirector(director domain.Director) *Date {
	d := &Date{
		director: director,
	}
	return d
}

func (d *Date) Changed() {
	d.director.Changed(d)
}

func NewDateFromString(date string) (Date, error) {
	d := new(Date)
	d.stringRepresentation = date
	integer, err := strconv.ParseInt(date, 10, 64)
	if err != nil {
		return *d, err
	}
	d.date = time.UnixMilli(integer)
	return *d, nil
}

// manage time zero and time future before its too late
func NewDate(date time.Time) Date {
	d := new(Date)
	d.date = date
	milliSeconds := d.date.UnixMilli()
	milliSecondsString := strconv.FormatInt(milliSeconds, 10)
	d.stringRepresentation = milliSecondsString
	return *d
}

func (d Date) Date() time.Time {
	return d.date
}
func (d Date) Age() int {
	today := time.Now()
	birhtDate := d.date
	age := today.Sub(birhtDate).Hours() / 24 / 365
	intAge := int(age)
	return intAge
}

func (d Date) StringRepresentation() string {
	return d.stringRepresentation
}

func StringDate(date time.Time) string {
	milliSeconds := date.UnixMilli()
	milliSecondsString := strconv.FormatInt(milliSeconds, 10)
	return milliSecondsString
}

func ParseDate(date string) (time.Time, error) {
	integer, err := strconv.ParseInt(date, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.UnixMilli(integer), nil
}
