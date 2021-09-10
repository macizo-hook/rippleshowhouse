package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/tip", showTip)
	mux.HandleFunc("/tip/create", createTip)

	// Instantiate a fileserver that serves our static assets.
	// Set up a mux handler func to all URL paths that match /static, and strip the prefix prior to the request reaching the fs
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}