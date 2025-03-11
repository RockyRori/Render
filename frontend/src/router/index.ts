import { createRouter, createWebHistory} from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
// 假设 Home.vue 实际路径为 ./pages/Home.vue，根据实际情况修改
import Home from '@/pages/Home.vue';

const routes: RouteRecordRaw[] = [
  { path: '/', component: Home }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;