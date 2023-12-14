package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bersennaidoo/etracker/backend/application/rest/router"
)

type Server struct {
	port   string
	server http.Server
	router *router.Router
	wg     sync.WaitGroup
}

func New(port int, router *router.Router) *Server {
	return &Server{
		router: router,
		port:   fmt.Sprintf(":%d", port),
	}
}

// MustStart will start the server and if it cannot bind to the port
// it will exit with a fatal log message
func (c *Server) MustStart() error {
	// Create the HTTP Server
	c.server = http.Server{
		Addr:           fmt.Sprintf("0.0.0.0%s", c.port),
		Handler:        c.router.Mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   0 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	// Start the listener
	go func() {
		log.Printf("API server started at %v on http://%s", time.Now().Format(time.Stamp), c.server.Addr)
		serverErrors <- c.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		log.Printf("shutdown status shutdown started signal %v\n", sig)
		defer log.Printf("shutdown status shutdown complete signal %v\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := c.server.Shutdown(ctx); err != nil {
			c.server.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
