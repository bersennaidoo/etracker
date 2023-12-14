package main

import (
	"flag"
	"fmt"

	"github.com/bersennaidoo/etracker/backend/application/rest/router"
	"github.com/bersennaidoo/etracker/backend/application/rest/server"
	"github.com/bersennaidoo/etracker/backend/infrastructure/logger"
	"github.com/bersennaidoo/etracker/backend/infrastructure/storage/pgstore"
	"github.com/bersennaidoo/etracker/backend/physical/config"
	"github.com/bersennaidoo/etracker/backend/physical/conn"
)

const serviceName = "etracker"

func main() {
	l := flag.Bool("local", false, "true - send to stdout, false - send to logging server")
	flag.Parse()

	/*sTracing, err := t.InitTracing(serviceName)
	if err != nil {
		log.Fatalf("Failed to setup tracing: %v\n", err)
	}
	defer func() {
		if err := sTracing(context.Background()); err != nil {
			log.Printf("Failed to shutdown tracing: %v\n", err)
		}
	}()

	/*sMetrics, err := m.InitMetrics(serviceName)
	if err != nil {
		log.Fatalf("Failed to setup metrics: %v\n", err)
	}
	defer func() {
		if err := sMetrics(context.Background()); err != nil {
			log.Printf("Failed to shutdown metrics: %v\n", err)
		}
	}()*/

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

	router := router.New()
	srv := server.New(config.GetAsInt("SERVER_PORT", 3000), router)

	logger.Logger.Info("Server Started on :3000")
	srv.MustStart()

}
