package common

import (
	"strconv"
	"time"
)

type Date struct {
	stringRepresentation string
	date                 time.Time
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
	age := today.Year() - d.date.Year()
	if today.YearDay() < d.date.YearDay() {
		age--
	}
	return age
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
