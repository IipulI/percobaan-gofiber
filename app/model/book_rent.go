package model

import "github.com/IipulI/percobaan-gofiber/app/utils"

type BookRent struct {
	Id                int32                 `json:"id"`
	BookCopyId        int32                 `json:"book_copy_id"`
	MemberId          int32                 `json:"member_id"`
	RentDate          *utils.CustomDate     `json:"rent_date"`
	DueDate           *utils.CustomDate     `json:"due_date"`
	ReturnDate        *utils.CustomDate     `json:"return_date"`
	ConditionReturned *string               `json:"condition_returned"`
	Status            *string               `json:"status"`
	CreatedAt         *utils.CustomDateTime `json:"created_at"`
	UpdatedAt         *utils.CustomDateTime `json:"updated_at"`
	DeletedAt         *utils.CustomDateTime `json:"deleted_at"`
}
