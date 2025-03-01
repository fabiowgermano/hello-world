package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/health", handleHealthCheck)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"message\":\"Hello world!\"}")
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"status\":\"healthy\"}")
}
