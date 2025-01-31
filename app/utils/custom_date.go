package utils

import (
	"bytes"
	"time"
)

type CustomDate time.Time

const dateFormat = "2006-01-02"

func NewCustomDate(t time.Time) *CustomDate {
	ct := CustomDate(t)
	return &ct
}

func (tm *CustomDate) ToTime() time.Time {
	return time.Time(*tm)
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	t := time.Time(cd)
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Format(dateFormat) + `"`), nil
}
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*cd = CustomDate{}
		return nil
	}

	parsedDate, err := time.Parse(`"`+dateFormat+`"`, string(data))
	if err != nil {
		return err
	}

	*cd = CustomDate(parsedDate)
	return nil
}
