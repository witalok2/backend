import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Listar from './views/Listar.vue'

Vue.use(Router)

export default new Router({
  routes: 
  [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/listar',
      name: 'listar',
      component: Listar
    }
  ]
})
