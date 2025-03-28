package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	// adds a new handler
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This should serve the root")
	})

	// serve static files by setting up a file server directory
	fs := http.FileServer(http.Dir("static/"))
	// point url to the file path with and handle function
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start up server
	http.ListenAndServe(":8080", nil)
}