import Vue from 'vue'
import Router from 'vue-router'
import * as firebase from 'firebase/app'
import 'firebase/auth'

import Top from '@/components/Top'
import Home from '@/components/Home'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'

import homeLayout from '@/components/layouts/homeLayout'
import simpleLayout from '@/components/layouts/simpleLayout'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '*',
      redirect: 'signin'
    },
    {
      path: '/',
      name: 'Top',
      component: simpleLayout(Top)
    },
    {
      path: '/home',
      name: 'Home',
      component: homeLayout(Home),
      meta: { requiresAuth: true }
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    }
  ]
})

router.beforeEach((to, from, next) => {
  const currentUser = firebase.auth().currentUser
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  if (currentUser && to.name === 'Top') {
    next('home')
  } else if (requiresAuth && !currentUser) {
    next('signin')
  } else {
    next()
  }
})

export default router
