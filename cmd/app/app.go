package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sandbox/config"
	"sandbox/controller"
	"sandbox/db/models"
	"sandbox/db/service"
	"sandbox/lib/db"
	"sandbox/lib/server"

	"github.com/labstack/echo/v4"
)

func run(ctx context.Context) error {
	// ---------------------------------------------------------------------------
	//
	// initialize configs and database here.
	//
	// ---------------------------------------------------------------------------
	config, err := config.NewConfig()
	if err != nil {
		return err
	}

	conn, err := db.ConnectDB(ctx, config.Database)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	// ---------------------------------------------------------------------------
	//
	// initialize all dependencies here.
	//
	// ---------------------------------------------------------------------------
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	models := models.New(conn)
	authService := service.NewAuthService(models, config.Jwt)
	authController := controller.NewAuthController(logger, authService)
	// authMiddleware := middleware.NewAuthMiddleware(config.Jwt, config.Auth)

	// ---------------------------------------------------------------------------
	//
	// initialize server and routes here.
	//
	// ---------------------------------------------------------------------------
	router := echo.New()
	router.HTTPErrorHandler = server.GlobalErrorHandler

	api := router.Group("/api")
	{
		api.POST("/login", authController.Login)
		api.POST("/register", authController.RegisterNewUser)
		// api.GET("/protected", protectedHandler.ProtectedRoute, authMiddleware.IsLoggedIn)
	}

	// ---------------------------------------------------------------------------
	//
	// start server process.
	//
	// ---------------------------------------------------------------------------
	address := config.Server.Address()
	return router.Start(address)
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}

// TODO: remove after testing.
// type ProtectedHandler struct {
// 	authConfig *config.AuthConfig
// }

// func (h *ProtectedHandler) ProtectedRoute(c echo.Context) error {
// 	user, err := server.CurrentUser(c, h.authConfig.AuthUserContextKey)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, map[string]any{
// 		"message":      "you have reached a protected route",
// 		"current_user": user,
// 	})
// }
