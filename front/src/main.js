// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import * as firebase from 'firebase/app'
import 'firebase/auth'

Vue.config.productionTip = false

/* initiaze firebase */
const firebaseConfig = {
  apiKey: 'AIzaSyAMyBWbN24rQp29SyvqwztX7QZMEv9npgo',
  authDomain: 'echo-sns.firebaseapp.com',
  databaseURL: 'https://echo-sns.firebaseio.com',
  projectId: 'echo-sns',
  storageBucket: 'echo-sns.appspot.com',
  messagingSenderId: '18493509738',
  appId: '1:18493509738:web:f9553df5bf85c44a03ba13',
  measurementId: 'G-4XJTKFFLWM'
}
firebase.initializeApp(firebaseConfig)

// firebase初期化完了後にVueインスタンスを作成する
let app
firebase.auth().onAuthStateChanged(user => {
  /* eslint-disable no-new */
  if (!app) {
    new Vue({
      el: '#app',
      router,
      template: '<App/>',
      components: { App }
    })
  }
})
