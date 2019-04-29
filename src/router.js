import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import Weather from "./views/Weather.vue";
import Calendar from "./views/Calendar.vue";
import Clock from "./views/Clock.vue";

export default new Router({
	mode: "hash",
	routes: [
		{
			name: "weather",
			path: "/weather",
			component: Weather
		},
		{
			name: "calendar",
			path: "/calendar",
			component: Calendar
		},
		{
			name: "clock",
			path: "/clock",
			component: Clock
		}
	]
});
