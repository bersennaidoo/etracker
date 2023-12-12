package conn

import (
	"database/sql"

	"github.com/bersennaidoo/etracker/server/infrastructure/logger"
	_ "github.com/lib/pq"
)

func NewPGCON(dbUri string) *sql.DB {
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		logger.Logger.Errorf("Error opening database: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		logger.Logger.Errorf("Error from database ping: %s", err.Error())
	}

	return db
}
