package database

import (
	"github.com/nurhidaylma/alterra-agmc/day-2/config"
	"github.com/nurhidaylma/alterra-agmc/day-2/models"
)

func GetBook() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func GetBookByID(bookID uint) (*models.Book, error) {
	book := models.Book{}
	if err := config.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func CreateBook(book *models.Book) (interface{}, error) {
	if err := config.DB.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(book *models.Book) (interface{}, error) {
	if err := config.DB.Model(book).Updates(book).Error; err != nil {
		return nil, err
	}
	if err := config.DB.First(&book, "id = ?", book.ID).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func DeleteBook(book *models.Book) error {
	if err := config.DB.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
