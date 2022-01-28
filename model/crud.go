package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gannett.com/api.grocery/data"
	"gannett.com/api.grocery/views"
	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /item: GetItems")
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
	reqBody, _ := ioutil.ReadAll(r.Body)
	x := bytes.TrimLeft(reqBody, " \t\r\n")
	isArray := len(x) > 0 && x[0] == '['

	itemsToCreate := make([]views.Item, 0)

	if isArray {
		decoder := json.NewDecoder(bytes.NewBufferString(string(reqBody)))
		err := decoder.Decode(&itemsToCreate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for _, item := range itemsToCreate {
			err := views.Validate(r, item)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	} else {
		var createdItem views.Item
		json.Unmarshal(reqBody, &createdItem)
		err := views.Validate(r, createdItem)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		itemsToCreate = append(itemsToCreate, createdItem)
	}

	handlerChannel := make(chan []views.Item)
	go func(items []views.Item) {
		for _, item := range items {
			//verifie if ID exist
			//Runs a time out but does the verification
			/* for i := range data.Items {
				if item.ID == data.Items[i].ID {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			} */
			data.Items = append(data.Items, item)
		}
		handlerChannel <- items
	}(itemsToCreate)

	itemsCreated := <-handlerChannel
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(itemsCreated)
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
