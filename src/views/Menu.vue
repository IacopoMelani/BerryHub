  <template>
  <v-container fluid class="bgcolor-2 centered">
    <v-layout row wrap justify-center>
      <v-flex xs3 v-for="item in itemsList" :key="item.routeName">
        <v-btn large flat icon :color="item.color" v-on:click="goToComponent(item.routeName)">
          <v-icon class="hover" x-large>{{item.icon}}</v-icon>
        </v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<style scoped>
</style>

<script>
export default {
  name: "Menu",
  data: function() {
    return {
      itemsList: [
        {
          color: "orange darken-2",
          icon: "wb_sunny",
          label: "Meteo",
          routeName: "weather"
        },
        {
          color: "green darken-2",
          icon: "calendar_today",
          label: "Calendario",
          routeName: "calendar"
        },
        {
          color: "indigo darken-2",
          icon: "access_time",
          label: "Orologio",
          routeName: "clock"
        },
        {
          color: "red darken-2",
          icon: "library_books",
          label: "Notizie",
          routeName: "news"
        }
      ]
    };
  },
  methods: {
    goToComponent: function(routeName) {
      this.$store.dispatch("components/goToComponent", routeName);
    }
  }
};
</script>
