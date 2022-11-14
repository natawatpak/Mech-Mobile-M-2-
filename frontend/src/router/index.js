import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoadingView from '../views/LoadingView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/callmech',
    name: 'callmech',
    component: () => import('../views/CallMechView.vue')
  },
  {
    path: '/callcar',
    name: 'callcar',
    component: () => import('../views/CallCarView.vue')
  },
  {
    path: '/loading',
    name: 'loading',
    component: LoadingView
  },
  {
    path: '/progress',
    name: 'progress',
    component: () => import('../views/ProgressView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
