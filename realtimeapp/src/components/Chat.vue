<template>
  <div class='hello' v-bind:class='{start : isStart}'>
  {{ msg }}
  </div>
</template>

<script>
import io from 'socket.io-client'
export default {
  name: 'Chat',
  data: () => ({
    msg: 'Precione START em outro dispositivo',
    isStart: true,
    socket: io('http://192.168.200.88:5000/socket.io/')
  }),
  methods: {
    chat1 () {
      this.socket.emit('join', 'chat', function () {
        console.log('ACK from server wtih data: ')
      })
    },

    chat2 () {
      this.socket.emit('join', 'chat2', function () {
        console.log('ACK from server wtih data: ')
      })
    }
  },
  mounted () {
    setTimeout(() => {
      this.socket.emit('join', 'chat', function () {
        console.log('ACK from server wtih data: ')
      })
    }, 2000)

    var that = this
    this.socket.on('message', function (val) {
    // Connected, let's sign-up for to receive messages for this room
      that.msg = val
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

.start {
   font-size: 150px;
}

</style>
