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
  
    created() {
      this.getSubmissions();
    },
  
    methods: {
      async getSubmissions() {
        const response = await backend.get('/submissions/current');
        console.log(response);
        this.submissions = response.data.submissions;
        this.pages = response.data.max_page;
      },
    },
    
  }
  </script>
  