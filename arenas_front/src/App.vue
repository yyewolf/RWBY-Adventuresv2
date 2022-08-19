<template>
  <v-app>
    <v-app-bar elevation="4">
      <v-row justify="center">
        <h1 class="text-primary">Hello, {{username}}</h1>
      </v-row>
    </v-app-bar>
    <v-main>
      <Arena :arena="arena" />
    </v-main>
  </v-app>
</template>

<script>
import Arena from './components/Arena.vue'

export default {
  name: 'App',

  components: {
    Arena,
  },

  data: () => ({
    username : 'Yewolf',
    arena : {
      title: "You have to defeat a Grimm",
      people: 1,
      image: "https://www.pngfind.com/pngs/m/78-781864_nevermore-rwby-png-download-rwby-volume-4-grimm.png"
    }
  }),
}
</script>
