package database

import (
	"database/sql"
	"log"
	"os"
	"test/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env not found")
	}
	conf := config.NewConfigFromEnv()
	conn := conf.Config()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println("err conn db", err)
		return nil, ""
	}

	return db, port
}
