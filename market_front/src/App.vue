<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">RWBYzon</v-tab>
      </template>

      <AdvancedSearch :open="dialogSearch" @close="dialogSearch = false;" ></AdvancedSearch>

      <v-row class="justify-center">
        <v-col cols="12" sm="6">
          <v-text-field 
            hide-details 
            append-icon="mdi-magnify" 
            single-line 
            outlined

            label="Search"
            density="compact"
            variant="outlined"
            @focus="dialogSearch = true"
          ></v-text-field>
        </v-col>
      </v-row>

      <template v-slot:append>
        <v-btn prepend-icon="mdi-discord" v-if="!logged_in" :href="login_link">Log in</v-btn>
        <v-btn prepend-icon="mdi-discord" v-else @click="logout()">Log out</v-btn>
      </template>
    </v-app-bar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import socket from '@/plugins/websocket';
import axios from 'axios';
import { authStore } from "@/store/authStore";
import AdvancedSearch from './components/AdvancedSearch.vue';
import { process } from 'ipaddr.js';

export default {
  name: 'App',

  components: { AdvancedSearch },

  data: () => ({
    dialogSearch: false,
    value_range: [0, 100],
    level_range: [1, 500],
    buffs_range: [0, 2],
    rarity_range: [0, 5],
    order_by: undefined,
    orderState: false,
    filters: {
      name_has: '',
      value_above: '0',
      value_below: '100',
      level_above: '1',
      level_below: '500',
      buffs_above: '0',
      buffs_below: '2',
      rarity_above: '0',
      rarity_below: '5',
      order_by: '',
      order_type: '',
    },
    login_link: undefined,
    logged_in: false,
  }),

  mounted() {
    this.waitForConnect();
    this.connect(true);
  },

  methods : {
    async waitForConnect() {
      socket.removeAllListeners()
      while(!socket.connected) {
        console.log("Waiting for socket connection...")
        await new Promise(r => setTimeout(r, 1000));
      }
    },
    logout: function() {
      // redirect to logginlink
      window.location=this.login_link
      authStore.commit("reset");
    },
    getToken: async function() {
      let data = await axios.get(process.env.VUE_APP_BACKEND + 'token', { withCredentials: true });
      await authStore.commit("setToken", data.data.token);
    },
    connect: function(retry) {
      let data = {
        body: {
          token: authStore.getters.token
        }
      }
      socket.emit("marketConnect", data, (data) => {
        if (data.success) {
          authStore.commit("setLogin", data.body.connected);
          this.logged_in = data.body.connected;
          this.login_link = data.body.link;
          if (!data.body.connected) {
            authStore.commit("setLoginLink", data.body.link);
          }
        } else if (retry){
          this.logged_in = false;
          this.getToken().then(() =>{
            this.connect(false);
          });
        }
      });
    },
  }
}
</script>
