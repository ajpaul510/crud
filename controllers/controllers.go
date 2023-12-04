package controllers

import (
	"crud/initialize"
	"crud/models"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

func (_ Controllers) Start() {
	env := "dev"
	config := initialize.AppSettings{}
	model := models.Models{}
	connstr, port := config.Init(env)
	model.Start(connstr, port)

	router := gin.Default()

	// GET
	router.GET("/user", models.GetUser)

	// POST
	router.POST("/user", models.PostUser)

	router.Run("localhost:8080")
}
