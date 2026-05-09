package main

import (
	"net/http"
)

type Mymux struct {
	routes map[string]http.Handler
}

func NewMyMux() *Mymux {
	return &Mymux{
		routes: make(map[string]http.Handler),
	}
}

// Register a route
func (m *Mymux) Handle(path string, handler http.Handler) {
	m.routes = map[string]http.Handler{
		path: handler,
	}
}

// THIS is what makes MyMux an http.Handler
// Without this method — MyMux is useless to ListenAndServe
func (m *Mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m.routes[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	handler.ServeHTTP(w, r)
}
