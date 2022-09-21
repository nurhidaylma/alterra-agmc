package database

import (
	"time"

	"github.com/nurhidaylma/alterra-agmc/day-4/models"
)

func GetBook() (interface{}, error) {
	now := time.Now().UTC()

	book := &models.Book{
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

func GetBookByID(bookID uint) (*models.Book, error) {
	now := time.Now().UTC()

	book := &models.Book{
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

func CreateBook(book *models.Book) (interface{}, error) {
	return book, nil
}

func UpdateBook(book *models.Book) (interface{}, error) {
	return &book, nil
}

func DeleteBook(book *models.Book) error {
	return nil
}
