package controller

import (
	"log"
	"net/http"

	"gannett.com/api.grocery/model"
	"github.com/gorilla/mux"
)

func Register() *mux.Router {

	/*
		HTTP request multiplexer". Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
	*/
	router := mux.NewRouter().StrictSlash(true)
	// registers a new route with a matcher for the URL path
	router.HandleFunc("/item", model.GetItems)
	router.HandleFunc("/item/{id}", model.GetItemByID)
	router.HandleFunc("/items", model.PostItems).Methods("POST")
	router.HandleFunc("/items/{id}", model.DeleteItems).Methods("DELETE")

	//Fatal is equivalent to Print() followed by a call to os.Exit(1)
	log.Fatal(http.ListenAndServe(":8080", router))

	return router
}
