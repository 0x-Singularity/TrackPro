package main

import (
	"TrackPro/handlers" // Adjust to your module path
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	// Parse all HTML templates in frontend/pages directory
	tmpl = template.Must(template.ParseGlob("../frontend/pages/*.html"))
	handlers.SetTemplate(tmpl)
}

func main() {
	// Serve static files from the frontend directory
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes using handlers
	http.HandleFunc("/", handlers.HomeHandler)           // Home page
	http.HandleFunc("/about", handlers.AboutHandler)     // About page
	http.HandleFunc("/contact", handlers.ContactHandler) // Contact page

	// Starting the web server
	log.Println("Starting the web server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
