package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// GET

	// POST
	router.POST("/authentication", authentication)
	router.Run("localhost:8080")
}

func authentication(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }
