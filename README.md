# Game Store Fullstack Project - Backend (Golang)

## 📂 目录结构
```bash
backend/
├── go-api/                 # Golang API（主服务器）
│   ├── main.go             # 入口文件
│   ├── router/             # API 路由
│   ├── models/             # 数据模型
│   ├── controllers/        # 业务逻辑
│   ├── services/           # 额外服务
│   ├── middleware/         # 中间件（认证、日志等）
│   ├── config/             # 配置文件
│   ├── database/           # 数据库连接
│   ├── tests/              # 单元测试
│   ├── go.mod              # Go 依赖管理
│   ├── Dockerfile          # Golang API Docker 配置
│   ├── .env                # 环境变量
```

## 1️⃣ **初始化 Golang 项目**
```bash
cd backend/go-api
go mod init game-store
```

### **安装依赖**
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

---
## 2️⃣ **主入口文件**
```go
// main.go
package main

import (
	"game-store/router"
	"game-store/config"
	"game-store/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := database.ConnectDB()
	r := gin.Default()
	router.SetupRoutes(r, db)

	r.Run(":8080")
}
```

---
## 3️⃣ **路由管理**
```go
// router/router.go
package router

import (
	"game-store/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	gameController := controllers.NewGameController(db)
	r.GET("/games", gameController.GetGames)
	r.POST("/games", gameController.CreateGame)
}
```

---
## 4️⃣ **数据库连接**
```go
// database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = db
	return db
}
```

---
## 5️⃣ **游戏数据模型（GORM）**
```go
// models/game.go
package models

type Game struct {
	ID    uint    `gorm:"primaryKey"`
	Title string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
}
```

---
## 6️⃣ **游戏控制器**
```go
// controllers/game_controller.go
package controllers

import (
	"game-store/models"
	"net/http"
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
```

---
## 7️⃣ **环境变量配置**
```ini
# .env
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root
DB_NAME=game_store
```

---
## 8️⃣ **Golang API Dockerfile**
```dockerfile
# Dockerfile
FROM golang:1.19

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]

EXPOSE 8080
```

---
## 9️⃣ **单元测试**
```go
// tests/game_test.go
package tests

import (
	"testing"
	"game-store/models"
)

func TestGameModel(t *testing.T) {
	game := models.Game{Title: "Test Game", Price: 29.99}
	if game.Title != "Test Game" {
		t.Errorf("Expected Title to be 'Test Game', got %s", game.Title)
	}
}
```

---
## ✅ **运行 API**
```bash
go run main.go
```

## ✅ **测试 API**
```bash
curl -X GET http://localhost:8080/games
```

## ✅ **使用 Docker 运行**
```bash
docker build -t game-store .
docker run -p 8080:8080 game-store
```

---
**🎮 现在你的 Golang 后端框架已经搭建好，支持 API、数据库、Docker、单元测试！** 🚀
