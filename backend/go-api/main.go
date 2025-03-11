package main

import (
	"render/config"
	"render/database"
	"render/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := database.ConnectDB()
	r := gin.Default()
	router.SetupRoutes(r, db)

	r.Run(":8080")
}
