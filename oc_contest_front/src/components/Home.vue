<template>
  <!-- Title -->
  <v-container>
    <h1 class="text-center"> Submissions </h1>
    <h3 v-if="status.logged" class="text-center"> You have {{5-(votes > status.votes ? votes : status.votes)}} votes left. </h3>
    <v-row class="mt-5">
      <v-col cols="12" xs="12" sm="12" md="6" lg="4" xl="3" v-for="s in submissions" :key="s">
        <submission :submission="s" @vote="submissionVote(s)"/>
      </v-col>
      <v-col cols="3" v-if="submissions.length == 0">
        There are no submissions yet...
      </v-col>
    </v-row>
    <div class="text-center">
      <v-pagination
        v-model="page"
        :length="pages+1"
        @click="onPageChange"
      ></v-pagination>
    </div>
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
    pages: 0,
    page: 1,

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
    async getSubmissions(page) {
      const response = await backend.get('/submissions/all/'+page);
      this.submissions = response.data.submissions;
      this.pages = response.data.max_page;
    },
  },
  
}
</script>
