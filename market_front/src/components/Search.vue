<template>
  <notifications position="top center" classes="notif vue-notification"/>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-card :loading="!noAuctions && auctions.length == 0" elevation="2" height="100%">
          <v-card-title class="mb-5">
            <p class="text-h5" style="float:left;">Auctions</p>
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
        <v-card :loading="!noListings && listings.length == 0" elevation="2" height="100%">
          <v-card-title class="mb-5">
            <p class="text-h5" style="float:left;">Listings</p> 
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

const listingsRoute = "listings/search"
const auctionsRoute = "auctions/search"

export default {
    name: "HomePage",
    data: () => ({
        listings: [],
        auctions: [],

        filters: undefined,
        noListings: false,
        noAuctions: false,
    }),
    mounted() {
      // Wait for socket connection
      this.filters = JSON.parse(localStorage.getItem("filters"))
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
        },
        getListings() {
            let data = {
                body: this.filters,
            }
            socket.emit(listingsRoute, data, (data) => {
                if (!data.body.listings || data.body.listings.length == 0) {
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
            let data = {
                body: this.filters,
            }
            socket.emit(auctionsRoute, data, (data) => {
                if (!data.body.auctions || data.body.auctions.length == 0) {
                  this.noAuctions = true; 
                  return
                }
                for (let i = 0; i < data.body.auctions.length; i++) {
                  data.body.auctions[i].icon = data.body.icons[i];
                }
                this.auctions = data.body.auctions;
            });
        },
    },
    components: { Listing, Auction }
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