package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/go-auth/config"
	"github.com/rafli-lutfi/go-auth/routes"
)

func init() {
	config.LoadEnv()
	config.ConnectDatabase()
}

func main() {
	db := config.GetDBConnection()
	r := gin.Default()

	routes.RunServer(db, r)

	r.Run()
}
