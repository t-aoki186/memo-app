package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("Server is running on *:8080")
	http.ListenAndServe(":8080", nil) // 指定なし, [net/http]のデフォルト処理を使用する。
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Hello, Go!",
	}

	json.NewEncoder(w).Encode(response)
}
