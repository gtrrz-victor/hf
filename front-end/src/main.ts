import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import axios from 'axios';
import VueCookies from 'vue-cookies'
const API_BASE_URL = process.env.API_BASE_URL || "http://192.168.1.104:3000"

Vue.prototype.$http = axios.create({baseURL: API_BASE_URL});
Vue.config.productionTip = false;
Vue.use(VueCookies)

new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount("#app");
