import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import axios from 'axios'

const backendApi = 'http://localhost:3000/'

function checkTicket(id){
  const data = new URLSearchParams({
    ticketID: id,
  });
  axios
    .post(backendApi + "shop/get-ticket", data)
    .catch((error) => {
      if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        console.log("data"+error.response.data);
        console.log("status"+error.response.status);
        console.log(error.response.headers);
        sessionStorage.removeItem("ticketID")
      } else if (error.request) {
        // The request was made but no response was received
        // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
        // http.ClientRequest in node.js
        console.log("request", error.request);
        sessionStorage.removeItem("ticketID")
      } else {
        // Something happened in setting up the request that triggered an Error
        console.log('Error', error.message);
      }
      console.log(error.config);
      return true
    })
    return true
}

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    beforeEnter: () => {
      if(!sessionStorage.getItem('auth')){return 'register'}
      if(sessionStorage.getItem('ticketID') && checkTicket(sessionStorage.getItem('ticketID'))){return 'progress'}
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
      if(sessionStorage.getItem('cusID')){return true}
      if(sessionStorage.getItem('ticketID')&& checkTicket(sessionStorage.getItem('ticketID'))){return 'progress'}
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
      if(sessionStorage.getItem('ticketID') && checkTicket(sessionStorage.getItem('ticketID'))){return true}
      if(sessionStorage.getItem('cusID')){return 'callmech'}
      return true
    },
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/RegisterView.vue'),
    beforeEnter(){
      if(sessionStorage.getItem('auth')){return '/'}
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
