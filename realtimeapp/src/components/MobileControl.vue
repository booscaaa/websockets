<template>
  <v-app>
  <v-card class="e4">
    <v-responsive
      :style="{ background: `rgb(${red}, ${green}, ${blue})` }"
      height="300px"
    ></v-responsive>

    <v-card-text>
      <v-container
        fluid
        grid-list-lg
      >
        <v-layout
          row
          wrap
        >
          <v-flex xs9>
            <v-slider
              v-model="red"
              :max="255"
              label="R"
            ></v-slider>
          </v-flex>

          <v-flex xs3>
            <v-text-field
              v-model="red"
              class="mt-0"
              type="number"
            ></v-text-field>
          </v-flex>

          <v-flex xs9>
            <v-slider
              v-model="green"
              :max="255"
              label="G"
            ></v-slider>
          </v-flex>

          <v-flex xs3>
            <v-text-field
              v-model="green"
              class="mt-0"
              type="number"
            ></v-text-field>
          </v-flex>

          <v-flex xs9>
            <v-slider
              :max="255"
              v-model="blue"
              label="B"
            ></v-slider>
          </v-flex>

          <v-flex xs3>
            <v-text-field
              v-model="blue"
              class="mt-0"
              type="number"
            ></v-text-field>
          </v-flex>
        </v-layout>
      </v-container>
    </v-card-text>
  </v-card>
  <v-layout row wrap justify-center align-center>
    <v-flex xs12 sm10 md8 lg5 xl4>
      <v-btn @click="start" name="button">Show me the controls</v-btn>
      <v-btn @click="hide" name="button">Hide the controls</v-btn>
    </v-flex>
  </v-layout>
  </v-app>
</template>

<script>
import io from 'socket.io-client'
export default {
  data: () => ({
    socket: io('http://192.168.200.88:5000'),
    // socket: io('http://localhost:5000/socket.io/'),
    red: 64,
    green: 128,
    blue: 0
  }),
  watch: {
    red: function (val) {
      let o = {
        'msg': this.rgbToHex(val, this.green, this.blue),
        'room': 'mobileControl',
        'red': val,
        'green': this.green,
        'blue': this.blue
      }
      this.socket.emit('message', JSON.stringify(o), function (data) {})
    },
    green: function (val) {
      let o = {
        'msg': this.rgbToHex(this.red, val, this.blue),
        'room': 'mobileControl',
        'red': this.red,
        'green': val,
        'blue': this.blue
      }
      this.socket.emit('message', JSON.stringify(o), function (data) {})
    },
    blue: function (val) {
      let o = {
        'msg': this.rgbToHex(this.red, this.green, val),
        'room': 'mobileControl',
        'red': this.red,
        'green': this.green,
        'blue': val
      }
      this.socket.emit('message', JSON.stringify(o), function (data) {})
    }
  },
  methods: {
    componentToHex (c) {
      var hex = c.toString(16)
      return hex.length === 1 ? '0' + hex : hex
    },

    rgbToHex (r, g, b) {
      return '#' + this.componentToHex(r) + this.componentToHex(g) + this.componentToHex(b)
    },

    start (){
      let o = {
        'msg': 'controls',
        'room': 'mobileControl'
      }
      this.socket.emit('message', JSON.stringify(o), function (data) {})
    },

    hide () {
      let o = {
        'msg': 'hide',
        'room': 'mobileControl'
      }
      this.socket.emit('message', JSON.stringify(o), function (data) {})
    }
  },
  mounted () {
    var that = this

    this.socket.on('connect', function (val) {
      setTimeout(() => {
        that.socket.emit('join', 'mobileControl', function () {})
        that.socket.emit('join', 'mobileControl1', function () {})
      }, 2000)
    })
  }
}
</script>
