package controller

import (
	"log"
	"net/http"

	"gannett.com/api.grocery/model"
	"github.com/gorilla/mux"
)

func Register() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/item", model.GetItems)
	router.HandleFunc("/item/{id}", model.GetItemByID)
	router.HandleFunc("/items", model.PostItems).Methods("POST")
	router.HandleFunc("/items/{id}", model.DeleteItems).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

	return router
}
