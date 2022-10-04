<template>
  <!-- Title -->
  <v-container>
    <v-row class="mt-5" justify="center">
      <v-col cols="2">
        <h1> Your submissions </h1>
      </v-col>
    </v-row>
    <v-row class="mt-5">
      <v-col cols="3" v-for="s in submissions" :key="s">
        <v-btn v-bind="props" variant="text" @click="deleteSubmission(s)">
          <v-icon large color="red darken-2">
              mdi-delete
          </v-icon>
        </v-btn>
        <submission :submission="s"/>
      </v-col>
      <v-col cols="3" v-if="submissions.length == 0">
        There are no submissions yet...
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import submission from '@/components/Submission.vue'
import {backend} from '@/plugins/axios';

export default {
  name: 'HomePage',

  components: {
    submission
  },

  data: () => ({
    submissions: [],
  }),

  mounted() {
    this.getSubmissions();
  },

  methods: {
    deleteSubmission(s) {
      backend.get('/submissions/delete/'+s.SubmissionID).then(() => {
        this.getSubmissions();
      });
    },
    async getSubmissions() {
      const response = await backend.get('/submissions/current');
      console.log(response);
      this.submissions = response.data.submissions;
      this.pages = response.data.max_page;
    },
  },
  
}
</script>x