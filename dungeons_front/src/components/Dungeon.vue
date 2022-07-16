<template>
  <v-container class="vertical-center">
    <notifications position="top center" classes="notif vue-notification"/>
    <v-row>
      <v-container ref="dungeon" :class="reflow ? '' : animation ">
        <v-col v-for="row in rows" :key="row" cols="12">
          <v-row :id="row" justify="center">
            <img v-for="col in columns" :key="col" :src="assets[grid[row][col]]"/>
          </v-row>
        </v-col>
      </v-container>
      
      <v-container class="mt-5">
          <v-row>
              <v-col cols="12" class="text-center">
                  <v-btn v-on:click="movePlayer('up');">
                      UP
                  </v-btn>
              </v-col>
              <v-col cols="3">
              </v-col>
              <v-col cols="2" class="text-right">
                  <v-btn v-on:click="movePlayer('left')">
                      LEFT
                  </v-btn>
              </v-col>
              <v-col cols="2" class="text-center">
                  <v-btn v-on:click="movePlayer('down')">
                      DOWN
                  </v-btn>
              </v-col>
              <v-col cols="2" class="text-left">
                  <v-btn v-on:click="movePlayer('right')">
                      RIGHT
                  </v-btn>
              </v-col>
              <v-col cols="3">
              </v-col>
          </v-row>
      </v-container>
    </v-row>
  </v-container>
</template>

<script>
import background from "@/assets/0.png"
import wall from "@/assets/1.png"
import fow from "@/assets/2.png"
import player from "@/assets/3.png"
import money from "@/assets/4.png"
import ding from "@/assets/ding.mp3"

import io from 'socket.io-client'

const connectRoute = "dungeonConnect";
const moveRoute = "dungeonMove";

export default {
  name: 'DungeonPage',
  
  data: function () {
      return {
          grid: [
            [0,0,0],
            [0,0,0],
            [0,0,0],
          ],
          assets:[
            background,
            wall,
            fow,
            player,
            money,
          ],
          reflow: false,
          animation: '',
          socket: undefined,
      }
  },

  computed: {
    columns() {
      return Array.from({ length: this.grid.length }, (_, i) => i)
    },
    rows() {
      return Array.from({ length: this.grid[0].length }, (_, i) => i)
    }
  },

  mounted() {
    this.connectToWS();
  },

  methods: {
    modifyDungeon(grid) {
      let centerOfGridX = Math.floor(grid.length / 2);
      let centerOfGridY = Math.floor(grid[0].length / 2);

      if (grid[centerOfGridX][centerOfGridY].message != "") {
        var audio = new Audio(ding);
        audio.volume = 0.2;
        audio.play();
        this.animation = "loot";
        this.activateClass();
        this.$notify({
          text:grid[centerOfGridX][centerOfGridY].message
        });
      }
      for (let row = 0; row < grid.length; row++) {
        for (let col = 0; col < grid[row].length; col++) {
          if (col == centerOfGridY && row == centerOfGridX) {
            this.grid[col][row] = 3;
            continue;
          }
          this.grid[col][row] = grid[col][row].type;
        }
      }
    },

    connectToWS() {
      this.socket = io('ws://localhost:9003/', { transports: ['websocket'] })
      this.socket.on('connect', () => {
        this.sendTokenToWS();
      })
    },

    sendTokenToWS() {
      // get token from local storage :
      // const token = localStorage.getItem('token')
      
      const token = "test"
      let data = {
        body : {
          token: token,
        }
      }
      this.socket.emit(connectRoute, data, (data) => {
        console.log("connected & tokened");
        this.modifyDungeon(data.body.g);
      })
    },

    activateClass(){
      this.reflow = 1;
      this.$refs.dungeon.offsetWidth;
      setTimeout(() => {
        this.reflow = 0;
      }, 0);
    },

    movePlayer(direction) {
      let dir;
      switch(direction) {
        case 'up':
          dir = 3;
          this.animation = 'animateU';
          //this.playerPos.row--;
          break;
        case 'down':
          dir = 1;
          this.animation = 'animateD';
          //this.playerPos.row++;
          break;
        case 'left':
          dir = 2;
          this.animation = 'animateL';
          //this.playerPos.col--;
          break;
        case 'right':
          dir = 0;
          this.animation = 'animateR';
          //this.playerPos.col++;
          break;
      }

      let data = {
        body : {
          direction: dir,
        }
      }
      this.socket.emit(moveRoute, data, (data) => {
        this.modifyDungeon(data.body.g);
      })

      // trigger reflow
      this.activateClass();
    },
  },
}
</script>

<style scoped>
@keyframes moveLeft {
  0% {
    transform: translateX(0);
  }
  50% {
    transform: translateX(.75%); /* .75 */
  }
  100% {
    transform: translateX(0);
  }
}

@keyframes moveRight {
  0% {
    transform: translateX(0);
  }
  50% {
    transform: translateX(-.75%); /* .75 */
  }
  100% {
    transform: translateX(0);
  }
}

@keyframes moveUp {
  0% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-4%); /* 4 */
  }
  100% {
    transform: translateY(0);
  }
}

@keyframes moveDown {
  0% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-4%); /* 4 */
  }
  100% {
    transform: translateY(0);
  }
}

@keyframes loot {
  30% { transform: scale(1.2); }
  40%, 60% { transform: rotate(-5deg) scale(1.2); }
  50% { transform: rotate(5deg) scale(1.2); }
  70% { transform: rotate(0deg) scale(1.2); }
  100% { transform: scale(1); }
}

.animateL {
  animation: moveLeft 0.25s ease-in-out;
}

.animateR {
  animation: moveRight 0.25s ease-in-out;
}

.animateU {
  animation: moveUp 0.25s ease-in-out;
}

.animateD {
  animation: moveDown 0.25s ease-in-out;
}

.loot{
  animation: loot 0.25s ease-in-out;
}

.vertical-center {
  display: flex;
  align-items: center;
  justify-content: center;
  height:100vh;
}
</style>

<style>
.notif {
  /* higher font size */
  font-size: 1.5em !important;
}
</style>