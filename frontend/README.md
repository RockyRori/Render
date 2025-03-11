# Render Fullstack Project - Frontend (Vue 3 + TypeScript)

## 📂 目录结构
```bash
frontend/
├── src/                   # Vue 源代码
│   ├── components/        # Vue 组件
│   │   ├── GameList.vue   # 游戏列表组件
│   │   ├── ChatBot.vue    # 聊天机器人组件
│   ├── pages/             # 页面
│   │   ├── Home.vue       # 主页
│   ├── router/            # Vue Router 配置
│   │   ├── index.ts       # 路由管理
│   ├── store/             # 状态管理（Pinia）
│   ├── api/               # API 请求封装
│   │   ├── gameApi.ts     # 游戏 API
│   │   ├── chatApi.ts     # 聊天 API
│   ├── App.vue            # Vue 主组件
│   ├── main.ts            # 入口文件
├── public/                # 静态资源
├── package.json           # 依赖管理
├── tsconfig.json          # TypeScript 配置
├── vite.config.ts         # Vite 配置
├── Dockerfile             # Vue Docker 配置
```

---
## 1️⃣ **App.vue - 主应用入口**
```vue
<!-- src/App.vue -->
<template>
  <div id="app">
    <header>
      <h1>🎮 Render 游戏商店</h1>
      <nav>
        <router-link to="/">主页</router-link>
      </nav>
    </header>
    <main>
      <router-view />
    </main>
    <footer>
      <p>© 2024 Render 游戏商店. All rights reserved.</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}

header {
  background-color: #34495e;
  padding: 10px;
  color: white;
}

nav a {
  color: white;
  text-decoration: none;
  margin: 0 15px;
  font-size: 18px;
}

nav a:hover {
  text-decoration: underline;
}

main {
  padding: 20px;
}

footer {
  margin-top: 30px;
  padding: 10px;
  background-color: #2c3e50;
  color: white;
}
</style>
```

---
## ✅ **运行 Vue 前端**
```bash
npm run dev
```

## ✅ **使用 Docker 运行**
```bash
docker build -t render-frontend .
docker run -p 3000:3000 render-frontend
```

---
**🎮 现在你的 `App.vue` 已经完成，包含导航栏、主页、动态路由显示和基本样式！🚀**
