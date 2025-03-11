# Render

game-store/
│── backend/                 # 后端（API 和业务逻辑）
│   ├── go-api/              # Golang API（主要负责高并发业务）
│   │   ├── main.go          # 主入口文件
│   │   ├── router/          # 路由
│   │   ├── models/          # 数据模型
│   │   ├── controllers/     # 控制器逻辑
│   │   ├── services/        # 业务逻辑
│   │   ├── middleware/      # 中间件（认证、日志等）
│   │   ├── config/          # 配置文件
│   │   ├── tests/           # 测试
│   │   ├── go.mod           # Golang 依赖管理
│   │   ├── go.sum           # Golang 依赖锁定文件
│   │   ├── Dockerfile       # Golang 后端 Docker 配置
│   │   ├── .env             # 环境变量
│   ├── python-worker/       # Python 异步任务（数据分析、推荐系统）
│   │   ├── main.py          # Python 任务入口
│   │   ├── tasks/           # 任务处理
│   │   ├── models/          # 数据模型
│   │   ├── utils/           # 工具类
│   │   ├── requirements.txt # Python 依赖管理
│   │   ├── Dockerfile       # Python 后端 Docker 配置
│── frontend/                # 前端（Vue 3）
│   ├── src/                 # 源代码
│   │   ├── components/      # Vue 组件
│   │   ├── pages/           # 页面
│   │   ├── router/          # 路由
│   │   ├── store/           # 状态管理（Pinia/Vuex）
│   │   ├── assets/          # 静态资源
│   │   ├── api/             # API 请求封装
│   │   ├── App.vue          # Vue 主组件
│   │   ├── main.js          # 入口文件
│   ├── public/              # 公共资源
│   ├── package.json         # 依赖管理
│   ├── vite.config.js       # Vite 配置
│   ├── Dockerfile           # Vue Docker 配置
│── database/                # 数据库
│   ├── migrations/          # 数据库迁移文件
│   ├── seeders/             # 数据填充
│   ├── schema.sql           # 数据库表结构（SQL）
│   ├── config.json          # 数据库配置
│── docs/                    # 项目文档
│   ├── api-docs.md          # API 说明
│   ├── system-design.md     # 系统设计文档
│── tests/                   # 测试代码
│   ├── backend/             # 后端测试
│   ├── frontend/            # 前端测试
│── scripts/                 # 部署和运维脚本
│   ├── deploy.sh            # 自动部署脚本
│   ├── setup.sh             # 服务器环境初始化
│── .gitignore               # Git 忽略规则
│── docker-compose.yml       # Docker Compose 文件（容器化部署）
│── README.md                # 项目介绍
