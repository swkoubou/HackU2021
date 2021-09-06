import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ComponentsPreviewPage from '../views/ComponentsPreviewPage.vue'
import AnswerTheQuestionPage from '../views/AnswerTheQuestionPage.vue'
import ScorePage from '../views/ScorePage.vue'
import LoginPage from '../views/LoginPage'
import LoginSuccessPreviewPage from '../views/LoginSuccessPreviewPage'

import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/components',
    name: 'ComponentsPreviewPage',
    component: ComponentsPreviewPage,
  },
  {
    path: '/answerthequestion',
    name: 'AnswerTheQuestionPage',
    component: AnswerTheQuestionPage,
  },
  {
    path: '/score',
    name: 'ScorePage',
    component: ScorePage,
  },
  {
    path: '/login',
    name: 'LoginPage',
    component: LoginPage,
  },
  {
    path: '/loginsuccesspreviewpage',
    name: 'LoginSuccessPreviewPage',
    component: LoginSuccessPreviewPage,
    meta: { requiresAuth: true },
  },
]

const router = new VueRouter({
  routes,
})

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth)
  if (requiresAuth) {
    firebase.auth().onAuthStateChanged(function (user) {
      if (user) {
        next()
      } else {
        next({
          path: '/login',
          query: { redirect: to.fullPath },
        })
      }
    })
  } else {
    next()
  }
})

export default router
