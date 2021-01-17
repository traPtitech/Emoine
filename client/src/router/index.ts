import { createRouter, createWebHistory } from 'vue-router'
import { useStore } from '/@/store'

export const routerHistory = createWebHistory()

const router = createRouter({
  history: routerHistory,
  routes: [
    {
      path: '/',
      name: 'index',
      component: () => import('/@/pages/Index.vue')
    },
    {
      path: '/admin',
      name: 'admin',
      component: () => import('/@/pages/Admin.vue')
    },
    {
      path: '/:catchAll(.*)',
      component: () => import('/@/pages/Null.vue')
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
