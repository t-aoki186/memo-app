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

func helloHandler(w http.ResponseWriter, r *http.Request) { // w,r: HTTPレスポンスの書き込みと取得。 , *はポインタ
	response := map[string]string{ // map[string]string: 文字列のキーと値を持つマップを作成する。
		"message": "hello, world!",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
