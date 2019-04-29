<template>
  <div class="content">
    <v-container fluid grid-list-md>
      <v-layout row wrap>
        <v-flex d-flex xs1 class="centered">
          <div>
            <v-btn flat icon color="primary" v-on:click="previousComponent()">
              <span class="fa fa-chevron-left hover"></span>
            </v-btn>
          </div>
        </v-flex>
        <v-flex xs10 class="component">
          <router-view></router-view>
        </v-flex>
        <v-flex d-flex xs1 class="centered">
          <div>
            <v-btn flat icon color="primary" v-on:click="nextComponent()">
              <span class="fa fa-chevron-right hover"></span>
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
  justify-content: center;
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
