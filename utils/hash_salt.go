package utils

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(ctx context.Context, password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(hash)
}

func VerifyHashSalt(ctx context.Context, rawIncomingPassword string, actualPasswordHash string) (isValid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(actualPasswordHash), []byte(rawIncomingPassword))
	return err == nil
}
