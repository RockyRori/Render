# Render Fullstack Project - Backend (Golang)

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
│   │   ├── config.go       # 读取环境变量
│   ├── database/           # 数据库连接
│   │   ├── schema.sql      # 数据库初始化 SQL
│   ├── tests/              # 单元测试
│   ├── go.mod              # Go 依赖管理
│   ├── Dockerfile          # Golang API Docker 配置
│   ├── .env                # 环境变量
```

---
## 1️⃣ **数据库初始化 SQL - Render**
```sql
-- schema.sql: 初始化数据库结构
CREATE DATABASE IF NOT EXISTS render;
USE render;

-- 游戏表
CREATE TABLE IF NOT EXISTS games (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    release_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 订单表
CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    game_id INT NOT NULL,
    status ENUM('pending', 'completed', 'cancelled') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);
```

---
## 2️⃣ **自动执行 SQL 脚本 - Golang**
```go
// database/init.go
package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	file, err := ioutil.ReadFile("database/schema.sql")
	if err != nil {
		log.Fatal("Failed to read schema.sql: ", err)
	}

	if _, err := db.Exec(string(file)); err != nil {
		log.Fatal("Failed to execute schema.sql: ", err)
	}

	fmt.Println("Database initialized successfully!")
}
```

---
## 3️⃣ **环境变量配置**
```ini
# .env
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root
DB_NAME=render
```

---
## ✅ **运行数据库初始化**
```bash
go run database/init.go
```

**🎮 现在 `schema.sql` 已经创建，包含 `games`、`users` 和 `orders` 表，并可以自动执行！** 🚀
