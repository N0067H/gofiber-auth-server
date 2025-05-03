package db

import (
	"context"

	"github.com/n0067h/gofiber-auth-server/ent"
)

func Migrate(client *ent.Client) error {
	return client.Schema.Create(context.Background())
}
