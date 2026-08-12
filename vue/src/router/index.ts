import { createRouter, createWebHistory } from 'vue-router'
import p_index from './pages/p_index.vue'
import P_list from './pages/p_list.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { name: 'Index', component: p_index, path: '/' },
    { name: 'list', component: P_list, path: '/list' },
  ],
})

export default router