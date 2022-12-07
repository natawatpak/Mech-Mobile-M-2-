import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    beforeEnter: () => {
      if(sessionStorage.getItem('ticketID')){return 'progress'}
      return true
    },
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
    component: () => import('../views/CallMechView.vue'),
    beforeEnter: () => {
      if(sessionStorage.getItem('cusID')){return true}
      if(sessionStorage.getItem('ticketID')){return 'progress'}
      return true
    },
  },
  {
    path: '/callcar',
    name: 'callcar',
    component: () => import('../views/CallCarView.vue')
  },
  {
    path: '/progress',
    name: 'progress',
    component: () => import('../views/ProgressView.vue'),
    beforeEnter: () => {
      if(sessionStorage.getItem('ticketID')){return true}
      if(sessionStorage.getItem('cusID')){return 'callmech'}
      return true
    },
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
