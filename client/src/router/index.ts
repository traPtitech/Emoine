import { createRouter, createWebHistory } from 'vue-router'
import Index from '/@/pages/Index.vue'
import Admin from '/@/pages/Admin.vue'
import Null from '/@/pages/Null.vue'
import { useStore } from '/@/store'

export const routerHistory = createWebHistory()

const router = createRouter({
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
    },
    {
      path: '/:catchAll(.*)',
      component: Null
    }
  ]
})

export default router

router.beforeEach(async (to, from, next) => {
  const store = useStore()

  if (store.state.me) {
    next(true)
    return
  }

  await store.dispatch.fetchMe()

  if (store.state.me) {
    next(true)
    return
  }

  location.href = '/api/oauth2/code'
})
