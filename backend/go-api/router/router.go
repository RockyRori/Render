package router

import (
	"render/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	gameController := controllers.NewGameController(db)
	r.GET("/games", gameController.GetGames)
	r.POST("/games", gameController.CreateGame)
}
