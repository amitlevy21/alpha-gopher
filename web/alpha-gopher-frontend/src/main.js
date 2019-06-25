import Vue from 'vue'
import VueResource from "vue-resource"
import VueRouter from "vue-router"
import App from './App.vue'

Vue.config.productionTip = false
Vue.use(VueResource);
Vue.use(VueRouter);

new Vue({
  render: h => h(App),
}).$mount('#app')
