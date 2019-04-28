import Vue from "vue";
import Vuex from "vuex";

// Modules
import components from "./modules/components";
import weather from "./modules/weather";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {components, weather},
	state: {},
	mutations: {},
	actions: {}
});
