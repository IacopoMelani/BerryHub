import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import Menu from "./views/Menu.vue";

import Weather from "./views/Weather.vue";
import Calendar from "./views/Calendar.vue";
import Clock from "./views/Clock.vue";
import News from "./views/News.vue";

export default new Router({
	mode: "hash",
	routes: [
		{
			name: "menu",
			path: "/menu",
			component: Menu
		},
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
		},
		{
			name: "news",
			path: "/news",
			component: News
		}
	]
});
