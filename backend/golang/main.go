package main

import (
	"game-store/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.SetupRoutes(app)
	app.Run(":8080")
}
