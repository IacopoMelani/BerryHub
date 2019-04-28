import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./stores/store";

Vue.config.productionTip = false;

// VUE-AXIOS
import axios from "axios";
import VueAxios from "vue-axios";
Vue.use(VueAxios, axios);

axios.defaults.withCredentials = true;

// VUETIFY
import Vuetify from "vuetify";
import "vuetify/dist/vuetify.min.css";
Vue.use(Vuetify);

new Vue({
	router,
	store,
	render: h => h(App)
}).$mount("#app");
