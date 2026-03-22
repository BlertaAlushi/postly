package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Postgresql() {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SSL_MODE"))
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Database not reachable:", err)
	}
	fmt.Println("Connected to PostgreSQL")
	DB = db
}
