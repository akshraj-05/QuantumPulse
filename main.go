package main

import (
	"fmt"
	"html/template"
	"html/template"
	"latency/ping"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// If the request method is GET, render the HTML template
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, "Failed to parse HTML template", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Failed to render HTML template", http.StatusInternalServerError)
			return
		}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// If the request method is GET, render the HTML template
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, "Failed to parse HTML template", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Failed to render HTML template", http.StatusInternalServerError)
			return
		}
	}

	// If the request method is POST, process the form data
	if r.Method == http.MethodPost {
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
			return
		}

		// Retrieve the value of the "domain" field
		domain := r.Form.Get("domain")
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
			return
		}

		// Retrieve the value of the "domain" field
		domain := r.Form.Get("domain")

		lat, err := ping.Ping(domain)
		if err != nil {
			http.Error(w, "Failed to ping domain", http.StatusInternalServerError)
			return
		}
		// Use the value of "domain" as needed
		fmt.Fprintf(w, "Latency of domain %s is : %v", domain, lat)
		lat, err := ping.Ping(domain)
		if err != nil {
			http.Error(w, "Failed to ping domain", http.StatusInternalServerError)
			return
		}
		// Use the value of "domain" as needed
		fmt.Fprintf(w, "Latency of domain %s is : %v", domain, lat)
	}
}

func main() {
	r := mux.NewRouter()

	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define route for root path ("/")
	r.HandleFunc("/", indexHandler)

	// Start HTTP server
	// Start HTTP server
	fmt.Println("Server listening on port 80")
	if err := http.ListenAndServe(":80", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
	if err := http.ListenAndServe(":80", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
