import Vue from 'vue'
import Router from 'vue-router'
import CardsControl from '@/components/CardsControl'
import Control from '@/components/Control'
import Grafico from '@/components/Grafico'
import MobileControl from '@/components/MobileControl'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'

Vue.use(Vuetify)

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'CardsControl',
      component: CardsControl
    },
    {
      path: '/controle',
      name: 'Control',
      component: Control
    },
    {
      path: '/mobile',
      name: 'Mobile',
      component: MobileControl
    },
    {
      path: '/grafico',
      name: 'Grafico',
      component: Grafico
    }
  ]
})
