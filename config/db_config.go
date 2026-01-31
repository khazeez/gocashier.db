package config

import (
	"database/sql"
	"log"

	"gocashier.db/pkg"
)

func InitDb() (*sql.DB, error) {
	var psqlinfo = pkg.Load().DBURL

	db, err := sql.Open("postgress", psqlinfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success connect to database")
	}

	return db, nil

}

func CloseDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
