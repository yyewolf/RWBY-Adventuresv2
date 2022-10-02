<template>
  <!-- Title -->
  <v-container>
    <v-row class="mt-5" justify="center">
      <v-col cols="1">
        <h1> Submissions </h1>
      </v-col>
    </v-row>
    <v-row class="mt-5">
      <v-col cols="3" v-for="s in submissions" :key="s">
        <submission :submission="s"/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import submission from '@/components/Submission.vue'
export default {
  name: 'HomePage',

  components: {
    submission
  },

  data: () => ({
    submissions: [
      {
        name: "Submission Name",
        color: "Color",
        short_desc: "This is a short description of the submission.",
        long_desc: "This is a long description of the submission.",
        icon: "https://img.rwbyadventures.com/Janina_Wolf/Xmas_Icon.webp",
        author: "@Someone#0000",
        votes: 51,
        files:[
            {
                name:"Image name.png",
                uri:"image file.png"
            }
        ],
      },
    ],
  }),
  
}
</script>
