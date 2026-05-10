package utils

import (
	"log"
	"net/http"
)

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		log.Printf("[%s] %s\n", method, path)
		next(w, r)
	}
}
