package utils

import (
	"context"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHashSalt(t *testing.T) {
	t.Run("when incoming password and actual password are same", func(t *testing.T) {
		ctx := context.Background()
		actualPasswordHash := HashAndSalt(ctx, []byte("password"))
		rawIncomingPassword := "password"
		isValid := VerifyHashSalt(ctx, rawIncomingPassword, actualPasswordHash)

		assert.Equal(t, isValid, true)
	})

	t.Run("when incoming password and actual password aren't same", func(t *testing.T) {
		ctx := context.Background()
		actualPasswordHash := HashAndSalt(ctx, []byte("password"))
		rawIncomingPassword := "wrong_password"
		isValid := VerifyHashSalt(ctx, rawIncomingPassword, actualPasswordHash)

		assert.Equal(t, isValid, false)
	})
}
