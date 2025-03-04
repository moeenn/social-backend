package server

import (
	"sandbox/lib/authToken"

	"github.com/labstack/echo/v4"
)

func CurrentUser(c echo.Context, authUserContextKey string) (*authToken.TokenPayload, error) {
	userPayload, ok := c.Get(authUserContextKey).(*authToken.TokenPayload)
	if !ok {
		return nil, echo.ErrUnauthorized
	}

	return userPayload, nil
}
