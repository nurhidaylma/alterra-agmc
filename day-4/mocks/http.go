package mocks

import (
	"io"
	"net/http/httptest"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nurhidaylma/alterra-agmc/day-4/middlewares"
)

type EchoMock struct {
	E *echo.Echo
}

func (em *EchoMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	em.E.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)

	return c, rec
}
