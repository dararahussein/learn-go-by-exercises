package main

import (
	"log"
	"net/http"

	"learn-go-by-exercises/set10/api"
	"learn-go-by-exercises/set10/task"
)

func main() {
	store, err := task.Load("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", api.NewHandler(store)))
}
