package main

import (
	"log"
	"net/http"

	"github.com/DNahar74/golang-playlist/23.mongoDB_API/routers"
)

func main() {
	router := routers.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}