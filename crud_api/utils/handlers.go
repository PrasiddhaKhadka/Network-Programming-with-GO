package utils

import (
	"log"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

type Server struct {
	mux   *http.ServeMux
	error ErrorHandler
}

func NewServer() *Server {
	return &Server{
		mux: &http.ServeMux{},
		error: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, "Server Error", 500)
		},
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("AppHandler ServeHTTP")
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Handle(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}

func (s *Server) handle(method string, path string, handler AppHandler) {
	pattern := method + " " + path
	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			log.Printf("handler error: %v\n", err)
			s.error(w, r, err)
		}
	})
}

func (s *Server) Get(path string, handler AppHandler)    { s.handle("GET", path, handler) }
func (a *Server) Post(path string, handler AppHandler)   { a.handle("POST", path, handler) }
func (a *Server) Put(path string, handler AppHandler)    { a.handle("PUT", path, handler) }
func (a *Server) Patch(path string, handler AppHandler)  { a.handle("PATCH", path, handler) }
func (a *Server) Delete(path string, handler AppHandler) { a.handle("DELETE", path, handler) }
