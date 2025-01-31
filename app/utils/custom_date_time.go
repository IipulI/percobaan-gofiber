package utils

import (
	"bytes"
	"time"
)

type CustomDateTime time.Time

const dateTimeFormat = "2006-01-02 15:04:05"

func NewCustomDateTime(t time.Time) *CustomDateTime {
	ct := CustomDateTime(t)
	return &ct
}

func (tm *CustomDateTime) ToTime() time.Time {
	return time.Time(*tm)
}

func (ct CustomDateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Format(dateTimeFormat) + `"`), nil
}

func (ct *CustomDateTime) UnmarshalJSON(data []byte) error {
	// Check for null value
	if bytes.Equal(data, []byte("null")) {
		*ct = CustomDateTime{}
		return nil
	}

	// Parse the date string
	parsedTime, err := time.Parse(`"`+dateTimeFormat+`"`, string(data))
	if err != nil {
		return err
	}

	*ct = CustomDateTime(parsedTime)
	return nil
}
