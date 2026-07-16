package main

import (
	"log"
	"net/http"

	server "learn-go-by-exercises/set06/01_server"
)

func main() {
	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server.NewMux()))
}
