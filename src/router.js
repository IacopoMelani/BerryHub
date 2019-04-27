import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import App from "./App.vue";

export default new Router({
	mode: "hash",
	routes: [
		{
			path: "/foffo",
			component: App
		}
	]
});
