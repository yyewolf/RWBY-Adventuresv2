<template>
  <v-app>
    <v-app-bar elevation="4">
      <v-row justify="center">
        <h1 class="text-primary">Hello, {{username}}</h1>
      </v-row>
    </v-app-bar>
    <v-main>
      <v-container>
        <v-row class="mb-3" justify="center">
          <h2>{{arena.title}}</h2>
        </v-row>
        <v-row class="mb-5" justify="center">
          <h4>There are {{playerAmount}} people with you.</h4>
        </v-row>
        <v-row justify="center">
          <v-card width="20vw" @click="click()">
            <v-progress-linear :model-value="curHealth" color="red"></v-progress-linear>
            <v-card-text>
              <v-img ref="mob" :class="reflow ? '' : 'animation'" :src="arena.image"/>
            </v-card-text>
          </v-card>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
  import io from 'socket.io-client'
  const token = localStorage.getItem('token') || "test"
  
  export default {
    name: 'App',
    data: () => ({
      hit:0,
      reflow: false,
      curHealth: 100,
      playerAmount: 1,
      socket: undefined,
      username : 'Yewolf',
      arena : {
        title: "You have to defeat a Grimm",
        image: "https://www.pngfind.com/pngs/m/78-781864_nevermore-rwby-png-download-rwby-volume-4-grimm.png"
      }
    }),
    mounted() {
      this.connectToWS();
    },
    methods: {
      connectToWS() {
        this.socket = io(process.env.VUE_APP_BACKEND_WS_URL, { transports: ['websocket'] });
        this.socket.on('connect', () => {
          this.sendTokenToWS();
        })
        this.socket.on('arenaLoop', (data) => {
          if (this.curHealth > data.body.h) {
            this.animate();
          }
          this.curHealth = data.body.h;
          this.playerAmount = data.body.n;
        });
      },
  
      animate() {
        this.reflow = true;
        this.$refs.mob.offsetWidth;
        setTimeout(() => {
          this.reflow = 0;
        }, 0);
      },
  
      sendTokenToWS() {
        console.log("connecting")
        let data = {
          body : {
            token: token,
          }
        }
        this.socket.emit('arenaConnect', data, (resp) => {
          if (resp.success) {
            this.arena = resp.body.arena;
          }
        })
      },
  
      click: function() {
        if (this.curHealth > 0) {
            this.socket.emit('arenaHit', {
                body: {
                    token: token,
                },
            });
        }
        this.animate();
      }
    }
  }
  </script>
  
  <style scoped>
  @-webkit-keyframes shake-lr {
    0%,
    100% {
      -webkit-transform: rotate(0deg);
              transform: rotate(0deg);
      -webkit-transform-origin: 50% 50%;
              transform-origin: 50% 50%;
    }
    10% {
      -webkit-transform: rotate(8deg);
              transform: rotate(8deg);
    }
    20%,
    40%,
    60% {
      -webkit-transform: rotate(-10deg);
              transform: rotate(-10deg);
    }
    30%,
    50%,
    70% {
      -webkit-transform: rotate(10deg);
              transform: rotate(10deg);
    }
    80% {
      -webkit-transform: rotate(-8deg);
              transform: rotate(-8deg);
    }
    90% {
      -webkit-transform: rotate(8deg);
              transform: rotate(8deg);
    }
  }
  @keyframes shake-lr {
    0%,
    100% {
      -webkit-transform: rotate(0deg);
              transform: rotate(0deg);
      -webkit-transform-origin: 50% 50%;
              transform-origin: 50% 50%;
    }
    10% {
      -webkit-transform: rotate(8deg);
              transform: rotate(8deg);
    }
    20%,
    40%,
    60% {
      -webkit-transform: rotate(-10deg);
              transform: rotate(-10deg);
    }
    30%,
    50%,
    70% {
      -webkit-transform: rotate(10deg);
              transform: rotate(10deg);
    }
    80% {
      -webkit-transform: rotate(-8deg);
              transform: rotate(-8deg);
    }
    90% {
      -webkit-transform: rotate(8deg);
              transform: rotate(8deg);
    }
  }
  .animation {
    -webkit-animation: shake-lr 0.2s cubic-bezier(0.455, 0.030, 0.515, 0.955) both;
    animation: shake-lr 0.2s cubic-bezier(0.455, 0.030, 0.515, 0.955) both;
  }
  </style>
  