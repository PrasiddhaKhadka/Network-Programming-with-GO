package main

import (
	"fmt"
	"net/http"
	Utils "prasiddhakhadka/np/pro/utils"
)

func main() {

	server := Utils.NewServer()

	server.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintf(w, "Hello loading......")
		return nil
	})
	server.Get("/about", func(w http.ResponseWriter, r *http.Request) error {
		// fmt.Fprintf(w, "about page")
		return fmt.Errorf("something went wrong!")
	})

	server.Get("/contact", func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(w, "contact page")
		return nil

	})

	http.ListenAndServe(":8080", server)

}
