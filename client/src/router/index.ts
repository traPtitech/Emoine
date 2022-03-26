import { createRouter, createWebHistory } from 'vue-router'

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
      path: '/overlay-viewer',
      name: 'overlay-viewer',
      component: () => import('/@/pages/OverlayViewer.vue')
    },
    {
      path: '/popup-comment-list',
      name: 'popup-comment-list',
      component: () => import('/@/pages/PopupCommentList.vue')
    },
    {
      path: '/:catchAll(.*)',
      component: () => import('/@/pages/Null.vue')
    }
  ]
})

export default router
