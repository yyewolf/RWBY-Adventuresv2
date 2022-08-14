<template>
  <notifications position="top center" classes="notif vue-notification"/>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-card :loading="!noAuctions && auctions.length == 0" elevation="2" height="100%">
          <v-card-title class="mb-5">
            <p class="text-h5" style="float:left;">Latest auctions</p>
            <v-btn variant="plain" href="/about" style="float:right">(View all)</v-btn>
          </v-card-title>
          <v-card-text>
            <v-card-subtitle v-if="noAuctions"> No auctions available... </v-card-subtitle>
            <v-layout row class="scrollbar" v-if="!noAuctions">
              <div class="ms-4 mb-5" v-for="a in auctions" :key="a">
                <Auction :data="a"></Auction>
              </div>
            </v-layout>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card :loading="personas.length == 0" elevation="2" height="100%">
          <v-card-title class="mb-5">
            <p class="text-h5" style="float:left;">Random Characters</p>
          </v-card-title>
          <v-card-text>
            <v-container class="d-flex justify-center">
              <div class="d-flex scrollbar">
                <div class="ms-4 mb-5" v-for="p in personas" :key="p">
                  <Persona :data="p"></Persona>
                </div>
              </div>
            </v-container>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card :loading="!noListings && listings.length == 0" elevation="2" height="100%">
          <v-card-title class="mb-5">
            <p class="text-h5" style="float:left;">Latest listings</p> 
            <v-btn variant="plain" href="/about" style="float:right;">(View all)</v-btn>
          </v-card-title>
          <v-card-text>
            <v-card-subtitle v-if="noListings"> No listings available... </v-card-subtitle>
            <v-layout row class="scrollbar" v-if="!noListings">
              <div class="ms-4 mb-5" v-for="l in listings" :key="l">
                <Listing :data="l"></Listing>
              </div>
            </v-layout>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import socket from '@/plugins/websocket';
import Listing from './Listing.vue'
import Auction from './Auction.vue'
import Persona from './Persona.vue'

const latestListingsRoute = "listings/latest"
const latestAuctionsRoute = "auctions/latest"
const randomPersonasRoute = "randomPersonas"

export default {
    name: "HomePage",
    data: () => ({
        listings: [],
        auctions: [],
        personas: [],

        noListings: false,
        noAuctions: false,
    }),
    mounted() {
      // Wait for socket connection
      this.waitForConnect();
    },
    methods: {
        async waitForConnect() {
          socket.removeAllListeners()
          while(!socket.connected) {
            console.log("Waiting for socket connection...")
            await new Promise(r => setTimeout(r, 1000));
          }
          this.getListings();
          this.getAuctions();
          this.getRandomPersonas();
        },
        getListings() {
            socket.emit(latestListingsRoute, {}, (data) => {
                if (!data.body.listings) {
                    this.noListings = true;
                    return
                }
                for (let i = 0; i < data.body.listings.length; i++) {
                  data.body.listings[i].icon = data.body.icons[i];
                }
                this.listings = data.body.listings;
            });
        },
        getAuctions() {
            socket.emit(latestAuctionsRoute, {}, (data) => {
                if (!data.body.auctions) {
                  this.noAuctions = true; 
                  return
                }
                for (let i = 0; i < data.body.auctions.length; i++) {
                  data.body.auctions[i].icon = data.body.icons[i];
                }
                this.auctions = data.body.auctions;
            });
        },
        getRandomPersonas() {
          socket.emit(randomPersonasRoute, {}, (data) => {
            this.personas = [];
            for (let p of data.body.characters) {
              this.personas.push(p);
            }
            for (let g of data.body.grimms) {
              this.personas.push(g);
            }
          });
        }
    },
    components: { Listing, Persona, Auction }
}
</script>

<style scoped>
.scrollbar {
  overflow-x: scroll;
  height: 100%;
}
.scrollbarF {
  flex-direction: row;
  justify-content: center;
}
.scrollbar::-webkit-scrollbar{
  height: 8px;
  width: 8px;
  background: rgb(75, 75, 75);
  border-radius: 5px;
}
.scrollbar::-webkit-scrollbar-thumb:horizontal{
  background: rgb(41, 39, 39);
  border-radius: 10px;
}
</style>