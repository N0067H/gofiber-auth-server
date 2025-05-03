package db

import (
	"fmt"

	"github.com/n0067h/gofiber-auth-server/ent"
	"github.com/n0067h/gofiber-auth-server/internal/config"

	_ "github.com/lib/pq"
)

var client *ent.Client

func Init(cfg *config.Config) error {
	var err error
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s", cfg.DBUser, cfg.DBName, cfg.DBPass)
	client, err = ent.Open("postgres", dsn)

	return err
}

func Get() *ent.Client {
	return client
}

func Close() error {
	return client.Close()
}
