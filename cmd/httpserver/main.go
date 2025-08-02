package main

import (
	"log"
	"net/http"

	"github.com/quii/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := httpserver.NewHandler()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
