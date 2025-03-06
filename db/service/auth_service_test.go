package service

import (
	"context"
	"sandbox/config"
	"sandbox/db/models"
	"sandbox/lib/db"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationAndLogin(t *testing.T) {
	ctx := context.Background()
	dbConfig, err := config.NewDatabaseConfig()
	assert.NoError(t, err)

	conn, err := db.ConnectDB(ctx, dbConfig)
	assert.NoError(t, err)

	queries := models.New(conn)
	jwtConfig := config.AuthConfig{
		Jwt: &config.JwtConfig{
			Secret: "some-super-secret-token",
			Issuer: "sample.com",
			Expiry: time.Hour,
		},
		AuthUserContextKey: "user",
	}

	userService := NewAuthService(queries, &jwtConfig)

	testUser := func(t testing.TB, userInput models.UserCreateParams, user *models.User) {
		t.Helper()
		assert.Equal(t, userInput.ID.String(), user.ID.String())
		assert.Equal(t, userInput.Email, user.Email)
		assert.NotEqual(t, userInput.Password, user.Password)
		assert.Equal(t, userInput.Name, user.Name)
		assert.Equal(t, userInput.Role, user.Role)
	}

	userInput := models.UserCreateParams{
		ID:       uuid.New(),
		Email:    "customer@site.com",
		Password: "secret-password",
		Name:     "Mr. Customer",
		Role:     "CUSTOMER",
	}

	t.Run("register new user", func(st *testing.T) {
		createdUser, err := userService.RegisterNewUser(ctx, userInput)
		assert.NoError(st, err)
		testUser(st, userInput, createdUser)
	})

	t.Run("valid user login", func(st *testing.T) {
		loginResult, err := userService.LoginUser(ctx, userInput.Email, userInput.Password)
		assert.NoError(st, err)
		testUser(st, userInput, loginResult.User)
		assert.NotEqual(st, loginResult.Token, "")
		assert.NotEqual(st, loginResult.Expiry, 0)
	})

	t.Run("invalid user email", func(st *testing.T) {
		_, err := userService.LoginUser(ctx, "some-wrong-email@site.com", userInput.Password)
		assert.Error(st, err)
	})

	t.Run("invalid password", func(st *testing.T) {
		_, err := userService.LoginUser(ctx, userInput.Email, "some-random-wrong-password")
		assert.Error(st, err)
	})

	t.Cleanup(func() {
		_, err := conn.Exec(ctx, "delete from users where id = $1", userInput.ID)
		assert.NoError(t, err)
		conn.Close(ctx)
	})
}
