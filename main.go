package main

import (
	"log"
	"net/http"

	"gannett.com/api.grocery/controller"
)

func main() {
	mux := controller.Register()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
