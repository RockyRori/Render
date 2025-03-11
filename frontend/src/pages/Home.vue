<template>
  <div>
    <h1>欢迎来到 Render 游戏商店</h1>
    <GameList />
    <ChatBot />
  </div>
</template>

<script setup lang="ts">
import GameList from '../components/GameList.vue';
import ChatBot from '../components/ChatBot.vue';
</script>