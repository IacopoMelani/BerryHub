<template>
  <div class="content">
    <v-container fluid grid-list-md>
      <v-layout row wrap>
        <v-flex d-flex xs1 class="centered">
          <div>
            <v-btn flat icon color="primary">
              <span class="fa fa-chevron-left hover" v-on:click="previousComponent()"></span>
            </v-btn>
          </div>
        </v-flex>
        <v-flex xs10 class="component">
          <router-view></router-view>
        </v-flex>
        <v-flex d-flex xs1 class="centered">
          <div>
            <v-btn flat icon color="primary">
              <span class="fa fa-chevron-right hover" v-on:click="nextComponent()"></span>
            </v-btn>
          </div>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<style scoped>
.centered {
  align-items: center;
  display: flex;
  justify-content: center;
}

.component {
  align-items: center;
  display: flex;
  height: 95vh;
  overflow-y: auto;
}

.content {
  display: flex;
  height: 100vh;
}

.hover {
  cursor: pointer;
}
</style>

<script>
export default {
  name: "Home",
  created() {
    this.$router.push("/weather");
  },
  methods: {
    nextComponent: function() {
      this.$store.dispatch("components/nextComponent");
    },
    previousComponent: function() {
      this.$store.dispatch("components/previousComponent");
    }
  }
};
</script>
