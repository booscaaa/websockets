<template>
  <v-app id="app">
    <v-container>
      <v-layout row wrap justify-center align-center>
        <v-flex xs12 sm10 md8 lg5 xl4>
          <v-btn @click="start" name="button">START/RESTART</v-btn>
          <v-btn @click="pause" name="button">PAUSE</v-btn>
          <v-btn @click="resume" name="button">RESUME</v-btn>
          <v-container fluid grid-list-lg>
            <v-layout row wrap>
              <v-flex xs12>
                <v-subheader class="pl-0">Altura do card</v-subheader>
                <v-slider color="primary" :max="500" :min="1" v-model="slider" thumb-label="always"></v-slider>
              </v-flex>
            </v-layout>
          </v-container>
        </v-flex>
      </v-layout>
      <v-layout row wrap justify-center align-center>
        <v-flex xs12 sm10 md8 lg5 xl4>
          <v-btn @click="start1" name="button">START/RESTART</v-btn>
          <v-btn @click="pause1" name="button">PAUSE</v-btn>
          <v-btn @click="resume1" name="button">RESUME</v-btn>
          <v-container fluid grid-list-lg>
            <v-layout row wrap>
              <v-flex xs12>
                <v-subheader class="pl-0">Altura do card</v-subheader>
                <v-slider color="primary" :max="500" :min="1" v-model="slider1" thumb-label="always"></v-slider>
              </v-flex>
            </v-layout>
          </v-container>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script>
import io from 'socket.io-client'
export default {
  name: 'Control',
  data: () => ({
    socket: io('http://192.168.200.88:5000'),
    slider: 45,
    interval: null,
    val: 0,
    slider1: 45,
    interval1: null,
    val1: 0
  }),
  watch: {
    slider: function (val) {
      let o = {
        'msg': val.toString(),
        'room': 'chat'
      }
      this.socket.emit('message', JSON.stringify(o), function (data) {})
    },

    slider1: function (val) {
      let o = {
        'msg': val.toString(),
        'room': 'chat1'
      }

      this.socket.emit('message', JSON.stringify(o), function (data) {})
    }
  },
  methods: {
    start () {
      clearInterval(this.interval)
      this.val = 0
      this.interval = setInterval(() => {
        this.val++
        let o = {
          'msg': this.val.toString(),
          'room': 'chat'
        }

        this.socket.emit('message', JSON.stringify(o), function (data) {})
      }, 100)
    },

    pause () {
      clearInterval(this.interval)
    },

    resume () {
      clearInterval(this.interval)
      this.interval = setInterval(() => {
        this.val++
        let o = {
          'msg': this.val.toString(),
          'room': 'chat'
        }

        this.socket.emit('message', JSON.stringify(o), function (data) {})
      }, 100)
    },

    start1 () {
      clearInterval(this.interval1)
      this.val1 = 0
      this.interval1 = setInterval(() => {
        this.val1++
        let o = {
          'msg': this.val1.toString(),
          'room': 'chat1'
        }

        this.socket.emit('message', JSON.stringify(o), function (data) {})
      }, 100)
    },

    pause1 () {
      clearInterval(this.interval1)
    },

    resume1 () {
      clearInterval(this.interval1)
      this.interval1 = setInterval(() => {
        this.val1++
        let o = {
          'msg': this.val1.toString(),
          'room': 'chat1'
        }
        this.socket.emit('message', JSON.stringify(o), function (data) {})
      }, 100)
    }
  },
  mounted () {
    var that = this

    this.socket.on('connect', function (val) {
      setTimeout(() => {
        that.socket.emit('join', 'chat', function () {})
        that.socket.emit('join', 'chat1', function () {})
      }, 2000)

      that.socket.on('message', function (val) {
      // Connected, let's sign-up for to receive messages for this room
        // document.getElementById('image').src = 'data:image/gif;base64,' + val
      })
    })
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
.hello {
  width: 100%;
  height: 500px;

  position:absolute; /*it can be fixed too*/
  left:0; right:0;
  top:0; bottom:0;
  margin:auto;

  /*this to solve 'the content will not be cut when the window is smaller than the content': */
  max-width:100%;
  max-height:100%;
  overflow:auto;
}

</style>
