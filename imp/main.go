package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request @:", r.Method)
	fmt.Println("Url Path @:", r.URL.Path)
	fmt.Println("ok")
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Get(path string, handle http.HandlerFunc) {
	pattern := "GET" + " " + path
	s.mux.HandleFunc(pattern, handle)
}
func main() {
	fmt.Println("hello")
	s := NewServer()

	s.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HOME")
	})

	s.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ABOUT")
	})

	s.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CONTACT")
	})

	http.ListenAndServe(":8080", s)
}
