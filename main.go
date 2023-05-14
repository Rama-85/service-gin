package main

import (
	"github.com/Rama-85/service-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()
	//router.Run()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8081")
}
