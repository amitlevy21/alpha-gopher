import Vue from 'vue'
import VueResource from "vue-resource"
import VueRouter from "vue-router"
import App from './App.vue'
import Routes from './routes'
import BootstrapVue from 'bootstrap-vue'

Vue.config.productionTip = false
Vue.use(VueResource);
Vue.use(VueRouter);
Vue.use(BootstrapVue)


const router = new VueRouter({
  routes: Routes
})

new Vue({
  render: h => h(App),
  router: router
}).$mount('#app')
