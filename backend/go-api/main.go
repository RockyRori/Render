// main.go
package main

import (
	"render/database"
	"render/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 连接数据库
	db := database.ConnectDB() // 确保 database.ConnectDB() 返回 *gorm.DB

	// 允许所有来源访问
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有域访问
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // 不需要携带 Cookie 或认证信息
		MaxAge:           12 * time.Hour,
	}))

	// 传递 `r` 和 `db` 给路由
	router.SetupRoutes(r, db)

	r.Run(":8080") // 运行服务器
}
