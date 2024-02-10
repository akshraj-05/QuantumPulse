package main

import (
	"fmt"
	"latency/ping"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// If the request method is GET, display the form
	if r.Method == "GET" {
		// Display the HTML form
		fmt.Fprintf(w, `
            <html>
            <body>
                <form method="POST" action="/">
                    <label for="">Enter Domain name:</label><br>
                    <input type="text" id="domain" name="domain"><br>
                    <input type="submit" value="Submit">
                </form>
            </body>
            </html>
        `)
		return
	}

	// If the request method is POST, process the form data
	if r.Method == "POST" {
		// Parse the form data
		r.ParseForm()
		domain1 := r.Form.Get("domain")

		domain, err := ping.Ping(string(domain1))
		if err == nil {
			fmt.Fprintln(w, "This is the latency", domain, "ms")
		}
	}
}

func main() {
	r := mux.NewRouter()

	// Define route for "/" path
	r.HandleFunc("/", helloHandler)

	// Start the server
	fmt.Println("Server listening on port 80")
	http.ListenAndServe(":80", r)
}
