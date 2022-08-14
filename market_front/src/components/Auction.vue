<template>
    <v-dialog v-model="confirmation">
        <v-card>
            <v-card-title>
                <span class="headline">Place your bid!</span>
            </v-card-title>
            <v-card-text class="d-flex" style="align-items: center;">
                <v-text-field
                    v-model="bidding"
                    hide-details
                    single-line
                    density="comfortable"
                    placeholder="How much ?"
                    variant="filled"
                    label="How much ?"
                    type="number"
                />
                <p class="ml-3">Liens</p>
            </v-card-text>
            <v-card-actions>
                <v-btn color="primary" @click="confirmBid(); confirmation = false;">Confirm</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
    <v-card variant="outlined" :style="'border-color:'+this.border" width="200" height="100%">
        <div>
            <v-avatar class="mt-3" size="100" rounded>
                <v-img :src="data.icon"></v-img>
            </v-avatar>
        </div>
        <v-card-text align="left">
            <p class="title">{{name}}</p>
            <p class="author" v-if="data.seller_name != undefined">@{{data.seller_name}}</p>
            <p class="star mt-2">
                <v-icon color="yellow" v-for="i in 5" :key="i">{{i < this.rarity ? "mdi-star" : "mdi-star-outline"}}</v-icon> <v-icon color="green" v-for="i in buffs" :key="i" :id="i">mdi-plus</v-icon>
            </p>
            <p class="ml-1">{{rarityString}}</p>
            <p ref="textanim" :class="animate ? '' : 'grow'">
                <v-icon color="green">mdi-cash</v-icon> {{price}}Ⱡ
            </p>
            <p>
                <v-icon>mdi-percent</v-icon> {{value}}%
            </p>
            <p>
                <v-icon>mdi-arrow-up-bold</v-icon> Level {{level}}
            </p>
        </v-card-text>
        <v-card-text>
            <p class="mb-2">{{time}}</p>
            <v-btn color="secondary" @click="confirmation = true">Bid</v-btn>
        </v-card-text>
    </v-card>
</template>

<script>
import socket from '@/plugins/websocket';

const auctionsBidRoute = 'auctions/bid';

export default {
    name: "r-auction",
    data () {
        return {
            name: undefined,
            border: undefined,
            rarity: undefined,
            value: undefined,
            level: undefined,
            price: undefined,
            time: undefined,
            ends_at: undefined,
            rarityString: undefined,
            buffs: 0,

            confirmation: false,
            bidding: 0,

            animate: false,
        }
    },
    props: ['data'],
    mounted() {
        let persona = this.data.char || this.data.grimm
        this.name = persona.Name
        this.rarity = persona.Rarity
        this.value = persona.Stats.Value
        this.level = persona.Level
        this.buffs = persona.Buffs
        this.ends_at = this.data.ends_at
        if (this.data.bidders.length > 0) {
            this.price = this.data.bidders[0].amount
        } else {
            this.price = 0
        }
        if (this.data.type == 0) {
            this.border = this.charColor();
            this.rarityString = this.charText();
        } else {
            this.border = this.grimmColor();
            this.rarityString = this.grimmText();
        }

        socket.on(this.data.ID, (data) => {
            this.price = data.body.amount;
            this.ends_at = data.body.ends_at;

            this.animate = true;
            this.$refs.textanim.offsetWidth;
            setTimeout(() => {
                this.animate = false;
            }, 100);
        })

        this.countDownTimer()
    },

    methods: {
        charColor() {
            switch (this.rarity) {
                case 0: // Common
                    return "#808080"
                case 1: // Uncommon
                    return "#7CFC00"
                case 2: // Rare
                    return "#87CEEB"
                case 3: // Very Rare
                    return "#BA55D3"
                case 4: // Legendary
                    return "#FFD700"
                case 5: // Collector
                    return "#FF0000"
            }
        },
        grimmColor() {
            switch (this.rarity) {
                case 0: // Common
                    return "#808080"
                case 1: // Uncommon
                    return "#285300"
                case 2: // Rare
                    return "#00008b"
                case 3: // Very Rare
                    return "#B22222"
                case 4: // Legendary
                    return "#800080"
                case 5: // Collector
                    return "#121212"
            }
        },
        charText() {
            switch (this.rarity) {
                case 0: // Common
                    return "Common"
                case 1: // Uncommon
                    return "Uncommon"
                case 2: // Rare
                    return "Rare"
                case 3: // Very Rare
                    return "Very Rare"
                case 4: // Legendary
                    return "Legendary"
                case 5: // Collector
                    return "Collector"
            }
        },
        grimmText() {
            switch (this.rarity) {
                case 0: // Common
                    return "Normal"
                case 1: // Uncommon
                    return "Abnormal"
                case 2: // Rare
                    return "Sparse"
                case 3: // Very Rare
                    return "Freaky"
                case 4: // Legendary
                    return "Mysterious"
                case 5: // Collector
                    return "Bloody"
            }
        },
        countDownTimer() {
            this.calcTime()
            setInterval(() => {
                this.calcTime()
            }, 1000)
        },
        calcTime() {
            let t = new Date(new Date(this.ends_at * 10000) - Date.now())
            this.time = "Ends in " + t.getHours()*t.getDay() + "h " + t.getMinutes() + "m " + t.getSeconds() + "s"
            if (t < 0) {
                this.time = "This auction is over."
            }
        },
        confirmBid() {
            let data = {
                body: {
                    auction_id: this.data.ID,
                    amount: this.bidding,
                },
            }
            socket.emit(auctionsBidRoute, data, (res) => {
                this.confirmation = false
                this.bidding = 0
                this.$notify({
                    text: res.text,
                });
            })
            this.confirmation = false
        }
    },
}
</script>

<style scoped>

/* ----------------------------------------------
 * Generated by Animista on 2022-8-14 12:6:41
 * Licensed under FreeBSD License.
 * See http://animista.net/license for more info. 
 * w: http://animista.net, t: @cssanimista
 * ---------------------------------------------- */

/**
 * ----------------------------------------
 * animation pulsate-fwd
 * ----------------------------------------
 */
@-webkit-keyframes pulsate-fwd {
  0% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  50% {
    -webkit-transform: scale(1.3);
            transform: scale(1.3);
  }
  100% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
}
@keyframes pulsate-fwd {
  0% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  50% {
    -webkit-transform: scale(1.3);
            transform: scale(1.3);
  }
  100% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
}

.grow {
    -webkit-animation: pulsate-fwd 0.5s cubic-bezier(0.230, 1.000, 0.320, 1.000) both;
            animation: pulsate-fwd 0.5s cubic-bezier(0.230, 1.000, 0.320, 1.000) both;
}

.title {
    font-size: 1rem;
    font-weight: bold;
}

.author {
    font-size: 0.8rem;
    font-style: italic;
}

.star {
    font-size:0.7rem
}
</style>