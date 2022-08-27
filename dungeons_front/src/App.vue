<template>
  <v-app class="bg">
    <v-main>
      <Dungeon/>
    </v-main>
  </v-app>
</template>

<script>
import Dungeon from './components/Dungeon.vue'

export default {
  name: 'App',

  components: {
    Dungeon,
  },

  data: () => ({
    //
  }),
}
</script>

<style>
body {
  background-color: #392258;
}
.bg {
  background-color: #392258;
}
</style>