package main

import (
	"check_republic/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	// adds a new handler
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This should serve the root")
	})
	handler := http.HandlerFunc(api.ChecklistHandler)

	//****look into this more for static assests
	// serve static files by setting up a file server directory
	// fs := http.FileServer(http.Dir("static/"))
	// point url to the file path with and handle function
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start up server
	log.Fatal(http.ListenAndServe(":8080", handler))
}