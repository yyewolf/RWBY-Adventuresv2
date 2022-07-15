<template>
  <v-container class="vertical-center">
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
                  <v-btn v-on:click="movePlayer('up')">
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

export default {
  name: 'DungeonPage',
  
  data: function () {
      return {
          grid: [
            [1,1,1,1,1],
            [1,3,0,0,1],
            [1,0,0,0,1],
            [1,0,0,0,1],
            [1,1,1,1,1],
          ],
          assets:[
            background,
            wall,
            fow,
            player,
          ],
          playerPos:{
            row: 1,
            col: 1,
          },
          memory: 0,
          reflow: false,
          animation: '',
      }
  },

  computed: {
    columns() {
      return Array.from({ length: 5 }, (_, i) => i)
    },
    rows() {
      return Array.from({ length: 5 }, (_, i) => i)
    }
  },

  methods: {
    activateClass(){
      this.reflow = 1;
      this.$refs.dungeon.offsetWidth;
      setTimeout(() => {
        this.reflow = 0;
      }, 0);
    },
    movePlayer(direction) {

      this.grid[this.playerPos.row][this.playerPos.col] = this.memory;
      switch(direction) {
        case 'up':
          this.animation = 'animateU';
          this.playerPos.row--;
          break;
        case 'down':
          this.animation = 'animateD';
          this.playerPos.row++;
          break;
        case 'left':
          this.animation = 'animateL';
          this.playerPos.col--;
          break;
        case 'right':
          this.animation = 'animateR';
          this.playerPos.col++;
          break;
      }
      if (this.playerPos.row < 0) {
        this.playerPos.row = 0;
      }
      if (this.playerPos.col < 0) {
        this.playerPos.col = 0;
      }
      if(this.playerPos.row > 4) {
        this.playerPos.row = 4;
      }
      if(this.playerPos.col > 4) {
        this.playerPos.col = 4;
      }

      // trigger reflow
      this.activateClass();

      this.memory = this.grid[this.playerPos.row][this.playerPos.col];
      this.grid[this.playerPos.row][this.playerPos.col] = 3
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
    transform: translateX(0.5%);
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
    transform: translateX(-0.5%);
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
    transform: translateY(-3%);
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
    transform: translateY(-3%);
  }
  100% {
    transform: translateY(0);
  }
}

.animateL{
  animation: moveLeft 0.25s ease-in-out;
}

.animateR{
  animation: moveRight 0.25s ease-in-out;
}

.animateU{
  animation: moveUp 0.25s ease-in-out;
}

.animateD{
  animation: moveDown 0.25s ease-in-out;
}

.vertical-center {
  display: flex;
  align-items: center;
  justify-content: center;
  height:100vh;
}

</style>