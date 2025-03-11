<template>
  <div class="game-list">
    <h2>🎮 游戏列表</h2>
    <ul>
      <li v-for="game in games" :key="game.id">
        <h3>{{ game.title }}</h3>
        <p>{{ game.description }}</p>
        <p>💰 价格: ${{ game.price }}</p>
        <p>📅 发布日期: {{ game.release_date }}</p>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { fetchGames } from '../api/gameApi';
import type { Game } from '../api/gameApi';

const games = ref<Game[]>([]);

// onMounted(async () => {
//   games.value = await fetchGames();
// });

onMounted(async () => {
  try {
    const data = await fetchGames();
    console.log("获取到的游戏数据:", data);
    games.value = data;
  } catch (error) {
    console.error("获取游戏数据失败:", error);
  }
});
</script>

<style scoped>
.game-list {
  padding: 20px;
  background: #f4f4f4;
  border-radius: 8px;
}
ul {
  list-style: none;
  padding: 0;
}
li {
  background: white;
  margin: 10px 0;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}
h3 {
  margin: 0;
  color: #2c3e50;
}
p {
  margin: 5px 0;
}
</style>