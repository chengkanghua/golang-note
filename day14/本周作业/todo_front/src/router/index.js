import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login'
import TodoList from '@/components/TodoList'
import Register from '@/components/Register'

Vue.use(Router)

export default new Router({
  linkActiveClass:'is-active',
  mode:'history',
  routes: [
    {
      path: '/',
      redirect:"Login"
      // name: 'app',
      // component: HelloWorld
    },
    {
      path:"/register",
      name:'Register',
      component:Register   // 这里不要加分号
    },
    {
      path:"/login",
      name:'Login',
      component:Login   // 这里不要加分号
    },
    {
      path:"/todolist",
      name:'todolist',
      component:TodoList
    }
  ]
})
