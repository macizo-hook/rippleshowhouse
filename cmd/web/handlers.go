package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"html/template"
)

// Define a handler function for home and offer a simple response body.
func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Read the files and store the templates in a set. Use a variadic parameter that consists of a slice of file paths.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Write the template content as the response body.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// showTip handler function
func showTip(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// We shouldn't have a negative int for id
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Interpolate id with the response and write it to the http.ResponseWriter
	fmt.Fprintf(w, "Display a specific tip with ID %d...", id)
}

// createTip handler function
func createTip(w http.ResponseWriter, r *http.Request){
	// We should check if the request is using POST, and if it isn't, send a 405
	if r.Method != http.MethodPost{
		// We should let folks know that POST methods are Allowed before nerfing them
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a Tip.."))
}
