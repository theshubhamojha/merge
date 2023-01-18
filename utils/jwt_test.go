package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateJWTToken(t *testing.T) {
	t.Run("generate a valid jwt token and verify it", func(t *testing.T) {
		response, err := GenerateJWTToken("random@gmail.com", "admin", "acc_123", "secret", 100)
		assert.Equal(t, err, nil)

		resp, err := VerifyJWTToken(response, "secret")
		assert.Equal(t, err, nil)
		assert.Equal(t, resp["email"], "random@gmail.com")
		assert.Equal(t, resp["role"], "admin")
		assert.Equal(t, resp["account_id"], "acc_123")
	})

	t.Run("pass an invalid jwt token to the verify function", func(t *testing.T) {
		_, err := VerifyJWTToken("invalid_jwt.invalid.invalid", "secret")
		require.Error(t, err)
	})
}
