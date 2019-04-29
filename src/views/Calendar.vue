<template>
  <v-container fluid class="bgcolor-2 centered">
    <v-layout wrap>
      <v-flex xs12 class="mb-3">
        <v-sheet>
          <v-calendar
            v-model="start"
            ref="calendar"
            :type="type"
            locale="it"
            :short-months="false"
            color="primary"
          ></v-calendar>
        </v-sheet>
      </v-flex>

      <v-flex xs6 class="text-sm-left text-xs-center">
        <v-btn @click="$refs.calendar.prev()">
          <v-icon dark left>keyboard_arrow_left</v-icon>Precedente
        </v-btn>
      </v-flex>
      <v-flex xs6 class="text-sm-right text-xs-center">
        <v-btn @click="$refs.calendar.next()">
          Successivo
          <v-icon right dark>keyboard_arrow_right</v-icon>
        </v-btn>
      </v-flex>
    </v-layout>
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
  }
};
</script>
