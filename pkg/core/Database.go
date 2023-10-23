package core

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return db, nil
}
