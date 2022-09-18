package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nurhidaylma/alterra-agmc/day-2/lib/database"
	"github.com/nurhidaylma/alterra-agmc/day-2/models"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
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

	user, err := database.GetUserByID(uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func CreateUserController(c echo.Context) error {
	users, ok := database.GetUserByEmail(c.FormValue("email"))
	if ok != nil {
		return ok
	}
	if users != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "email already exists",
		})
	}

	age, _ := strconv.Atoi(c.FormValue("age"))
	user, err := database.CreateUser(&models.User{
		Email:    c.FormValue("email"),
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
		FullName: c.FormValue("full_name"),
		Gender:   models.GenderTypes(c.FormValue("gender")),
		Age:      age,
	})
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

	user, err := database.UpdateUser(&models.User{
		ID:       uint(id),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Username: c.FormValue("username"),
		FullName: c.FormValue("full_name"),
		Gender:   models.GenderTypes(c.FormValue("gender")),
		Age:      age,
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

	err := database.DeleteUser(uint(id))
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
