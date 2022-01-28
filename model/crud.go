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
		//Picks the Id form the url
		var foundItem views.Item
		for index, item := range data.Items {
			if item.ID == id {
				//checks if the item id exist and place it in fountItem
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
	//encodes the foundItem
	json.NewEncoder(w).Encode(foundItem)
}

func PostItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /items: PostItems")
	reqBody, _ := ioutil.ReadAll(r.Body)
	//takes the JSON payload
	x := bytes.TrimLeft(reqBody, " \t\r\n")
	isArray := len(x) > 0 && x[0] == '['
	//prepares to create an item
	itemsToCreate := make([]views.Item, 0)

	//checks if the payload is a single JSON or an array os JSONs
	if isArray {
		decoder := json.NewDecoder(bytes.NewBufferString(string(reqBody)))
		err := decoder.Decode(&itemsToCreate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for _, item := range itemsToCreate {
			//check items in the array validating the data
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
		//itemsToCreate becomes the single JSON payload
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
			//apends the data JSON with the new items
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
			//search for the id of the item
			if item.ID == id {
				//deletes that item
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
