package router

import "github.com/gorilla/mux"

type Router struct {
	Mux *mux.Router
}

func New() *Router {
	return &Router{
		Mux: mux.NewRouter().StrictSlash(true),
	}
}
