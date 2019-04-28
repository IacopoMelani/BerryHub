import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import Weather from "./views/Weather.vue";

export default new Router({
	mode: "hash",
	routes: [{ name: "weather", path: "/weather", component: Weather}]
});
