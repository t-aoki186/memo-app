package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "hello, world!",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
