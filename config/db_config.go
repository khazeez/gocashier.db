package config

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gocashier.db/pkg"
)

func InitDb() (*sql.DB, error) {
	dsn:= pkg.Load().DBURL
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("success connect to database")
	return db, nil
}

func CloseDb(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Println("error closing database:", err)
	}
}
