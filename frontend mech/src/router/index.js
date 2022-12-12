import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/order',
    name: 'order',
    component: () => import('../views/OrderView.vue')
  },
  {
    path: '/details',
    name: 'details',
    component: () => import('../views/DetailView.vue')
  },
  {
    path: '/example',
    name: 'examples',
    component: () => import('../views/ExampleView.vue')
  },
  {
    path: '/example2',
    name: 'examples2',
    component: () => import('../views/ExampleView2.vue')
  },
  {
    path: '/history',
    name: 'history',
    component: () => import('../views/HistoryView.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/RegisterView.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'error404',
    component: () => import('../views/Error404View.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
