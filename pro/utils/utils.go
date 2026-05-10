package utils

import (
	"fmt"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

type Server struct {
	mux          *http.ServeMux
	errorHandler func(w http.ResponseWriter, r *http.Request, err error)
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
		errorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("request received")

	s.mux.ServeHTTP(w, r)

}

func (s *Server) Get(path string, handler AppHandler) {

	// REGISTER
	s.mux.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			s.errorHandler(w, r, err)
		}
	})
}
