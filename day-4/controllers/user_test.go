package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nurhidaylma/alterra-agmc/day-4/config"
	"github.com/nurhidaylma/alterra-agmc/day-4/database"
	"github.com/nurhidaylma/alterra-agmc/day-4/middlewares"
	"github.com/nurhidaylma/alterra-agmc/day-4/mocks"
	"github.com/stretchr/testify/assert"
)

var echoMock = mocks.EchoMock{E: echo.New()}

func TestLoginUserControllerInvalidPayload(t *testing.T) {
	// setup database
	config.GetConnection()
	database.NewSeeder().Seed()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPost, "/login", nil)
	c.SetPath("/jwt/login")

	// setup handler
	asserts := assert.New(t)
	db := config.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Invalid parameters or payload","info": null},"error": "bad_request"}`, body)
	}
}

func TestLoginUserControllerWrongPassword(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "vincentlhubbard@superrito.com",
		Password: "1234567890",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Email or password is incorrect","info": null},"error": "bad_request"}`, body)
	}
}

func TestLoginUserControllerSuccess(t *testing.T) {
	// setup database
	config.GetConnection()
	database.NewSeeder().Seed()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPost, "/login", nil)
	c.SetPath("/jwt/login")

	// setup handler
	asserts := assert.New(t)
	db := config.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Invalid parameters or payload","info": null},"error": "bad_request"}`, body)
	}
}

func TestEmployeeHandlerGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := middlewares.CreateToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/employees")
	c.QueryParams().Add("page", "a")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(employeeHandler.Get(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}
