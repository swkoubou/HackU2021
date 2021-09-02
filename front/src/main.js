import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'

async function loadFirebaseConfig(url) {
  const response = await fetch(url)
  const firebaseConfig = await response.json()
  return firebaseConfig
}

async function initApp() {
  // Initialize Firebase
  const config = await loadFirebaseConfig('auth.json')
  firebase.initializeApp(config)
  firebase.auth().onAuthStateChanged(() => {
    // Initialize Vue.js
    new Vue({
      store,
      router,
      render: (h) => h(App),
    }).$mount('#app')
  })
}

initApp()
