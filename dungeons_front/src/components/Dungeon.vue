<template>
  <v-container class="vertical-center">
    <notifications position="top center" classes="notif vue-notification"/>

    <!-- Dungeon ending -->
    <v-dialog transition="dialog-top-transition" persistent v-model="finished" max-width="600">
      <v-card>
        <v-toolbar color="secondary" dark>Dungeon {{rewards.win ? "finished" : "lost"}} !</v-toolbar>
        <v-card-text>
            <v-row>
              <v-col cols="12" class="text-center">
                <v-img class="gods" :src="rewards.win ? require('@/assets/gods/w.png') : require('@/assets/gods/l.png')"/>
              </v-col>
              <v-col cols="12" class="text-center" style="margin-top:-200px;">
                <v-container>
                  <p>Lien(s) : {{rewards.liens}}Ⱡ</p>
                  <p>Box(es) : {{rewards.ccBox}}</p>
                  <p>Arm(s) : {{rewards.arms}}</p>
                  <p>Minion(s) : {{rewards.minions}}</p>
                </v-container>
              </v-col>
            </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>

    <!-- Ambrosius -->
    <v-dialog transition="dialog-top-transition" v-model="ambrosius">
      <v-card>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12" class="text-center">
                <v-img class="ambrosius" :src="require('@/assets/ambrosius.png')"/>
              </v-col>
              <v-col v-for="choice in choices" :key="choice" cols="12" class="text-center">
                <v-btn block v-on:click="ambrosiusChoice(choice.index)">
                    {{choice.message}}
                </v-btn>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-row>
      <!-- Title -->
      <v-container>
        <v-row>
          <v-col cols="12" class="text-center">
            <h1>
              Escape the dungeon !
            </h1>
          </v-col>
        </v-row>
      </v-container>

      <!-- Life -->
      <v-container>
        <v-row>
          <v-col cols="12" class="text-center">
            <v-container>
              <img class="heart" style="vertical-align: bottom;" :src="require('@/assets/health.png')"/>
              <span class="text-warning ml-2">{{life}}</span>
            </v-container>
          </v-col>
        </v-row>
      </v-container>

      <!-- Dungeon box -->
      <v-container ref="dungeon" :class="reflow ? '' : animation ">
        <v-col v-for="row in rows" :key="row" cols="12">
          <v-row :id="row" justify="center">
            <img :class="grid[row][col] != 6 ? 'tile' : 'tileVoid' " v-for="col in columns" :key="col" :src="chooseAsset(row, col)"/>
          </v-row>
        </v-col>
      </v-container>
      
      <!-- Controller -->
      <v-container class="mt-5">
          <v-row>
              <v-col cols="5">
              </v-col>
              <v-col cols="2" class="text-center controller">
                  <v-btn block v-on:click="movePlayer('up')">
                      UP
                  </v-btn>
              </v-col>
              <v-col cols="5">
              </v-col>
              <v-col cols="3">
              </v-col>
              <v-col cols="2" class="text-right controller">
                  <v-btn block v-on:click="movePlayer('left')">
                      LEFT
                  </v-btn>
              </v-col>
              <v-col cols="2" class="text-center controller">
                  <v-btn block v-on:click="movePlayer('down')">
                      DOWN
                  </v-btn>
              </v-col>
              <v-col cols="2" class="text-left controller">
                  <v-btn block v-on:click="movePlayer('right')">
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
import background from "@/assets/floor.png"
import wall from "@/assets/wall.jpg"
import fow from "@/assets/2.png" // Unused
// import player from "@/assets/3.png" // Unused
import money from "@/assets/money.png"
import door from "@/assets/door.png"
import empty from "@/assets/void.png"
import ennemy from "@/assets/ennemy.png"
import ambrosius from "@/assets/ambrosius_tile.png"
import ding from "@/assets/ding.mp3"
import other_wall from "@/assets/other_wall.png"

import player_up from "@/assets/player/up.png"
import player_down from "@/assets/player/down.png"
import player_left from "@/assets/player/left.png"
import player_right from "@/assets/player/right.png"

import arm_1 from "@/assets/arm/1.png"
import arm_2 from "@/assets/arm/2.png"

import minion_1 from "@/assets/minion/0.png"
import minion_2 from "@/assets/minion/1.png"
import minion_3 from "@/assets/minion/2.png"

import io from 'socket.io-client'

const connectRoute = "dungeonConnect";
const moveRoute = "dungeonMove";
const ambrosiusChoice = "ambrosiusChoice";

const token = localStorage.getItem('token') || "test"
export default {
  name: 'DungeonPage',
  
  data: function () {
      return {
          grid: [
            [0,0,0],
            [0,0,0],
            [0,0,0],
          ],
          lastUpdate: [],
          random: {},
          assets:[
            background,
            wall,
            fow,
            player_up,
            money,
            door,
            empty,
            ennemy,
            ambrosius,
            [arm_1, arm_2],
            [minion_1, minion_2, minion_3],
            other_wall,
          ],
          player_directions: [
            player_right,
            player_down,
            player_left,
            player_up,
          ],
          rewards: undefined,
          reflow: false,
          animation: '',
          socket: undefined,
          finished: false,
          life: 150,
          ambrosius: false,
          choices: undefined,
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
    window.addEventListener("keyup", e => {
      switch (e.key) {
        case "ArrowUp":
          this.movePlayer("up");
          break;
        case "ArrowDown":
          this.movePlayer("down");
          break;
        case "ArrowLeft":
          this.movePlayer("left");
          break;
        case "ArrowRight":
          this.movePlayer("right");
          break;
      }
    });
  },

  methods: {
    modifyDungeon(grid) {
      this.lastUpdate = grid;
      let centerOfGridX = Math.floor(grid.length / 2);
      let centerOfGridY = Math.floor(grid[0].length / 2);

      if (grid[centerOfGridX][centerOfGridY].message != undefined) {
        var audio = new Audio(ding);
        audio.volume = 0.2;
        audio.play();
        this.animation = "loot";
        this.activateClass();
        this.$notify({
          text:grid[centerOfGridX][centerOfGridY].message
        });
      }

      if (grid[centerOfGridX][centerOfGridY].choices != undefined) {
        this.ambrosius = true;
        this.choices = grid[centerOfGridX][centerOfGridY].choices;
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

    receiveUpdate(data) {
      this.modifyDungeon(data.g); // Grid update
      this.life = data.h; // Life update
      if (data.e) {
        // Dungeon ended
        this.rewards = data.r;
        this.rewards.win = data.w;
        this.finished = true;
      }
    },

    connectToWS() {
      this.socket = io(process.env.VUE_APP_BACKEND_WS_URL, { transports: ['websocket'] })
      this.socket.on('connect', () => {
        this.sendTokenToWS();
      })
    },

    sendTokenToWS() {
      // get token from local storage :
      let data = {
        body : {
          token: token,
        }
      }
      this.socket.emit(connectRoute, data, (data) => {
        console.log("connected & tokened");
        this.receiveUpdate(data.body);
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
      this.assets[3] = this.player_directions[dir];
      let data = {
        body : {
          token: token,
          direction: dir,
        }
      }
      this.socket.emit(moveRoute, data, (data) => {
        this.receiveUpdate(data.body);
      })

      // trigger reflow
      this.activateClass();
    },

    ambrosiusChoice(choice) {
      let data = {
        body : {
          choice: choice,
        }
      }
      console.log(choice);
      this.socket.emit(ambrosiusChoice, data, (data) => {
        this.ambrosius = false;
        this.$notify({
          token: token,
          text:data.text
        });
      })
    },

    chooseAsset(row, col) {
      let cell;
      try {
        cell = this.lastUpdate[row][col];
        cell.type = this.grid[row][col];
      } catch (e) {
        cell = {
          type:0,
          id:0,
        }
      }
      let asset = this.assets[cell.type];
      if (asset == undefined) {
        asset = empty;
      }
      if (typeof(asset) == "object") {
        if (this.random[cell.id] == undefined) {
          asset = asset[Math.floor(Math.random() * asset.length)];
          this.random[cell.id] = asset;
        } else {
          asset = this.random[cell.id];
        }
      }
      return asset;
    }
  },
}
</script>

<style scoped>
@keyframes moveLeft {
  0% {transform: translateX(0);}
  50% {transform: translateX(.75%); /* .75 */}
  100% {transform: translateX(0);}
}

@keyframes moveRight {
  0% {transform: translateX(0);}
  50% {transform: translateX(-.75%); /* .75 */}
  100% {transform: translateX(0);}
}

@keyframes moveUp {
  0% {transform: translateY(0);}
  50% {transform: translateY(-4%); /* 4 */}
  100% {transform: translateY(0);}
}

@keyframes moveDown {
  0% {transform: translateY(0);}
  50% {transform: translateY(-4%); /* 4 */}
  100% {transform: translateY(0);}
}

@keyframes loot {
  30% { transform: scale(1.2); }
  40%, 60% { transform: rotate(-5deg) scale(1.2); }
  50% { transform: rotate(5deg) scale(1.2); }
  70% { transform: rotate(0deg) scale(1.2); }
  100% { transform: scale(1); }
}

@keyframes heartbeat{
  0%, 40%, 80%, 100%{transform: scale( 1 );}
  20%,60%{transform: scale( 1.25 );}
  60%{transform: scale( 1.25 );}
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

.tile {
  width: 100px;
  height: 100px;
  background-color:white;
  margin: -1px;
}
.tileVoid {
  width: 100px;
  height: 100px;
  margin: -1px;
}

.heart {
  animation: heartbeat 1.75s ease-in-out infinite;
}

h1 {
  color:white;
}

.controller {
  padding: 3px;
}

.ambrosius {
  height: 175px;
  width: auto;
}

.gods {
  height: 200px;
  width: auto;
  object-fit: cover;
  opacity: 35%;
}
</style>

<style>
.notif {
  /* higher font size */
  font-size: 1.5em !important;
}
</style>