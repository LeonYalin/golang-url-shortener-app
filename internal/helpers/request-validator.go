package helpers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (this *RequestValidator) Validate(i interface{}) error {
	if err := this.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
