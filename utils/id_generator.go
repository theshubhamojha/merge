package utils

import (
	"context"

	"github.com/rs/xid"
)

func GenerateRandomID(ctx context.Context, prefix string) (id string) {
	guid := xid.New()
	id = prefix + guid.String()
	return id
}
