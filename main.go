package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("[v3] Hello, Kubernetes from host: %s!", host))
}

func main() {
	host, _ := os.Hostname()

	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	fmt.Printf("Server is running on port 8080 from %s", host)
	http.ListenAndServe(":8080", nil)
}
