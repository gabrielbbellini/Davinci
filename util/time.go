package util

import (
	"fmt"
	"strings"
	"time"
)

func NewDateTime(convert *time.Time) *DateTime {
	if convert != nil {
		x := DateTime(*convert)
		return &x
	}
	return nil
}

func NewDate(convert *time.Time) *Date {
	if convert != nil {
		x := Date(*convert)
		return &x
	}
	return nil
}

func DateTimeByString(convert string) *DateTime {
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, convert)
	if err == nil {
		d := DateTime(t)
		return &d
	}

	return nil
}

func DateByStringTemplate(tmpl, convert string) *Date {
	t, err := time.Parse(tmpl, convert)
	if err == nil {
		d := Date(t)
		return &d
	}

	return nil
}

// <editor-fold  desc="DateTime faz a serialização dos timestamps no formato AAAA-MM-DD HH:mm:ss" defaultstate="collapsed">
type DateTime time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	date := time.Time(d).Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("%q", date)), nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	date := strings.Trim(string(b), `"`)
	if date == "null" {
		*d = DateTime(time.Time{})
		return nil
	} else {
		tmp, err := time.Parse("2006-01-02 15:04:05", date)
		if err == nil {
			*d = DateTime(tmp)
		}
		return err
	}
}

func (d *DateTime) Time() *time.Time {
	if d == nil {
		return nil
	}
	t := time.Time(*d)
	return &t
}

// </editor-fold>

// <editor-fold desc="Date faz a serialização de datas simples, no formato AAAA-MM-DD" defaultstate="collapsed">
type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	df := time.Time(d).Format("2006-01-02")
	return []byte(fmt.Sprintf("%q", df)), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	date := strings.Trim(string(b), `"`)
	if date == "null" {
		*d = Date(time.Time{})
		return nil
	} else {
		tmp, err := time.Parse("2006-01-02", date)
		if err == nil {
			*d = Date(tmp)
		}
		return err
	}
}

func (d *Date) Time() *time.Time {
	if d == nil {
		return nil
	}
	t := time.Time(*d)
	return &t
}

// </editor-fold>
