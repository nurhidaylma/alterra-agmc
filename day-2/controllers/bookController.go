package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nurhidaylma/alterra-agmc/day-2/lib/database"
	"github.com/nurhidaylma/alterra-agmc/day-2/models"
)

func GetBookController(c echo.Context) error {
	books, err := database.GetBook()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookByIDController(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("id"))

	book, err := database.GetBookByID(uint(bookID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func CreateBookController(c echo.Context) error {
	noOfPage, _ := strconv.Atoi(c.FormValue("no_of_page"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))
	price, _ := strconv.Atoi(c.FormValue("price"))

	book, err := database.CreateBook(&models.Book{
		Title:    c.FormValue("title"),
		Writer:   c.FormValue("writer"),
		NoOfPage: noOfPage,
		Stock:    stock,
		Price:    float32(price),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	noOfPage, _ := strconv.Atoi(c.FormValue("no_of_page"))
	stock, _ := strconv.Atoi("stock")
	price, _ := strconv.Atoi("price")

	book, err := database.UpdateBook(&models.Book{
		ID:       uint(id),
		Title:    c.FormValue("title"),
		Writer:   c.FormValue("writer"),
		NoOfPage: noOfPage,
		Stock:    stock,
		Price:    float32(price),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := database.GetBookByID(uint(id))
	if err != nil {
		return err
	}

	err = database.DeleteBook(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
