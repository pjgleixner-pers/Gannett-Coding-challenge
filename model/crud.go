package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gannett.com/api.grocery/data"
	"gannett.com/api.grocery/views"
	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /items: GetItems")
	json.NewEncoder(w).Encode(data.Items)
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /item/{id}: GetItemByID")
	handlerChannel := make(chan views.Item)

	go func() {
		vars := mux.Vars(r)
		id := vars["id"]
		var foundItem views.Item
		for index, item := range data.Items {
			if item.ID == id {
				foundItem = data.Items[index]
				break
			}
		}
		handlerChannel <- foundItem
	}()

	foundItem := <-handlerChannel
	if foundItem.ID != "" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(foundItem)
}

func PostItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /items: PostItems")

}

func DeleteItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE /items/{id}: DeleteItems")
	handlerChannel := make(chan bool)

	go func() {
		vars := mux.Vars(r)
		id := vars["id"]
		var deletedItem bool = false
		for index, item := range data.Items {
			if item.ID == id {
				data.Items = append(data.Items[:index], data.Items[index+1:]...)
				deletedItem = true
				break
			}
		}
		handlerChannel <- deletedItem
	}()

	itemDeleted := <-handlerChannel
	if itemDeleted {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
