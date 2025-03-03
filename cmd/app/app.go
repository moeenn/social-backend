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
