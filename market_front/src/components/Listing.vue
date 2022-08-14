<template>
    <v-dialog v-model="confirmation">
        <v-card>
            <v-card-title>
                <span class="headline">Listing purchase</span>
            </v-card-title>
            <v-card-text class="d-flex" style="align-items: center;">
                <p>
                    Confirm transaction for <b> {{data.price}}</b>Ⱡ ?
                </p>
            </v-card-text>
            <v-card-actions>
                <v-btn color="primary" @click="confirmation = false;">NO</v-btn>
                <v-btn color="secondary" @click="confirmBuy(); confirmation = false;">YES</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
    <v-card variant="outlined" :style="'border-color:'+this.border" width="200" height="100%">
        <v-overlay :model-value="bought" scroll-strategy="allow" scrim="#000" contained class="align-center justify-center">
            <h1>SOLD</h1>
        </v-overlay>
        <div :class="bought ? 'blur' : ''">
            <div>
                <v-avatar class="mt-3" size="100" rounded>
                    <v-img :src="data.icon"></v-img>
                </v-avatar>
            </div>
            <v-card-text align="left">
                <p class="title">{{name}}</p>
                <p class="author" v-if="data.seller_name != undefined">@{{data.seller_name}}</p>
                <p class="star mt-2">
                    <v-icon color="yellow" v-for="i in 5" :key="i">{{i <= this.rarity ? "mdi-star" : "mdi-star-outline"}}</v-icon> <v-icon color="green" v-for="i in buffs" :key="i" :id="i">mdi-plus</v-icon>
                </p>
                <p class="ml-1">{{rarityString}}</p>
                <p>
                    <v-icon color="green">mdi-cash</v-icon> {{data.price}}Ⱡ
                </p>
                <p>
                    <v-icon>mdi-percent</v-icon> {{value.toFixed(2)}}%
                </p>
                <p>
                    <v-icon>mdi-arrow-up-bold</v-icon> Level {{level}}
                </p>
            </v-card-text>
            <v-card-text>
                <v-btn color="secondary" @click="confirmation = true" :disabled="bought">Purchase</v-btn>
            </v-card-text>
        </div>
    </v-card>
</template>

<script>
import socket from '@/plugins/websocket';

const listingsBuyRoute = 'listings/buy';

export default {
    name: "r-listing",
    data () {
        return {
            name: undefined,
            border: undefined,
            rarity: undefined,
            value: 0.0,
            level: undefined,
            rarityString: undefined,
            buffs: 0,

            confirmation: false,
            bought: false,
        }
    },
    props: ['data'],
    mounted() {
        let persona = this.data.char || this.data.grimm
        this.name = persona.Name;
        this.rarity = persona.Rarity;
        this.value = persona.Stats.Value;
        this.level = persona.Level;
        this.buffs = persona.Buffs;
        if (this.data.type == 0) {
            this.border = this.charColor();
            this.rarityString = this.charText();
        } else {
            this.border = this.grimmColor();
            this.rarityString = this.grimmText();
        }
        
        socket.on(this.data.ID, () => {
            this.bought = true
        })
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
        confirmBuy() {
            let data = {
                body: {
                    listing_id: this.data.ID,
                },
            }
            socket.emit(listingsBuyRoute, data, (res) => {
                this.confirmation = false
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
.blur {
    filter: blur(5px);
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