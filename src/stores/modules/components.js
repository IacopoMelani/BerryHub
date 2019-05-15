import router from "../../router";

export default {
	namespaced: true,
	state: {
		componentsList: ["weather", "calendar", "clock", "news"],
		currentComponent: "weather",
		updating: false,
		isMenu: false
	},
	mutations: {
		changeCurrentComponent(state, path) {
			state.currentComponent = path;
			state.updating = false;
			state.isMenu = false;
			if (path == "menu") {
				state.isMenu = true;
			}
			router.push({name: path});
		},
		changeCurrentComponentUpdating(state) {
			state.updating = true;
		}
	},
	actions: {
		goToComponent: function(context, component) {
			if (context.state.updating) {
				return;
			}
			context.commit("changeCurrentComponentUpdating");
			context.commit("changeCurrentComponent", component);
		},
		goToMenu: function(context) {
			if (context.state.updating) {
				return;
			}
			context.commit("changeCurrentComponentUpdating");
			context.commit("changeCurrentComponent", "menu");
		},
		nextComponent: function(context) {
			if (context.state.updating) {
				return;
			}
			context.commit("changeCurrentComponentUpdating");
			for (var i = 0; i < context.state.componentsList.length; i++) {
				if (context.state.currentComponent == context.state.componentsList[i]) {
					if (i == context.state.componentsList.length - 1) {
						context.commit("changeCurrentComponent", context.state.componentsList[0]);
					} else {
						context.commit("changeCurrentComponent", context.state.componentsList[i + 1]);
					}
					return;
				}
			}
		},
		previousComponent: function(context) {
			if (context.state.updating) {
				return;
			}
			context.commit("changeCurrentComponentUpdating");
			for (var i = 0; i < context.state.componentsList.length; i++) {
				if (context.state.currentComponent == context.state.componentsList[i]) {
					if (i == 0) {
						context.commit(
							"changeCurrentComponent",
							context.state.componentsList[context.state.componentsList.length - 1]
						);
					} else {
						context.commit("changeCurrentComponent", context.state.componentsList[i - 1]);
					}
					return;
				}
			}
		}
	}
};
