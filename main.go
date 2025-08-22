package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the robots.txt content
	robots := `
User-agent: *
Disallow: /`

	// Handler that always returns robots.txt
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, robots)
	})

	// Start server on port 80
	fmt.Println("Serving robots.txt on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
