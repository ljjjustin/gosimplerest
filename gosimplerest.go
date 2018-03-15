package main

import (
	"log"
	"net/http"

	"github.com/ljjjustin/gosimplerest/api"
)

func main() {
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
