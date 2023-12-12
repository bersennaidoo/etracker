package conn

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPGCON(dbUri string) *sql.DB {
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}

	return db
}
