<template>
  <v-container>
    <v-card class="elevation-8 pa-5" :class="{'shake': count > 100}">
      <line-chart :chart-data="grafico" :options="op"></line-chart>
    </v-card>
    <audio id="myAudio" muted>
      <source src="/static/alert.wav" type="audio/wav">
    </audio>
  </v-container>
</template>

<script>
import LineChart from './LineChart.js'
import io from 'socket.io-client'
export default {
  components: {
    LineChart
  },
  data: () => ({
    // socket: io('https://websocketservice.vinicius.rs/socket.io/'),
    socket: io('http://localhost:5000/socket.io/'),
    grafico: [],
    op: {},
    audio: null,
    count: 0,
    labels: [],
    data: []
  }),
  watch: {
    count(val) {
      if (val > 100) {
        this.play();
      }
    }
  },
  mounted (){
    this.grafico = {
      labels: [],
      datasets: [
        {
          label: 'Data One',
          backgroundColor: '#f87979',
          data: []
        }
      ]
    }

    this.audio = document.getElementById("myAudio")
    this.op = {responsive: true, maintainAspectRatio: false}


    var that = this
    that.socket.on('connect', function (val) {
      setTimeout(() => {
        that.socket.emit('join', 'chat', function () {})
      }, 2000)
      setTimeout(() => {
        that.socket.emit('join', 'chat1', function () {})
      }, 2000)

      that.socket.on('message', function (val) {
        that.count++
        that.labels.push(val.msg)
        that.data.push(val.msg)
        that.grafico = {
          labels: that.labels,
          datasets: [
            {
              label: 'Data One',
              backgroundColor: '#f87979',
              data: that.data
            }
          ]
        }
      })
    })
  },
  methods: {
    play () {
      this.audio.loop = true
      this.audio.play()
    }
  }
}
</script>

<style>
.small {
  width: 100%;
}

body{
  background-color:#FB861D;
}
.content {
  margin: 5% 0 0 0%;
}
.shake {
	display:inline-block
}
.shake {
	display:block;
	position:relative;
  -webkit-animation-name: spaceboots;
	-webkit-animation-duration: 0.8s;
	-webkit-transform-origin:50% 50%;
	-webkit-animation-iteration-count: infinite;
	-webkit-animation-timing-function: linear;
}
img{
display: block;
height: 250px;
width: 250px;
border-radius: 50% 50% 50% 50%;
}
h2.shake,
.shake.inline {
	display:inline-block
}
h1,p,h2{
  color:#FFF;
}
@-webkit-keyframes spaceboots {
	0% { -webkit-transform: translate(2px, 1px) rotate(0deg); }
	10% { -webkit-transform: translate(-1px, -2px) rotate(-4deg); }
	20% { -webkit-transform: translate(-3px, 0px) rotate(4deg); }
	30% { -webkit-transform: translate(0px, 2px) rotate(0deg); }
	40% { -webkit-transform: translate(1px, -1px) rotate(4deg); }
	50% { -webkit-transform: translate(-1px, 2px) rotate(-4deg); }
	60% { -webkit-transform: translate(-3px, 1px) rotate(0deg); }
	70% { -webkit-transform: translate(2px, 1px) rotate(-4deg); }
	80% { -webkit-transform: translate(-1px, -1px) rotate(4deg); }
	90% { -webkit-transform: translate(2px, 2px) rotate(0deg); }
	100% { -webkit-transform: translate(1px, -2px) rotate(-4deg); }
}
.shake:hover,
.shake:focus {
	-webkit-animation-name: spaceboots;
	-webkit-animation-duration: 0.8s;
	-webkit-transform-origin:50% 50%;
	-webkit-animation-iteration-count: infinite;
	-webkit-animation-timing-function: linear;
}
</style>
