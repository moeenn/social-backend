package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GlobalErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		message, ok := he.Message.(string)
		if !ok {
			message = "unknown error occurred"
		}

		_ = c.JSON(he.Code, ErrorResponse{
			Error: message,
		})

		return
	}

	_ = c.JSON(code, ErrorResponse{
		Error: err.Error(),
	})
}
