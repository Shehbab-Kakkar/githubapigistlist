package main

import (
	"log"
	"net/http"
)

func main() {
	handler := NewServer()

	log.Println("Server is listening on :8080 and test on self hosted runner")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

