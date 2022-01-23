package model

import (
	"net/http"

	"gannett.com/api.grocery/data"
	"gannett.com/api.grocery/views"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Items)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range data.Items {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
	//var newItems controller.
	var newItems views.Item
	//TODO: make acept many inputs
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newItems); err != nil {
		return
	}

	// Add the new album to the slice.
	data.Items = append(data.Items, newItems)
	c.IndentedJSON(http.StatusCreated, newItems)
}

func DeleteAlbums(c *gin.Context) {
	id := c.Param("id")

	for i, a := range data.Items {
		if a.ID == id {
			data.Items = append(data.Items[:i], data.Items[i+1:]...) //I dont get how it works
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
