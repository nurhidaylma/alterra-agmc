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

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserByIDController)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}
