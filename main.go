package main

import (
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)

	const addr = "localhost:8080"

	http.ListenAndServe(addr, serveMux)
}
