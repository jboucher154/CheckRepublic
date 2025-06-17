package main

import (
	"check_republic/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	

	// use multiplexer to match url patterns!

	// adds a new handler
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This should serve the root")
	})
	// can pass idectly to the listen and serve for root behavior, but can also add as a specific endpoint
	// handler := http.HandlerFunc(api.ChecklistHandler)
	http.HandleFunc("/checklist", api.GetChecklistHandler)
	//****look into this more for static assests
	// serve static files by setting up a file server directory
	// fs := http.FileServer(http.Dir("static/"))
	// point url to the file path with and handle function
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start up server
	log.Fatal(http.ListenAndServe(":8080", nil))
}