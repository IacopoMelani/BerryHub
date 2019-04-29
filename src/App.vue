<template>
  <div id="app">
    <Home></Home>
  </div>
</template>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-bottom: 20px;
}
html,
body {
  height: 100%;
  widows: 100%;
  overflow: hidden;
}

#scroll::-webkit-scrollbar-track {
  border-radius: 10px;
  background-color: transparent;
}

#scroll::-webkit-scrollbar {
  width: 5px;
  background-color: whitesmoke;
}

#scroll::-webkit-scrollbar-thumb {
  border-radius: 25%;
  background-color: whitesmoke;
}

@import url("./assets/css/main.css");
</style>

<script>
import Home from "./views/Home";
export default {
  name: "App",
  components: {
    Home
  },
  mounted() {
    this.$router.push("/weather");
  }
};
</script>

