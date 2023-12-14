package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Mux *mux.Router
	mwf []mux.MiddlewareFunc
}

func New(mwff ...mux.MiddlewareFunc) *Router {
	router := &Router{
		Mux: mux.NewRouter().StrictSlash(true),
		mwf: mwff,
	}
	return router
}

func (r *Router) AddRoute(path, method string, handler http.HandlerFunc) {
	subRouter := r.Mux.PathPrefix(path).Subrouter()
	subRouter.Use(r.mwf...)
	subRouter.HandleFunc("", handler).Methods(method)
	log.Printf("Added route: [%v] [%v]", path, method)
}
