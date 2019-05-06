<template>
  <v-app>
    <v-container fluid class="centered component">
      <v-layout wrap>
        <v-flex xs12 class="mb-3">
          <v-card>
            <v-time-picker :full-width="true" :use-seconds="true" :readonly="true" v-model="time"></v-time-picker>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<style scoped>
.centered {
  align-items: center;
  display: flex;
}
</style>

<script>
export default {
  data: () => ({
    time: new Date().toTimeString().split(" ")[0],
    timeId: undefined
  }),
  mounted: function() {
    this.time = new Date().toTimeString().split(" ")[0];

    this.timeId = setInterval(() => {
      this.time = new Date().toTimeString().split(" ")[0];
    }, 1000);
  },
  destroyed: function() {
    clearInterval(this.timeId);
  }
};
</script>