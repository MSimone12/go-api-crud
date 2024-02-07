package main

import (
	api "go-api-crud/router"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
