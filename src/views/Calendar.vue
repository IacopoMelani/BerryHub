<template>
  <v-container fluid class="bgcolor-2 centered">
    <v-card class="bgcolor">
      <v-layout wrap>
        <v-flex xs12 class="mb-3">
          <v-sheet>
            <v-calendar
              v-model="start"
              ref="calendar"
              :type="type"
              locale="it"
              :short-months="false"
              :weekdays="[1, 2, 3, 4, 5, 6, 0]"
              color="primary"
            ></v-calendar>
          </v-sheet>
        </v-flex>

        <v-flex xs4 class="text-sm-left text-xs-center mb-3">
          <v-btn @click="$refs.calendar.prev()">
            <v-icon dark left>keyboard_arrow_left</v-icon>Precedente
          </v-btn>
        </v-flex>
        <v-flex xs4 class="text-sm-center text-xs-center mb-3">
          <v-btn @click="today()">
            <v-icon dark left>explore</v-icon>Oggi
          </v-btn>
        </v-flex>
        <v-flex xs4 class="text-sm-right text-xs-center mb-3">
          <v-btn @click="$refs.calendar.next()">
            Successivo
            <v-icon right dark>keyboard_arrow_right</v-icon>
          </v-btn>
        </v-flex>
      </v-layout>
    </v-card>
  </v-container>
</template>

<style scoped>
.centered {
  align-items: center;
  display: flex;
}
.my-event {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  border-radius: 2px;
  background-color: #1867c0;
  color: #ffffff;
  border: 1px solid #1867c0;
  width: 100%;
  font-size: 12px;
  padding: 3px;
  cursor: pointer;
  margin-bottom: 1px;
}
</style>

<script>
export default {
  data: function() {
    return {
      start: new Date().toISOString().split("T")[0],
      type: "month"
    };
  },
  methods: {
    today: function() {
      this.start = new Date().toISOString().split("T")[0];
    }
  }
};
</script>
