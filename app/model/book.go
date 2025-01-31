package model

import "github.com/IipulI/percobaan-gofiber/app/utils"

type Book struct {
	Id        int32                 `json:"id"`
	Title     string                `json:"title"`
	Author    string                `json:"author"`
	ISBN      string                `json:"isbn"`
	Page      int32                 `json:"page"`
	CreatedAt *utils.CustomDateTime `json:"created_at"`
	UpdatedAt *utils.CustomDateTime `json:"updated_at"`
	DeletedAt *utils.CustomDateTime `json:"deleted_at"`
}
