package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nurhidaylma/alterra-agmc/day-6/internal/repository"
	"github.com/nurhidaylma/alterra-agmc/day-6/internal/repository/mysql"
)

func GetBookController(c echo.Context) error {
	books, err := mysql.GetBook()
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

	book, err := mysql.GetBookByID(uint(bookID))
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

	book, err := mysql.CreateBook(&repository.Book{
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

	book, err := mysql.UpdateBook(&repository.Book{
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

	book, err := mysql.GetBookByID(uint(id))
	if err != nil {
		return err
	}

	err = mysql.DeleteBook(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
