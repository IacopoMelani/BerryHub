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

new Vue({
	router,
	store,
	render: h => h(App)
}).$mount("#app");
