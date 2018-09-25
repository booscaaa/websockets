import Vue from 'vue'
import Router from 'vue-router'
import Chat from '@/components/Chat'
import Start from '@/components/Start'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Chat',
      component: Chat
    },
    {
      path: '/chat-maquina-um',
      name: 'Chat',
      component: Chat
    },
    {
      path: '/start',
      name: 'Start',
      component: Start
    }
  ]
})
