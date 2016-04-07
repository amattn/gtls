package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

const DEFAULT_LISTEN_ADDRESS = ":8080"

func main() {
	log.Println("gtls", Version(), "build", BuildNumber())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, you came from: %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening from:", DEFAULT_LISTEN_ADDRESS)
	log.Fatal(http.ListenAndServe(DEFAULT_LISTEN_ADDRESS, nil))
}
