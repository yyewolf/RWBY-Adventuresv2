<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">OC Contest</v-tab>
      </template>

      <template v-slot:append>
        <v-btn prepend-icon="mdi-discord" v-if="!logged_in" :href="login_link">Log in</v-btn>
        <v-btn prepend-icon="mdi-discord" v-else>Log out</v-btn>
      </template>
    </v-app-bar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'App',

  data: () => ({
    //
  }),
}
</script>
