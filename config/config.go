package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

func NewConfigFromEnv() *Config {
	port := "5432"
	if parsed := os.Getenv("POSTGRES_PORT"); parsed != "" {
		port = parsed
	}

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}

	return &Config{
		Host:     host,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Port:     port,
	}
}

func (c *Config) Config() string {
	if c.User == "" || c.Database == "" || c.Password == "" {
		log.Fatal("required database parameters not found")
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.Database, c.Port)
}
