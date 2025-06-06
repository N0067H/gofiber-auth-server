package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

var cfg *Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	cfg = &Config{
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func Get() *Config {
	return cfg
}
