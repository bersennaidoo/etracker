package main

import (
	"flag"
	"fmt"

	"github.com/bersennaidoo/etracker/backend/infrastructure/logger"
	"github.com/bersennaidoo/etracker/backend/infrastructure/storage/pgstore"
	"github.com/bersennaidoo/etracker/backend/physical/config"
	"github.com/bersennaidoo/etracker/backend/physical/conn"
)

func main() {
	l := flag.Bool("local", false, "true - send to stdout, false - send to logging server")
	flag.Parse()

	logger.SetLoggingOutput(*l)

	logger.Logger.Debugf("Application logging to stdout = %v", *l)
	logger.Logger.Info("Starting the application...")

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.GetAsString("DB_USER", "postgres"),
		config.GetAsString("DB_PASSWORD", "bersen"),
		config.GetAsString("DB_HOST", "localhost"),
		config.GetAsInt("DB_PORT", 5432),
		config.GetAsString("DB_NAME", "postgres"),
	)

	dbcl := conn.NewPGCON(dbURI)

	logger.Logger.Info("Database connection fine")

	_ = pgstore.New(dbcl)

}
