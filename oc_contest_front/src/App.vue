<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tabs>
          <v-tab to="/">Home</v-tab>
          <template v-if="logged">
            <v-tab to="/self">My submissions</v-tab>
            <v-tab to="/create">New submission</v-tab>
          </template>
        </v-tabs>
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
