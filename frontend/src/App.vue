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