package service

import (
	"context"
	"errors"
	"fmt"
	"sandbox/config"
	"sandbox/db/constraints"
	"sandbox/db/models"
	"sandbox/lib/authToken"
	"sandbox/lib/hash"
)

type AuthService struct {
	db        *models.Queries
	jwtConfig *authToken.JwtConfig
}

func NewAuthService(db *models.Queries, jwtConfig *config.JwtConfig) *AuthService {
	return &AuthService{
		db: db,
		jwtConfig: &authToken.JwtConfig{
			Secret: jwtConfig.Secret,
			Issuer: jwtConfig.Issuer,
			Expiry: jwtConfig.Expiry,
		},
	}
}

func (s *AuthService) RegisterNewUser(ctx context.Context, input models.UserCreateParams) (*models.User, error) {
	passwordHash, err := hash.HashPassword(input.Password, hash.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	input.Password = passwordHash
	createdUser, err := s.db.UserCreate(ctx, input)
	err = constraints.ProcessConstraintError(err, constraints.UserConstraints)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

type LoginResult struct {
	User   *models.User
	Token  string
	Expiry int64
}

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (*LoginResult, error) {
	credentialsErr := errors.New("invalid email or password")
	user, err := s.db.UserByEmail(ctx, email)
	if err != nil {
		return nil, credentialsErr
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, credentialsErr
	}

	tokenPayload := authToken.TokenPayload{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	token, expiry, err := authToken.CreateToken(&tokenPayload, s.jwtConfig)
	if err != nil {
		return nil, err
	}

	loginResult := LoginResult{
		User:   &user,
		Token:  token,
		Expiry: expiry,
	}

	return &loginResult, nil
}
