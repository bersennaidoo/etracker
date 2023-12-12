package main

import (
	"fmt"

	"github.com/bersennaidoo/etracker/server/infrastructure/storage/pgstore"
	"github.com/bersennaidoo/etracker/server/physical/config"
	"github.com/bersennaidoo/etracker/server/physical/conn"
)

func main() {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.GetAsString("DB_USER", "postgres"),
		config.GetAsString("DB_PASSWORD", "bersen"),
		config.GetAsString("DB_HOST", "localhost"),
		config.GetAsInt("DB_PORT", 5432),
		config.GetAsString("DB_NAME", "postgres"),
	)
	dbcl := conn.NewPGCON(dbURI)
	_ = pgstore.New(dbcl)

}
