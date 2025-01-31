package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/IipulI/percobaan-gofiber/app/model"
)

type UserDetailRepository interface {
	GetUserDetail(ctx context.Context, username string) (model.UserDetail, error)
	UpdateUserDetail(ctx context.Context, username string, paramUserDetail *model.UserDetail) (model.UserDetail, error)
}

type UserDetailRepositoryImpl struct{ DB *sql.DB }

func NewUserDetailRepository(db *sql.DB) UserDetailRepository {
	return &UserDetailRepositoryImpl{DB: db}
}

func (repository *UserDetailRepositoryImpl) GetUserDetail(ctx context.Context, username string) (model.UserDetail, error) {
	script := "SELECT id, username, first_name, last_name, address, phone_number, gender, date_of_birth, profile_picture FROM user_details where username = ?"
	rows, err := repository.DB.QueryContext(ctx, script, username)

	userDetail := model.UserDetail{}
	if err != nil {
		return userDetail, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&userDetail.Id, &userDetail.Username, &userDetail.FirstName, &userDetail.LastName, &userDetail.Address, &userDetail.PhoneNumber, &userDetail.Gender, &userDetail.DateOfBirth, &userDetail.ProfilePicture)

		return userDetail, nil
	} else {
		return userDetail, errors.New("username :" + username + " tidak ditemukan")
	}
}

func (repository *UserDetailRepositoryImpl) UpdateUserDetail(ctx context.Context, username string, paramUserDetail *model.UserDetail) (model.UserDetail, error) {
	script := "UPDATE user_details SET first_name=?, last_name=?, address=?, phone_number=?, gender=?, date_of_birth=?, profile_picture=? where username=?"
	result, err := repository.DB.ExecContext(ctx, script, paramUserDetail.FirstName, paramUserDetail.LastName, paramUserDetail.Address, paramUserDetail.PhoneNumber, paramUserDetail.Gender, paramUserDetail.DateOfBirth, paramUserDetail.ProfilePicture, username)

	if err != nil {
		return *paramUserDetail, errors.New("error executing update query")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return *paramUserDetail, fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected != 1 {
		return *paramUserDetail, fmt.Errorf("update failed: expected 1 row affected, but got %d", rowsAffected)
	}

	return *paramUserDetail, nil
}
