<template>
    <div>
      <h2>聊天机器人</h2>
      <input v-model="message" placeholder="输入消息..." />
      <button @click="sendMessageToBot">发送</button>
      <p>回复: {{ response }}</p>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import { sendMessage } from '@/api/chatApi';
  
  const message = ref('');
  const response = ref('');
  
  async function sendMessageToBot() {
    response.value = await sendMessage(message.value);
  }
  </script>