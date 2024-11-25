package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/IipulI/percobaan-gofiber/app/model"
)

type BookRepository interface {
	Insert(ctx context.Context, book model.Book) (model.Book, error)
	FindById(ctx context.Context, id int) (model.Book, error)
	FindAll(ctx context.Context) ([]model.Book, error)
	Update(ctx context.Context, id int, b *model.Book) (string, error)
	Delete(ctx context.Context, id int) (string, error)
}

type BookRepositoryImpl struct{ DB *sql.DB }

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (repository *BookRepositoryImpl) Insert(ctx context.Context, book model.Book) (model.Book, error) {
	script := "INSERT INTO book (id, name) values(?,?)"
	_, err := repository.DB.ExecContext(ctx, script, book.Id, book.Name)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, id int) (model.Book, error) {
	script := "SELECT id, name from book where id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	book := model.Book{}
	if err != nil {
		return book, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&book.Id, &book.Name)
		return book, nil
	} else {
		return book, errors.New("Id" + strconv.Itoa(int(id)) + "tidak ditemukan")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context) ([]model.Book, error) {
	script := "SELECT id, name from book"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var books []model.Book

	for rows.Next() {
		book := model.Book{}
		rows.Scan(&book.Id, &book.Name)
		books = append(books, book)
	}

	return books, nil
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, id int, b *model.Book) (string, error) {
	script := "UPDATE book SET name=? WHERE id=?"
	result, err := repository.DB.ExecContext(ctx, script, b.Name, id)
	if err != nil {
		return "", fmt.Errorf("error executing update query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected != 1 {
		return "", fmt.Errorf("update failed: expected 1 row affected, but got %d", rowsAffected)
	}

	return "1 row successfully updated", nil
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, id int) (string, error) {
	script := "delete from book where id=?"
	_, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return "", err
	} else {
		return "Row deleted successfully", nil
	}
}
