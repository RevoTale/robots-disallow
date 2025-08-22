package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Define the robots.txt content
	robots := `
User-agent: *
Disallow: /`

	// Handler that always returns robots.txt for root and other paths
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Special endpoint for version/status info
		if r.URL.Path == "/status" {
			handleStatus(w, r)
			return
		}
		
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

// handleStatus provides a status endpoint demonstrating repository editing capability
func handleStatus(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"service":    "robots-disallow",
		"version":    "1.0.0",
		"timestamp":  time.Now().UTC(),
		"message":    "Repository editing capability demonstrated",
		"endpoints": map[string]string{
			"/":       "Returns robots.txt with 'User-agent: * Disallow: /'",
			"/status": "Returns this status information",
		},
	}
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(status)
}
