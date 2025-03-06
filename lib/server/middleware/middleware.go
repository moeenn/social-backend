package middleware

import (
	"strings"

	"sandbox/config"
	"sandbox/lib/authToken"

	"github.com/labstack/echo/v4"
)

type AuthMiddeware struct {
	jwtConfig  *authToken.JwtConfig
	authConfig *config.AuthConfig
}

func NewAuthMiddleware(authConfig *config.AuthConfig) *AuthMiddeware {
	return &AuthMiddeware{
		jwtConfig: &authToken.JwtConfig{
			Issuer: authConfig.Jwt.Issuer,
			Secret: authConfig.Jwt.Secret,
			Expiry: authConfig.Jwt.Expiry,
		},
		authConfig: authConfig,
	}
}

func (m *AuthMiddeware) IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		authTokenPayload, err := authToken.VerifyToken(tokenString, m.jwtConfig)
		if err != nil {
			return echo.ErrUnauthorized
		}

		c.Set(m.authConfig.AuthUserContextKey, authTokenPayload)
		if err := next(c); err != nil {
			c.Error(err)
			return nil
		}

		return nil
	}
}

// TODO: implement HasRole middleware.
