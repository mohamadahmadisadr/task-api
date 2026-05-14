package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "user=mohamad dbname=taskapi sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	er := db.Ping()
	if er != nil {
		log.Fatal(er)
	}

	log.Println("connected to postgres")

	return db
}
