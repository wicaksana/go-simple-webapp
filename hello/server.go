package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handles equest to "/"
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go Web Server Landing Page</title>
			<style>
				body { font-family: sans-serif; text-align: center; padding-top: 50px; }
				h1 { color: #333; }
				p { color: #555; }
			</style>
		</head>
		<body>
			<h1>Welcome!</h1>
			<p>This page is served by a simple Go web server.</p>
			<p>Current request path: %s</p>
		</body>
		</html>
	`, r.URL.Path)
}

// Handles request to "/health"
func handlerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status": "ok"}`)
}

func main() {
	port := ":8900"
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/health", handlerHealth)
	log.Printf("Server starting on http://localhost:%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
