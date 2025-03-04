package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Run("test valid passwords", func(st *testing.T) {
		passwords := []string{
			"secret-password",
			"short",
			"#4acna324u934203-str0ng-Pas3Word!",
		}

		for _, password := range passwords {
			hashedPassword, err := HashPassword(password, 2)
			assert.NoError(t, err)
			isValid := CheckPasswordHash(password, hashedPassword)
			assert.True(t, isValid)
		}
	})

	t.Run("test invalid password", func(st *testing.T) {
		invalidPassword := "invalid-password"
		hashedPassword, err := HashPassword(invalidPassword, 2)
		assert.NoError(t, err)
		isValid := CheckPasswordHash("wrong-password", hashedPassword)
		assert.False(t, isValid)
	})
}
