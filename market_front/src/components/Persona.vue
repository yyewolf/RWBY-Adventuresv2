<template>
    <AdvancedSearch :open="dialogSearch" @close="dialogSearch = false;" :default_name="data.Name" ></AdvancedSearch>
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
            <v-btn variant="outlined" color="secondary" @click="dialogSearch = true;">Search</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
import AdvancedSearch from '@/components/AdvancedSearch.vue';

export default {
    name: "r-persona",
    components: { AdvancedSearch },
    data () {
        return {
            dialogSearch: false,
        }
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