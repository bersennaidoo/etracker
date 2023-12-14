package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/etracker/backend/application/rest/handler"
	"github.com/bersennaidoo/etracker/backend/application/rest/mid"
	"github.com/bersennaidoo/etracker/backend/application/rest/router"
	"github.com/bersennaidoo/etracker/backend/application/rest/server"
	"github.com/bersennaidoo/etracker/backend/infrastructure/logger"
	"github.com/bersennaidoo/etracker/backend/infrastructure/storage/pgstore"
	"github.com/bersennaidoo/etracker/backend/physical/config"
	"github.com/bersennaidoo/etracker/backend/physical/conn"
	"github.com/gorilla/mux"
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

	st := pgstore.New(dbcl)

	hd := handler.New(st)

	defaultMiddleware := []mux.MiddlewareFunc{
		mid.JSONMiddleware,
		mid.CORSMiddleware(config.GetAsSlice("CORS_WHITELIST",
			[]string{
				"http://localhost:3000",
				"http://0.0.0.0:3000",
			}, ","),
		),
	}

	router := router.New(defaultMiddleware...)
	router.AddRoute("/login", http.MethodPost, hd.HandleLogin)

	srv := server.New(config.GetAsInt("SERVER_PORT", 3000), router)

	logger.Logger.Info("Server Started on :3000")
	srv.MustStart()
}
