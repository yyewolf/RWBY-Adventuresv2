<template>
    <v-card height="100%" tonal @click="dialog = true" class="d-flex flex-column justify-space-between">
        <div class="d-flex flex-no-wrap justify-space-between">
            <div>
                <v-card-text>
                    <h2>{{submission.name}}</h2>
                </v-card-text>
                <v-card-subtitle>
                    {{submission.color}}
                </v-card-subtitle>
                <v-card-text>
                    {{submission.short_desc}}
                </v-card-text>
            </div>
            <v-avatar class="ma-3" size="125" rounded="0">
                <v-img :src="filepath+submission.icon.uri"></v-img>
            </v-avatar>
        </div>

        <v-card-actions>
            <v-row justify="center">
                <v-col cols="1">
                </v-col>
                <v-col cols="5" class="text-center">
                    <p class="author">Author : <b>{{submission.author}}</b></p>
                </v-col>
                <v-col cols="6" class="text-center">
                    <v-icon>mdi-thumb-up</v-icon>
                    <p class="ml-3">{{submission.votes.length}} Votes</p>
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
                    <v-card tonal v-if="dialog" width="500px" class="center">
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
                                    <v-img height="50px" :src="filepath+file.uri" class="image" @click="image_dialog = true; image = filepath+file.uri" v-bind="props"/>
                                </template>
                                <span>{{file.name}}</span>
                            </v-tooltip>
                        </v-col>
                    </template>
                </v-row>
            </v-card-text>
            <v-card-actions>
                <v-btn color="primary" @click="clickVote()">
                    <v-icon>mdi-thumb-up</v-icon>
                    <p class="ml-3">{{submission.votes.length}} Votes</p>
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
  emits: ['vote'],

  data: () => ({
    filepath: process.env.VUE_APP_BACKEND,
    image_dialog: false,
    image: undefined,
    dialog: false,
  }),

  methods: {
    clickVote() {
        this.$emit('vote');
    }
  }
}
</script>

<style scoped>
.author {
    color:rgb(249, 231, 255);
}

.image {
    background-color: rgba(245, 245, 220, 0);
    border-radius: 5px;
}

.notif {
    /* higher font size */
    font-size: 1.5em !important;
}

.center {
    position:absolute;
    top:50%;
    left:50%;
    transform:translate(-50%, -50%);
}
</style>