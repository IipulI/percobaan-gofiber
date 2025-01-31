package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/IipulI/percobaan-gofiber/app/model"
)

type BookRentRepository interface {
	GetBookRent(ctx context.Context) ([]model.BookRent, error)
	GetBookRentById(ctx context.Context, id int) (model.BookRent, error)
	InsertBookRent(ctx context.Context, exec interface{}, b model.BookRent) (model.BookRent, error)
	UpdateBookRent(ctx context.Context, exec interface{}, b *model.BookRent) (string, error)
}

type BookRentRepositoryImpl struct{ DB *sql.DB }

func NewBookRentRepository(db *sql.DB) BookRentRepository {
	return &BookRentRepositoryImpl{DB: db}
}

func (repository *BookRentRepositoryImpl) GetBookRent(ctx context.Context) ([]model.BookRent, error) {
	script := "SELECT * FROM book_rents"
	result, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	var bookrents []model.BookRent

	for result.Next() {
		bookrent := model.BookRent{}
		result.Scan(&bookrent.Id, &bookrent.BookCopyId, &bookrent.MemberId, &bookrent.RentDate, &bookrent.DueDate, &bookrent.ReturnDate, &bookrent.ConditionReturned, &bookrent.Status, &bookrent.CreatedAt, &bookrent.UpdatedAt, &bookrent.DeletedAt)
		// log.Printf("CreatedAt: %v, Status: %v", bookrent.CreatedAt, bookrent.Status)
		bookrents = append(bookrents, bookrent)
	}

	return bookrents, nil
}

func (repository *BookRentRepositoryImpl) GetBookRentById(ctx context.Context, id int) (model.BookRent, error) {
	script := "SELECT * FROM book_rents WHERE id=?"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	if err != nil {
		return model.BookRent{}, err
	}

	bookRent := model.BookRent{}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(
			&bookRent.Id,
			&bookRent.BookCopyId,
			&bookRent.MemberId,
			&bookRent.RentDate,
			&bookRent.DueDate,
			&bookRent.ReturnDate,
			&bookRent.ConditionReturned,
			&bookRent.Status,
			&bookRent.CreatedAt,
			&bookRent.UpdatedAt,
			&bookRent.DeletedAt,
		)

		return bookRent, nil
	} else {
		return model.BookRent{}, errors.New("terjadi kesalahan mencari book rent")
	}
}

func (repository *BookRentRepositoryImpl) InsertBookRent(ctx context.Context, exec interface{}, b model.BookRent) (model.BookRent, error) {
	script := "INSERT INTO book_rents (book_copy_id, member_id, rent_date, due_date, status) VALUES (?,?,?,?,?)"

	var err error

	switch e := exec.(type) {
	case *sql.Tx:
		_, err = e.ExecContext(ctx, script, b.BookCopyId, b.MemberId, b.RentDate.ToTime(), b.DueDate.ToTime(), b.Status)
	case *sql.DB:
		_, err = e.ExecContext(ctx, script, b.BookCopyId, b.MemberId, b.RentDate.ToTime(), b.DueDate.ToTime(), b.Status)
	default:
		return model.BookRent{}, sql.ErrConnDone
	}

	if err != nil {
		return model.BookRent{}, err
	}

	return b, nil
}

func (repository *BookRentRepositoryImpl) UpdateBookRent(ctx context.Context, exec interface{}, b *model.BookRent) (string, error) {
	script := "UPDATE book_rents SET return_date=?, condition_returned=?, status=?, updated_at=? WHERE id=?"

	var rows sql.Result
	var err error

	switch e := exec.(type) {
	case *sql.Tx:
		rows, err = e.ExecContext(ctx, script, b.ReturnDate.ToTime(), b.ConditionReturned, b.Status, b.UpdatedAt.ToTime(), b.Id)
	case *sql.DB:
		rows, err = e.ExecContext(ctx, script, b.ReturnDate.ToTime(), b.ConditionReturned, b.Status, b.UpdatedAt.ToTime(), b.Id)
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
