package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgres() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set in environment variables")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to Postgres: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Cannot ping Postgres: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully!")
	return db
}
