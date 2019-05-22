import axios from "axios";

import utility from "./utility";

export default {
	namespaced: true,
	state: {
		updating: false,
		updatingErrMsg: "",
		counter: 0,
		data: {
			temp: Number,
			pressure: Number,
			humidity: Number,
			tempMin: Number,
			tempMax: Number,
			windSpeed: Number,
			city: "",
			icon: "",
			message: ""
		}
	},
	mutations: {
		incrementCounter: function(state) {
			state.counter++;
		},
		resetCounter: function(state) {
			state.counter = 0;
		},
		getWeatherDataError: function(state) {
			state.updating = false;
			state.updatingErrMsg = "Errore: richiesta scaduta";
		},
		getWeatherDataFailed: function(state, message) {
			state.updating = false;
			state.updatingErrMsg = message;
		},
		getWeatherDataSuccess: function(state, payload) {
			state.updating = false;
			state.data.temp = utility.roudOneDigit(payload.main.temp);
			state.data.pressure = payload.main.pressure;
			state.data.humidity = payload.main.humidity;
			state.data.tempMin = utility.roudOneDigit(payload.main.temp_min);
			state.data.tempMax = utility.roudOneDigit(payload.main.temp_max);
			state.data.windSpeed = Math.round(payload.wind.speed);
			state.data.city = payload.name;
			state.data.icon = payload.weather[0].icon + ".png";
			state.data.message = payload.weather[0].description;
		},
		getWeatherDataUpdating: function(state) {
			state.updating = true;
			state.updatingErrMsg = "";
		}
	},
	actions: {
		getWeatherData: function(context) {
			if (context.state.updating || context.state.counter >= 5) {
				return;
			}
			context.commit("getWeatherDataUpdating");
			context.commit("incrementCounter");

			axios
				.post("/weather/data")
				.then(result => {
					if (result.data && result.data.success && result.data.content) {
						context.commit("getWeatherDataSuccess", result.data.content);
					} else {
						context.commit("getWeatherDataFailed", result.data.message);
					}
				})
				.catch(() => {
					context.commit("getWeatherDataError");
				});
		}
	}
};
