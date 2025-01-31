package model

import (
	"github.com/IipulI/percobaan-gofiber/app/utils"
)

type BookCopies struct {
	Id         int32                 `json:"id"`
	BookId     int32                 `json:"book_id"`
	CopyNumber int32                 `json:"copy_number"`
	Status     string                `json:"status"`
	CreatedAt  *utils.CustomDateTime `json:"created_at"`
	UpdatedAt  *utils.CustomDateTime `json:"updated_at"`
	DeletedAt  *utils.CustomDateTime `json:"deleted_at"`
}
