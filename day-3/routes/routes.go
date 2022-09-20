package routes

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nurhidaylma/alterra-agmc/day-3/config"
	c "github.com/nurhidaylma/alterra-agmc/day-3/constants"
	"github.com/nurhidaylma/alterra-agmc/day-3/controllers"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(config.GetValue(c.SECRET_JWT))))

	// User
	r.GET("/users/:id", controllers.GetUserByIDController)
	r.GET("/users", controllers.GetUsersController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)

	// Book
	r.POST("/books", controllers.CreateBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)

	// Book
	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookByIDController)

	// User
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)

	return e
}
