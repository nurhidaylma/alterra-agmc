package routes

import (
	"github.com/labstack/echo"
	"github.com/nurhidaylma/alterra-agmc/day-2/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookByIDController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)

	return e
}
