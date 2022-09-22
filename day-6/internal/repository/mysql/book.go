package mysql

import (
	"time"

	"github.com/nurhidaylma/alterra-agmc/day-6/internal/repository"
)

func GetBook() (interface{}, error) {
	now := time.Now().UTC()

	book := &repository.Book{
		ID:        1,
		Title:     "The Little House",
		Writer:    "Laura Wilder",
		NoOfPage:  150,
		Stock:     12,
		Price:     1500,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return book, nil
}

func GetBookByID(bookID uint) (*repository.Book, error) {
	now := time.Now().UTC()

	book := &repository.Book{
		ID:        1,
		Title:     "The Little House",
		Writer:    "Laura Wilder",
		NoOfPage:  150,
		Stock:     12,
		Price:     1500,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return book, nil
}

func CreateBook(book *repository.Book) (interface{}, error) {
	return book, nil
}

func UpdateBook(book *repository.Book) (interface{}, error) {
	return &book, nil
}

func DeleteBook(book *repository.Book) error {
	return nil
}
