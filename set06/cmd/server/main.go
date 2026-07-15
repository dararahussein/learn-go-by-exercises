package main

import (
	"log"
	"net/http"

	"learn-go-by-exercises/set06"
)

func main() {
	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", set06.NewMux()))
}
