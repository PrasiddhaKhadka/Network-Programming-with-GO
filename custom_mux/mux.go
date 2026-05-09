package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := NewMyMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home page")
	}))

	// pre bult mux
	// pmux := http.NewServeMux()
	// pmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// })

	http.ListenAndServe(":8080", mux)
}
