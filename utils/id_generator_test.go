package utils

import (
	"context"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGenerateRandomID(t *testing.T) {
	t.Run("generate a random id with given prefix", func(t *testing.T) {
		prefix := "acc_"
		ctx := context.Background()
		response := GenerateRandomID(ctx, prefix)

		// checking if prefix is added in the id at front
		assert.Equal(t, response[:len(prefix)], prefix)
		// random id generated will be of 20 characters long string
		assert.Equal(t, len(response), len(prefix)+20)
	})
}
