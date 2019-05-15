import axios from "axios";

export default {
	namespaced: true,
	state: {
		counter: 0,
		list: [],
		updating: false,
		updatingErrMsg: ""
	},
	mutations: {
		getNewsListError: function(state) {
			state.updating = false;
			state.updatingErrMsg = "Errore: richiesta scaduta";
		},
		getNewsListFailed: function(state, message) {
			state.updating = false;
			state.updatingErrMsg = message;
		},
		getNewsListSuccess: function(state, payload) {
			state.updating = false;
			state.updatingErrMsg = "";
			state.list = payload;
		},
		getNewsListUpdating(state) {
			state.updating = true;
			state.updatingErrMsg = "";
		},
		incrementCounter: function(state) {
			state.counter++;
		},
		resetCounter: function(state) {
			state.counter = 0;
		}
	},
	actions: {
		getNewsList: function(context) {
			if (context.state.updating || context.state.counter >= 5) {
				return;
			}
			context.commit("getNewsListUpdating");
			context.commit("incrementCounter");
			axios
				.get("/news/data")
				.then(result => {
					if (result.data && result.data.success && result.data.content) {
						context.commit("getNewsListSuccess", result.data.content);
					} else {
						context.commit("getNewsListFailed", result.data.message);
					}
				})
				.catch(() => {
					context.commit("getNewsListError");
				});
		}
	}
};
