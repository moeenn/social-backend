package controller

import (
	"log/slog"
	"net/http"
	"sandbox/db/models"
	"sandbox/db/service"

	"sandbox/lib/server"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	logger      *slog.Logger
	authService *service.AuthService
}

func NewAuthController(logger *slog.Logger, authService *service.AuthService) *AuthController {
	return &AuthController{
		logger:      logger,
		authService: authService,
	}
}

func (c *AuthController) Login(ctx echo.Context) error {
	var loginRequestBody LoginRequestBody
	if err := ctx.Bind(&loginRequestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := loginRequestBody.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	loginResult, err := c.authService.LoginUser(
		ctx.Request().Context(), loginRequestBody.Email, loginRequestBody.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := LoginResponseFromLoginResult(loginResult)
	return ctx.JSON(http.StatusOK, resp)
}

func (c *AuthController) RegisterNewUser(ctx echo.Context) error {
	var registerNewUserRequestBody RegisterNewUserResquestBody
	if err := ctx.Bind(&registerNewUserRequestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := registerNewUserRequestBody.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	createUserParams := models.UserCreateParams{
		ID:       uuid.New(),
		Email:    registerNewUserRequestBody.Email,
		Password: registerNewUserRequestBody.Password,
		Name:     registerNewUserRequestBody.Name,
		Role:     "USER",
	}

	_, err := c.authService.RegisterNewUser(ctx.Request().Context(), createUserParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, server.MessageResponse{
		Message: "user registered successfully",
	})
}
