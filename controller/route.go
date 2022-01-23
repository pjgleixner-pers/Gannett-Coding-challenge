package controller

import (
	"gannett.com/api.grocery/model"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	router := gin.Default()
	router.GET("/items", model.GetAlbums)
	router.GET("/items/:id", model.GetAlbumByID)
	router.POST("/items", model.PostAlbums)
	router.DELETE("/items/:id", model.DeleteAlbums)

	router.Run("localhost:8080")

	return router
}
