package authToken

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTokenCreateVerify(t *testing.T) {
	testUser := TokenPayload{
		Id:    uuid.NewString(),
		Email: "some-test-user@site.com",
		Role:  "ADMIN",
	}

	config := JwtConfig{
		Issuer: "sample.com",
		Secret: "some-super-secret-token",
		Expiry: time.Hour,
	}

	token, expiry, err := CreateToken(&testUser, &config)
	assert.NoError(t, err)
	assert.NotEqual(t, token, "")
	assert.NotEqual(t, expiry, 0)

	payload, err := VerifyToken(token, &config)
	assert.NoError(t, err)
	assert.Equal(t, payload.Id, testUser.Id)
	assert.Equal(t, payload.Email, testUser.Email)
	assert.Equal(t, payload.Role, testUser.Role)
}
