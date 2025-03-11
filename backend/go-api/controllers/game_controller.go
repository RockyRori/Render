package controllers

import (
	"net/http"
	"render/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameController struct {
	DB *gorm.DB
}

func NewGameController(db *gorm.DB) *GameController {
	return &GameController{DB: db}
}

func (gc *GameController) GetGames(c *gin.Context) {
	var games []models.Game
	gc.DB.Find(&games)
	c.JSON(http.StatusOK, games)
}

func (gc *GameController) CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gc.DB.Create(&game)
	c.JSON(http.StatusCreated, game)
}
