# Render 游戏商店 - 全栈项目

## 📌 项目简介
**Render 游戏商店** 是一个全栈开发项目，允许用户浏览、购买游戏，并支持简单的聊天机器人。

- **前端**：Vue 3 + TypeScript + Vite
- **后端**：Golang（Gin + GORM）+ Python（Flask）
- **数据库**：MySQL
- **部署**：Docker + Docker Compose

---
## 📂 目录结构
```bash
render-project/
├── backend/
│   ├── go-api/              # Golang 后端 API
│   │   ├── main.go          # 入口文件
│   │   ├── router/          # API 路由
│   │   ├── models/          # 数据模型
│   │   ├── controllers/     # 业务逻辑
│   │   ├── database/        # 数据库连接
│   │   ├── config/          # 配置管理
│   ├── python-worker/       # Python 聊天机器人后端
│   │   ├── app/             # Flask API
│   │   ├── chatbot.py       # 聊天逻辑
│   │   ├── requirements.txt # 依赖
├── frontend/                # Vue 3 + TypeScript 前端
│   ├── src/                 # 源代码
│   │   ├── components/      # 组件
│   │   ├── pages/           # 页面
│   │   ├── router/          # Vue Router
│   │   ├── api/             # API 调用封装
│   ├── package.json         # 依赖管理
│   ├── vite.config.ts       # Vite 配置
│   ├── Dockerfile           # 前端 Docker 配置
├── database/                # 数据库 SQL
│   ├── schema.sql           # 数据库结构
│   ├── seed.sql             # 初始数据填充
├── docker-compose.yml       # Docker Compose 配置
├── README.md                # 项目文档
```

---
## 🎯 功能
### **🖥️ 前端（Vue 3 + TypeScript）**
- 游戏列表展示
- 聊天机器人交互
- API 请求管理

### **🚀 后端（Golang + Python）**
- **Golang API**：游戏数据管理（CRUD）
- **Python Chatbot**：简单 AI 聊天功能

### **🗄️ 数据库（MySQL）**
- 游戏信息表
- 用户信息表
- 订单表

---
## ⚙️ **安装 & 运行**

### **1️⃣ 克隆项目**
```bash

```

### **2️⃣ 运行数据库**
```bash

```

### **3️⃣ 初始化数据库**
```bash

```

### **4️⃣ 运行后端（Golang）**
```bash
cd backend/go-api
go run main.go
```

### **5️⃣ 运行聊天机器人（Python）**
```bash
cd backend/python-worker
pip install -r requirements.txt
python app/main.py
```

### **6️⃣ 运行前端（Vue 3）**
```bash
cd frontend
npm install
npm run dev
```

### **7️⃣ 使用 Docker 运行全部服务**
```bash

```

---
## 📡 API 端点
### **🎮 游戏 API**（Golang）
| Method | Endpoint   | Description         |
|--------|-----------|---------------------|
| GET    | /games    | 获取所有游戏        |
| POST   | /games    | 添加新游戏          |

### **💬 聊天机器人 API**（Python）
| Method | Endpoint   | Description      |
|--------|-----------|------------------|
| POST   | /chat     | 发送聊天消息      |

---
## 🛠️ **技术栈**
- **前端**：Vue 3 + TypeScript + Vite + Pinia + Vue Router
- **后端**：Golang（Gin + GORM）+ Python（Flask）
- **数据库**：MySQL
- **部署**：Docker + Docker Compose

---
## 📜 **许可证**
AGPL License

**🎮 现在你可以运行 Render 游戏商店，并开始开发更多功能！🚀**
