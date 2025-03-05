package authToken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPayload struct {
	Id    uuid.UUID
	Email string
	Role  string
}

type JwtConfig struct {
	Issuer string
	Secret string
	Expiry time.Duration
}

func CreateToken(payload *TokenPayload, config *JwtConfig) (string, int64, error) {
	expiry := time.Now().Add(config.Expiry).Unix()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": payload.Id,
		"sub":    payload.Email,
		"iss":    config.Issuer,
		"aud":    payload.Role,
		"exp":    expiry,
		"iat":    time.Now().Unix(), // Issued at
	})

	tokenString, err := claims.SignedString([]byte(config.Secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiry, nil
}

func VerifyToken(tokenString string, config *JwtConfig) (*TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	email, err := token.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	role, err := token.Claims.GetAudience()
	if err != nil || len(role) == 0 {
		return nil, err
	}

	claimsMap, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	userId, ok := claimsMap["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid token claims")
	}

	expiry, err := token.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if time.Now().After(expiry.Time) {
		return nil, fmt.Errorf("token expired")
	}

	tokenPayload := TokenPayload{
		Id:    parsedUserId,
		Email: email,
		Role:  role[0],
	}

	return &tokenPayload, nil
}
