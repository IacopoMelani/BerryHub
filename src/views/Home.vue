<template>
  <div class="content">
    <v-container fluid grid-list-md>
      <v-layout
        row
        wrap
        v-touch="{
          left: () => nextComponent(),
          right: () => previousComponent(),
        }"
      >
        <v-flex d-flex xs1 class="centered">
          <div v-if="!isMenu">
            <v-btn flat icon color="primary" v-on:click="previousComponent()">
              <span class="fa fa-chevron-left hover"></span>
            </v-btn>
          </div>
        </v-flex>
        <v-flex xs10 class="component">
          <router-view></router-view>
        </v-flex>
        <v-flex d-flex xs1 class="centered">
          <div v-if="!isMenu">
            <v-btn flat icon color="primary" v-on:click="nextComponent()">
              <span class="fa fa-chevron-right hover"></span>
            </v-btn>
          </div>
        </v-flex>
        <v-flex d-flex xs12 class="centered">
          <div v-if="!isMenu">
            <v-btn flat icon color="primary" v-on:click="goToMenu()">
              <span class="fa fa-th hover"></span>
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
  height: 85vh;
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
  computed: {
    isMenu: function() {
      return this.$store.state.components.isMenu;
    }
  },
  methods: {
    goToMenu: function() {
      this.$store.dispatch("components/goToMenu");
    },
    nextComponent: function() {
      this.$store.dispatch("components/nextComponent");
    },
    previousComponent: function() {
      this.$store.dispatch("components/previousComponent");
    }
  }
};
</script>
