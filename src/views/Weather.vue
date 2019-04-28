<template>
  <v-container fluid class="bgcolor-2 centered">
    <v-layout row wrap justify-center>
      <v-flex xs12>
        <v-card class="bgcolor">
          <v-card-title class="headline pb-0">{{currentCityName}}</v-card-title>
          <v-subheader>{{formatDate(new Date())}}, {{currentMessage}}</v-subheader>
          <v-layout row wrap align-center text-xs-center>
            <v-flex xs6>
              <v-card-text class="display-3">{{currentTemp}}&deg;C</v-card-text>
              <h3>({{currentTempMin}}&deg;C/{{currentTempMax}}&deg;C)</h3>
            </v-flex>
            <v-flex xs6>
              <img class="weather-icon" :src="currentIncon">
            </v-flex>
          </v-layout>
          <v-card-text>
            <v-icon>send</v-icon>
            <span class="pl-3">{{currentWindSpeed}} km/h</span>
          </v-card-text>
          <v-card-text>
            <v-icon>cloud_download</v-icon>
            <span class="pl-3">{{currentHumidity}}%</span>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<style scoped>
.centered {
  align-items: center;
  display: flex;
}
.weather-icon {
  height: 100px;
  width: 100px;
}
</style>

<script>
export default {
  computed: {
    currentCityName() {
      return this.$store.state.weather.data.city;
    },
    currentHumidity() {
      return this.$store.state.weather.data.humidity;
    },
    currentIncon() {
      if (this.$store.state.weather.data.icon != "") {
        return "/weather-icons/" + this.$store.state.weather.data.icon;
      }
    },
    currentMessage() {
      return this.$store.state.weather.data.message;
    },
    currentTemp() {
      return this.$store.state.weather.data.temp;
    },
    currentTempMin() {
      return this.$store.state.weather.data.tempMin;
    },
    currentTempMax() {
      return this.$store.state.weather.data.tempMax;
    },
    currentWindSpeed() {
      return this.$store.state.weather.data.windSpeed;
    }
  },
  data() {
    return {
      time: 0,
      getWeatherDataId: undefined
    };
  },
  methods: {
    formatDate: function(date) {
      var monthNames = [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December"
      ];

      var day = date.getDate();
      var monthIndex = date.getMonth();
      var year = date.getFullYear();

      return day + " " + monthNames[monthIndex] + " " + year;
    }
  },
  mounted: function() {
    this.$store.dispatch("weather/getWeatherData");

    this.getWeatherDataId = setInterval(() => {
      this.$store.dispatch("weather/getWeatherData");
    }, 60000);
  },
  destroyed: function() {
    clearInterval(this.getWeatherDataId);
  }
};
</script>

