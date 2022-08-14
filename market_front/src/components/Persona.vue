<template>
    <v-card class="d-flex flex-column" height="100%" variant="outlined" width="175">
        <div>
            <v-avatar class="mt-3" size="75" rounded>
                <v-img :src="data.IconURL"></v-img>
            </v-avatar>
        </div>
        <v-card-text>
            <p class="title">{{data.Name}}</p>
        </v-card-text>
        <v-card-actions class="justify-center">
            <v-btn class="flex-end" variant="outlined" color="secondary">Search</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
export default {
    name: "r-persona",
    data () {
        return {}
    },
    props: ['data'],
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
</style>