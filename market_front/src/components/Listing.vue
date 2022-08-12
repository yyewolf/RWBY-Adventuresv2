<template>
    <v-card variant="outlined" :style="'border-color:'+this.border">
        <div>
            <v-avatar class="mt-3" size="100" rounded>
                <v-img :src="data.icon"></v-img>
            </v-avatar>
        </div>
        <v-card-text align="left">
            <p class="title">{{name}}</p>
            <p class="author" v-if="data.seller_name != undefined">@{{data.seller_name}}</p>
            <p class="star mt-2">
                <v-icon color="yellow" v-for="i in 5" :key="i">{{i < this.rarity ? "mdi-star" : "mdi-star-outline"}}</v-icon>
                
            </p>
            <p class="ml-1">{{rarityString}}</p>
            <p>
                <v-icon color="green">mdi-cash</v-icon> {{data.price}}Ⱡ
            </p>
            <p>
                <v-icon>mdi-percent</v-icon> {{value}}%
            </p>
            <p>
                <v-icon>mdi-arrow-up-bold</v-icon> Level {{level}}
            </p>
        </v-card-text>
        <v-card-text>
            <v-btn color="secondary">Purchase</v-btn>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    name: "r-listing",
    data () {
        return {
            name: undefined,
            border: undefined,
            rarity: undefined,
            value: undefined,
            level: undefined,
            rarityString: undefined,
        }
    },
    props: ['data'],
    mounted() {
        console.log(this.data)
        this.name = this.data.char.Name || this.data.grimm.Name;
        this.rarity = this.data.char.Rarity || this.data.grimm.Rarity;
        this.value = this.data.char.Stats.Value || this.data.grimm.Stats.Value;
        this.level = this.data.char.Level || this.data.grimm.Level;
        if (this.data.type == 0) {
            this.border = this.charColor();
            this.rarityString = this.charText();
        } else {
            this.border = this.grimmColor();
            this.rarityString = this.grimmText();
        }
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
    },
}
</script>

<style scoped>
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