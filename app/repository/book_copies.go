package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/IipulI/percobaan-gofiber/app/model"
)

type BookCopiesRepository interface {
	GetAllBookCopies(ctx context.Context) ([]model.BookCopies, error)
	GetBookCopyById(ctx context.Context, id int) (model.BookCopies, error)
	GetBookCopyByBook(ctx context.Context, id int) ([]model.BookCopies, error)
	ValidateBookCopyNumber(ctx context.Context, b model.BookCopies) error
	ValidateBookCopyStatus(ctx context.Context, id int) error
	InsertBookCopy(ctx context.Context, b model.BookCopies) (model.BookCopies, error)
	UpdateBookCopy(ctx context.Context, exec interface{}, b *model.BookCopies) (string, error)
}

type BookCopiesRepositoryImpl struct{ DB *sql.DB }

func NewBookCopiesRepository(db *sql.DB) BookCopiesRepository {
	return &BookCopiesRepositoryImpl{DB: db}
}

func (repository *BookCopiesRepositoryImpl) GetAllBookCopies(ctx context.Context) ([]model.BookCopies, error) {
	script := "SELECT id, book_id, copy_number, status, created_at, updated_at, deleted_at from book_copies"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var bookCopies []model.BookCopies

	for rows.Next() {
		bookCopy := model.BookCopies{}
		rows.Scan(&bookCopy.Id, &bookCopy.BookId, &bookCopy.CopyNumber, &bookCopy.Status, &bookCopy.CreatedAt, &bookCopy.UpdatedAt, &bookCopy.DeletedAt)
		bookCopies = append(bookCopies, bookCopy)
	}

	return bookCopies, nil
}

func (repository *BookCopiesRepositoryImpl) GetBookCopyById(ctx context.Context, id int) (model.BookCopies, error) {
	script := "SELECT * FROM book_copies where id=?"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	bookCopy := model.BookCopies{}
	if err != nil {
		return bookCopy, nil
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&bookCopy.Id, &bookCopy.BookId, &bookCopy.CopyNumber, &bookCopy.Status, &bookCopy.CreatedAt, &bookCopy.UpdatedAt, &bookCopy.DeletedAt)
		return bookCopy, nil
	} else {
		return bookCopy, errors.New("Salinan Buku id " + strconv.Itoa(int(id)) + " tidak ditemukan")
	}
}

func (repository *BookCopiesRepositoryImpl) GetBookCopyByBook(ctx context.Context, id int) ([]model.BookCopies, error) {
	script := "SELECT * FROM book_copies as bc WHERE bc.book_id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	var bookCopies []model.BookCopies
	if err != nil {
		return bookCopies, err
	}

	defer rows.Close()
	for rows.Next() {
		bc := model.BookCopies{}
		rows.Scan(&bc.BookId, &bc.Id, &bc.CopyNumber, &bc.Status, &bc.CreatedAt, &bc.UpdatedAt, &bc.DeletedAt)
		bookCopies = append(bookCopies, bc)
	}

	return bookCopies, nil
}

func (repository *BookCopiesRepositoryImpl) ValidateBookCopyNumber(ctx context.Context, b model.BookCopies) error {
	script := "SELECT book_id, copy_number FROM book_copies WHERE book_id=? ORDER BY copy_number DESC LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, b.BookId)
	if err != nil {
		return err
	}

	bookCopy := model.BookCopies{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&bookCopy.BookId, &bookCopy.CopyNumber)

		if b.CopyNumber <= bookCopy.CopyNumber {
			return errors.New("copy number invalid")
		}
	}

	return nil
}

func (repository *BookCopiesRepositoryImpl) ValidateBookCopyStatus(ctx context.Context, id int) error {
	script := "SELECT id, status FROM book_copies where id=?"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	if err != nil {
		return err
	}

	bookCopy := model.BookCopies{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&bookCopy.BookId, &bookCopy.Status)

		if string(bookCopy.Status) != "available" {
			return errors.New("status buku sedang tidak available")
		}
	}

	return nil
}

func (repository *BookCopiesRepositoryImpl) InsertBookCopy(ctx context.Context, b model.BookCopies) (model.BookCopies, error) {
	script := "INSERT into book_copies (book_id, copy_number, status) VALUES(?,?,?)"
	_, err := repository.DB.ExecContext(ctx, script, b.BookId, b.CopyNumber, b.Status)

	if err != nil {
		return b, err
	}

	return b, nil
}

func (repository *BookCopiesRepositoryImpl) UpdateBookCopy(ctx context.Context, exec interface{}, b *model.BookCopies) (string, error) {
	script := "UPDATE book_copies SET copy_number=?, status=?, updated_at=? WHERE id=?"

	var rows sql.Result
	var err error

	switch e := exec.(type) {
	case *sql.Tx:
		rows, err = e.ExecContext(ctx, script, b.CopyNumber, b.Status, b.UpdatedAt, b.Id)
	case *sql.DB:
		rows, err = e.ExecContext(ctx, script, b.CopyNumber, b.Status, b.UpdatedAt, b.Id)
	default:
		return "", sql.ErrConnDone
	}

	if err != nil {
		return "", err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected != 1 {
		return "", err
	}

	return "Update success", nil
}
