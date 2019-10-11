import Vue from 'vue'
import Router from 'vue-router'
import DefaultLayout from './layouts/Default.vue'
import Lista from './views/Lista.vue'
import Inicio from './views/Inicio.vue'
import Criar from './views/Criar.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {
          path: '',
          name: 'inicio',
          component: Inicio
        },
        {
          path: '/criar',
          name: 'criar',
          component: Criar
        },
        {
          path: '/lista',
          name: 'lista',
          component: Lista
        }
      ]
    }
  ]
})
