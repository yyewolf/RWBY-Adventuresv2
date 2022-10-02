<template>
    <v-card height="100%" tonal @click="dialog = true" class="d-flex flex-column justify-space-between">
        <div class="d-flex flex-no-wrap justify-space-between">
            <div>
                <v-card-title>
                    {{submission.name}}
                </v-card-title>
                <v-card-subtitle>
                    {{submission.color}}
                </v-card-subtitle>
                <v-card-text>
                    {{submission.short_desc}}
                </v-card-text>
            </div>
            <v-avatar class="ma-3" size="125" rounded="0">
                <v-img :src="submission.icon"></v-img>
            </v-avatar>
        </div>

        <v-card-actions>
            <v-row justify="center">
                <v-col cols="6" class="text-center">
                    <p class="author">Author : <b>{{submission.author}}</b></p>
                </v-col>
            </v-row>
        </v-card-actions>
    </v-card>
    <v-dialog v-model="dialog" max-width="65vw">
        <v-card tonal v-show="dialog" width="65vw" class="d-flex flex-column">
            <v-card-title style="font-size: 1.35rem;">
                Description
            </v-card-title>
            <v-card-text style="font-size: 1rem;">
                <v-row align="center" justify="center" style="height:100%">
                    {{submission.long_desc}}
                </v-row>
                <v-dialog v-model="image_dialog">
                    <v-card tonal v-if="dialog" width="500px">
                        <v-img :src="image" class="image"/>
                    </v-card>
                </v-dialog>
            </v-card-text>
            <v-card-text style="flex:none;">
                <v-row justify="center" class="mb-1">
                    <template v-for="file in submission.files" :key="file">
                        <v-col cols="2">
                            <v-tooltip location="top">
                                <template v-slot:activator="{ props }">
                                    <template v-if="file.name.split('.').pop() == 'jpeg' || file.name.split('.').pop() == 'jpg' || file.name.split('.').pop() == 'png'">
                                        <v-img height="50px" :src="file.uri" class="image" @click="image_dialog = true; image = file.uri" v-bind="props"/>
                                    </template>
                                    <template v-else>
                                        <v-btn height="100%" :href="file.uri" v-bind="props" variant="text">
                                            <v-icon large color="orange darken-2">
                                                mdi-file
                                            </v-icon>
                                        </v-btn>
                                    </template>
                                </template>
                                <span>{{file.name}}</span>
                            </v-tooltip>
                        </v-col>
                    </template>
                </v-row>
            </v-card-text>
            <v-card-actions>
                <v-btn color="primary">
                    <v-icon>mdi-thumb-up</v-icon>
                    <p class="ml-3">{{submission.votes}} Votes</p>
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
  name: 'SubmissionCard',

  props: {
    submission: {
      type: Object,
      required: true
    }
  },

  data: () => ({
    image_dialog: false,
    image: undefined,
    dialog: false,
  }),
  
  methods: {
  },
}
</script>

<style scoped>
.author {
    color:rgb(249, 231, 255);
}

.image {
    background-color: beige;
    border-radius: 5px;
}

.notif {
    /* higher font size */
    font-size: 1.5em !important;
}
</style>