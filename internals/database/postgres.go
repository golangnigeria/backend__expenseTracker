package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	conection, err := openDB(dsn)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")

	return conection, nil
}
