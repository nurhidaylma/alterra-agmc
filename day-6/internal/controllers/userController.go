package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nurhidaylma/alterra-agmc/day-6/internal/public"
	"github.com/nurhidaylma/alterra-agmc/day-6/internal/repository"
	"github.com/nurhidaylma/alterra-agmc/day-6/internal/repository/mysql"
	"github.com/nurhidaylma/alterra-agmc/day-6/middlewares"
)

func LoginUserController(c echo.Context) error {
	user := public.LoginRequest{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := mysql.LoginUser(&repository.User{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"token":  token,
	})
}

func GetUsersController(c echo.Context) error {
	users, err := mysql.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIDController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	user, err := mysql.GetUserByID(uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func CreateUserController(c echo.Context) error {
	userRepo, err := mysql.GetUserByEmail(c.FormValue("email"))
	if err != nil {
		return err
	}
	if userRepo != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "email already exists",
		})
	}

	age, _ := strconv.Atoi(c.FormValue("age"))
	payload := repository.User{
		Email:    c.FormValue("email"),
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
		FullName: c.FormValue("full_name"),
		Gender:   repository.GenderTypes(c.FormValue("gender")),
		Age:      age,
	}

	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := mysql.CreateUser(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	age, _ := strconv.Atoi(c.FormValue("age"))
	payload := public.UpdateUserRequest{
		ID:       uint(id),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Username: c.FormValue("username"),
		Fullname: c.FormValue("full_name"),
		Gender:   repository.GenderTypes(c.FormValue("gender")),
		Age:      age,
	}

	if err := c.Validate(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID := middlewares.ExtractTokenUserID(c)
	if userID != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	user, err := mysql.UpdateUser(&repository.User{
		ID:       uint(id),
		Email:    payload.Email,
		Password: payload.Password,
		Username: payload.Username,
		FullName: payload.Fullname,
		Gender:   payload.Gender,
		Age:      payload.Age,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userID := middlewares.ExtractTokenUserID(c)
	if userID != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := mysql.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "failed",
			"message": "record not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
