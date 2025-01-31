package utils

import (
	"bytes"
	"time"
)

type CustomTime time.Time

const timeFormat = "2006-01-02 15:04:05"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Format(timeFormat) + `"`), nil
}
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*ct = CustomTime{}
		return nil
	}

	// Parse the date string
	parsedTime, err := time.Parse(`"`+timeFormat+`"`, string(data))
	if err != nil {
		return err
	}

	*ct = CustomTime(parsedTime)
	return nil
}
