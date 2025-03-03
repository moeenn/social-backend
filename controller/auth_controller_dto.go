package controller

import (
	"errors"
	"sandbox/db/service"
	"strings"
)

const PASSWORD_MIN_LENGTH = 8

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequestBody) Validate() error {
	if r.Email == "" || !strings.Contains(r.Email, "@") {
		return errors.New("invalid email address")
	}

	if r.Password == "" || len(r.Password) < PASSWORD_MIN_LENGTH {
		return errors.New("invalid password")
	}

	return nil
}

type LoginResponseUser struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type LoginResponse struct {
	User   LoginResponseUser `json:"user"`
	Token  string            `json:"token"`
	Expiry int64             `json:"expiry"`
}

func LoginResponseFromLoginResult(result *service.LoginResult) *LoginResponse {
	return &LoginResponse{
		User: LoginResponseUser{
			Id:    result.User.ID.String(),
			Email: result.User.Email,
			Role:  result.User.Role,
		},
		Token:  result.Token,
		Expiry: result.Expiry,
	}
}

type RegisterNewUserResquestBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *RegisterNewUserResquestBody) Validate() error {
	if r.Email == "" || !strings.Contains(r.Email, "@") {
		return errors.New("invalid email address")
	}

	if r.Password == "" || len(r.Password) < PASSWORD_MIN_LENGTH {
		return errors.New("invalid password")
	}

	if r.Password != r.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	return nil
}
