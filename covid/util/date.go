package util

import (
	"fmt"
	"strconv"
	"time"
)

const YEARFORMAT = "2006"
const DATEFORMAT = "2006-01-02"
const TIMEFORMAT = "15:04:05"
const TIMEFORMAT2 = "15:04"
const DATETIMEFORMAT = "2006-01-02 15:04:05"
const DATETIMEFORMA2 = "2006-01-02 15:04"

type JSONDateTime struct {
	time.Time
}

func (t JSONDateTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(DATETIMEFORMAT))
	return []byte(stamp), nil
}

func (t *JSONDateTime) UnmarshalJSON(in []byte) error {
	s, err := strconv.Unquote(string(in))
	if err != nil {
		return err
	}
	t.Time, err = time.ParseInLocation(DATETIMEFORMAT, s, time.Local)
	if err != nil {
		t.Time, err = time.ParseInLocation(DATETIMEFORMA2, s, time.Local)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t JSONDateTime) String() string {
	return t.Format(DATETIMEFORMAT)
}

type JSONDate struct {
	time.Time
}

func (t JSONDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(DATEFORMAT))
	return []byte(stamp), nil
}

func (t *JSONDate) UnmarshalJSON(in []byte) error {
	s, err := strconv.Unquote(string(in))
	if err != nil {
		return err
	}
	t.Time, err = time.Parse(DATEFORMAT, s)
	return err
}

func (t JSONDate) String() string {
	return t.Format(DATEFORMAT)
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(TIMEFORMAT))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(in []byte) error {
	s, err := strconv.Unquote(string(in))
	if err != nil {
		return err
	}
	t.Time, err = time.Parse(TIMEFORMAT, s)
	return err
}

func (t JSONTime) String() string {
	return t.Format(TIMEFORMAT)
}