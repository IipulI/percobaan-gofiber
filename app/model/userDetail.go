package model

import (
	"encoding/json"
	"errors"
	"time"
)

type CustomDate struct {
	time.Time
}

// UnmarshalJSON implements the custom logic for parsing dates
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Remove surrounding quotes from the JSON string
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// Define the expected layout (e.g., "2006-01-02")
	layout := "2006-01-02"

	// Parse the date using the layout
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}

	// Assign the parsed time to the CustomDate
	cd.Time = parsedTime
	return nil
}

// MarshalJSON (optional) formats the date back into a string when sending JSON
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// Use the same layout to format the date
	layout := "2006-01-02"
	return json.Marshal(cd.Time.Format(layout))
}

type UserDetail struct {
	Id             int32      `json:"id"`
	Username       string     `json:"username"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Address        string     `json:"address"`
	PhoneNumber    string     `json:"phone_number"`
	Gender         string     `json:"gender"`
	DateOfBirth    CustomDate `json:"date_of_birth"`
	ProfilePicture string     `json:"profile_picture"`
}
