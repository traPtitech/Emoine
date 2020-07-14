import { createRouter, createWebHistory } from 'vue-router'
import Index from '/@/pages/Index.vue'
import Admin from '/@/pages/Admin.vue'

export const routerHistory = createWebHistory()

export default createRouter({
  history: routerHistory,
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin
    }
  ]
})
