package main

import (
	"log"
	"net/http"
)

func main() {

	directory := "./images"

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, "8080")

	http.ListenAndServe(":8080", nil)
}
