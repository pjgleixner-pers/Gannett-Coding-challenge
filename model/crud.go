package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gannett.com/api.grocery/data"
	"gannett.com/api.grocery/views"
	"github.com/gin-gonic/gin"
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

func PostItems(c *gin.Context) {
	//var newItems controller.
	var newItems views.Item
	//TODO: make acept many inputs
	// Call BindJSON to bind the received JSON to
	// newItem.
	if err := c.BindJSON(&newItems); err != nil {
		return
	}

	// Add the new Item to the slice.
	data.Items = append(data.Items, newItems)
	c.IndentedJSON(http.StatusCreated, newItems)
}

func DeleteItems(c *gin.Context) {
	id := c.Param("id")

	for i, a := range data.Items {
		if a.ID == id {
			data.Items = append(data.Items[:i], data.Items[i+1:]...) //I dont get how it works
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
