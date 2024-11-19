package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/IipulI/percobaan-gofiber/app/model"
	"github.com/IipulI/percobaan-gofiber/app/utils"
)

type UserRepository interface {
	Login(ctx context.Context, user string, password string) (model.User, error)
}

type UserRepositoryImpl struct{ DB *sql.DB }

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, user string, password string) (model.User, error) {
	script := "select username, email, password  from `user` u where username = ? or email = ?"
	rows, err := repository.DB.QueryContext(ctx, script, user, user)

	userStruct := model.User{}
	if err != nil {
		return userStruct, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&userStruct.Username, &userStruct.Email, &userStruct.Password)
	} else {
		return userStruct, errors.New("user not found")
	}

	if !utils.ComparePassword(userStruct.Password, password) {
		return userStruct, errors.New("password not same")
	}

	return userStruct, nil
}