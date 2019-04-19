package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "こんにちは世界")
}

func main() {
	// Fake health check
	if len(os.Args) > 1 && os.Args[1] == "healthz" {
		os.Exit(0)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
