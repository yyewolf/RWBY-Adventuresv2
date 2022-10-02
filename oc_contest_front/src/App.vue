<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">OC Contest</v-tab>
      </template>

      <template v-slot:append>
        <v-btn prepend-icon="mdi-discord" v-if="!logged" :href="login">Log in</v-btn>
        <v-btn prepend-icon="mdi-discord" v-else :href="logout">Log out</v-btn>
      </template>
    </v-app-bar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import {loggedIn} from '@/auth/logged';

export default {
  name: 'App',

  data: () => ({
    logged: false,
    login: process.env.VUE_APP_BACKEND + 'auth/login',
    logout: process.env.VUE_APP_BACKEND + 'auth/logout',
  }),

  created() {
    loggedIn().then(data => {
      this.logged = data;
      console.log(data);
    });
  },
}
</script>
