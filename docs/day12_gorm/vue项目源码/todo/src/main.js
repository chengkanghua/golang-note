import Vue from 'vue'
import App from './App.vue'
import './plugins/element.js'
import router from './router'
import axios from './plugins/axios'

Vue.prototype.$axios = axios;
Vue.config.productionTip = true

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
