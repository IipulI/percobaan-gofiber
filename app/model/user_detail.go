package model

import "github.com/IipulI/percobaan-gofiber/app/utils"

type UserDetail struct {
	Id             int32            `json:"id"`
	Username       string           `json:"username"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	Address        string           `json:"address"`
	PhoneNumber    string           `json:"phone_number"`
	Gender         string           `json:"gender"`
	DateOfBirth    utils.CustomDate `json:"date_of_birth"`
	ProfilePicture string           `json:"profile_picture"`
}
