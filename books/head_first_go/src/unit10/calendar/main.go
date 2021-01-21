package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetYear(y int) error {
	if y < 1 {
		return errors.New("invalid y")
	}
	d.year = y
	return nil
}

func (d *Date) SetMonth(m int) error {
	if m < 1 {
		return errors.New("invalid m")
	}
	d.month = m
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
