<template>
  <v-container fluid class="bgcolor-2 centered">
    <v-layout row wrap justify-center>
      <v-flex xs12 class="scroll-content">
        <v-card v-for="news in newsList.articles" :key="news.title">
          <v-img :src="news.urlToImage" aspect-ratio="2.75"></v-img>
          <v-card-title primary-title>
            <div>
              <h3 class="headline mb-0">{{news.title}}</h3>
              <div>{{news.description}}</div>
            </div>
          </v-card-title>
          <v-card-actions>
            <v-btn flat color="orange" :href="news.url">Explore</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
      <v-progress-circular v-if="isUpdating" :size="70" :width="7" color="red" indeterminate></v-progress-circular>
      <h4 class="mb-1">powered by NewsAPI.org</h4>
    </v-layout>
  </v-container>
</template>

<style scoped>
.scroll-content {
  height: 80vh;
  overflow: auto;
}
</style>

<script>
export default {
  name: "News",
  computed: {
    isUpdating: function() {
      return this.$store.state.news.updating;
    },
    newsList: function() {
      return this.$store.state.news.list;
    }
  },
  mounted: function() {
    this.$store.dispatch("news/getNewsList");

    this.getNewsListId = setInterval(() => {
      this.$store.commit("news/resetCounter");
      this.$store.dispatch("news/getNewsList");
    }, 60000);
  },
  destroyed: function() {
    clearInterval(this.getNewsListId);
  }
};
</script>
