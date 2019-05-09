<template>
  <div class="clock">
    <div class="inside">
      <div class="content">
        <p class="days">
          <span :class="{'today': dayToday == 1}" id="mon">LUN</span>&nbsp;&nbsp;
          <span :class="{'today': dayToday == 2}" id="tus">MAR</span>&nbsp;&nbsp;
          <span :class="{'today': dayToday == 3}" id="wed">MER</span>&nbsp;&nbsp;
          <span :class="{'today': dayToday == 4}" id="thu">GIO</span>&nbsp;&nbsp;
          <span :class="{'today': dayToday == 5}" id="fri">VEN</span>&nbsp;&nbsp;
          <span :class="{'today': dayToday == 6}" id="sat">SAB</span>&nbsp;&nbsp;
          <span :class="{'today': dayToday == 0}" id="sun">DOM</span>
        </p>
        <br>
        <p class="text">
          <span id="hours">{{h}}</span> :
          <span id="min">{{m}}</span> :
          <span id="sec">{{s}}</span>&nbsp;&nbsp;
          <span id="time">{{day}}</span>
        </p>
        <p id="date">
          <span id="cal">
            <img src="https://i.imgur.com/DzEzjHh.png?1" alt="cal">
          </span>
          <span id="day">{{date}}</span>
          <span id="month">{{month}}</span>
          <span id="year">{{year}}</span>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.centered {
  align-items: center;
  display: flex;
}
.clock {
  position: absolute;
  margin: auto;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 600px;
  height: 350px;
  background-color: #333;
  text-align: center;
  color: #fff;
  font-family: "Quicksand", "Helvetica", Arial, sans-serif;
  font-weight: lighter;
}

.inside {
  position: absolute;
  margin: auto;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 500px;
  height: 250px;
  background-color: #222;
}

.content {
  margin-top: 40px;
  width: auto;
  height: auto;
  text-align: center;
}

.days {
  color: rgba(255, 255, 255, 0.2);
}

.today {
  color: whitesmoke;
}

.days span {
  margin-left: 10px;
}

.text {
  position: absolute;
  margin: auto;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  margin-top: 100px;
  font-size: 3.5em;
}

#time {
  font-size: 1.5rem;
}

#date {
  margin-top: 100px;
  margin-left: -100px;
}

#cal img {
  width: 30px;
  height: 30px;
}

span#day {
  position: absolute;
  margin-top: 15px;
  margin-left: 10px;
}

span#month {
  position: absolute;
  margin-top: 15px;
  margin-left: 35px;
}

span#year {
  position: absolute;
  margin-top: 15px;
  margin-left: 55px;
}
</style>

<script>
export default {
  data: () => ({
    time: new Date(),
    timeId: undefined
  }),
  computed: {
    h: function() {
      return this.time.getHours() < 10
        ? "0" + this.time.getHours()
        : this.time.getHours();
    },
    m: function() {
      return this.time.getMinutes() < 10
        ? "0" + this.time.getMinutes()
        : this.time.getMinutes();
    },
    s: function() {
      return this.time.getSeconds() < 10
        ? "0" + this.time.getSeconds()
        : this.time.getSeconds();
    },
    day: function() {
      return this.h < 11 ? "AM" : "PM";
    },
    dayToday: function() {
      return this.time.getDay();
    },
    date: function() {
      return this.time.getDate();
    },
    month: function() {
      return this.time.getMonth();
    },
    year: function() {
      return this.time.getFullYear();
    }
  },
  mounted: function() {
    this.time = new Date();

    this.timeId = setInterval(() => {
      this.time = new Date();
    }, 1000);
  },
  destroyed: function() {
    clearInterval(this.timeId);
  }
};
</script>