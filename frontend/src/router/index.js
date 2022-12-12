import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import axios from 'axios'

const backendApi = 'http://localhost:3000/'

function checkActiveTicket(id){
  const data = new URLSearchParams({
    cusID: id,
  });
  return axios
    .post(backendApi + "customer/get-active-ticket", data)
    .then((response) => {
      console.log(response.data === null)
      if (response.data === null){
        return false
      }
      console.log(response.data.ticketID)
      sessionStorage.setItem("ticketID", response.data.ticketID)
      return true
    })

}

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    beforeEnter: () => {
      if(!sessionStorage.getItem('auth')){return 'register'}
      console.log(checkActiveTicket(sessionStorage.getItem('cusID')))
      if(checkActiveTicket(sessionStorage.getItem('cusID'))){return 'progress'}
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
      if(!sessionStorage.getItem('auth')){return 'register'}
      if(checkActiveTicket(sessionStorage.getItem('cusID'))){return 'progress'}
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
      if(!sessionStorage.getItem('auth')){return 'register'}
      if(checkActiveTicket(sessionStorage.getItem('cusID'))){return true}
      if(sessionStorage.getItem('cusID')){return '/'}
      return true
    },
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/RegisterView.vue'),
    beforeEnter(){
      if(sessionStorage.getItem('auth')){return '/'}
      return true
    }
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
