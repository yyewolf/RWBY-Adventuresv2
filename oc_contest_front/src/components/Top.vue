<template>
  <!-- Title -->
  <v-container>
    <h1 class="text-center"> Top </h1>
    <h3 v-if="status.logged" class="text-center"> You have {{5-status.votes}} votes left. </h3>
    <v-row class="mt-5" justify="center" v-for="s in submissions" :key="s">
      <v-col cols="12" xs="12" sm="12" md="6" lg="4" xl="3">
        <submission :submission="s" @vote="submissionVote(s)"/>
      </v-col>
      <v-col cols="3" v-if="submissions.length == 0">
        There are no submissions yet...
      </v-col>
    </v-row>
  </v-container>
  <v-snackbar v-model="alert.active" :timeout="2000">
    {{ alert.text }}

    <template v-slot:actions>
      <v-btn color="blue" variant="text" @click="alert.active = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
import submission from '@/components/Submission.vue'
import {backend} from '@/plugins/axios';

export default {
  name: 'HomePage',

  props: {
    status: Object,
  },

  components: {
    submission
  },

  data: () => ({
    alert: {
      active:false,
      text: '',
    },
    submissions: [],
    votes: 0,
  }),

  mounted() {
    this.getSubmissions(this.page-1);

    if (this.status.logged) {
      this.votes = this.status.votes;
    }
  },

  methods: {
    onPageChange() {
      this.getSubmissions(this.page-1);
    },
    submissionVote(s) {
      backend.get('/submissions/vote/'+s.SubmissionID).then(() => {
        this.alert.active = true;
        this.alert.text = 'Vote successful!';
        this.votes++;
        s.votes.push({});
      }).catch((e) => {
        this.alert.active = true;
        this.alert.text = e.response.data.error;
      });
    },
    async getSubmissions() {
      const response = await backend.get('/submissions/top');
      console.log(response);
      this.submissions = response.data.submissions;
      this.pages = response.data.max_page;
    },
  },
  
}
</script>
