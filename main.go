package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/go-auth/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDatabase()
}

func main() {
	// db := config.GetDBConnection()
	r := gin.Default()

	// For Check Connection Only
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
