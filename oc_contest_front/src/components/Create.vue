<template>
  <!-- Title -->
  <v-container>
    <h1 class="text-center">Creation form</h1>
    <v-row class="mt-5">
      <v-col cols="3" height="100%" class="d-flex flex-column">
          <v-text-field v-model="s.name" label="Name" counter="30" required></v-text-field>
          <v-textarea v-model="s.short_desc" counter="150" label="Short Description" required></v-textarea>
      </v-col>
      <v-col cols="9" height="100%" class="d-flex flex-column">
        <v-textarea height="100%" v-model="s.long_desc" counter="5000" label="Long Description" required></v-textarea>
      </v-col>
      <v-col cols="6" height="100%" class="d-flex flex-column">
        <v-file-input v-model="icon" @change="changeIcon" label="Icon" accept="image/*" prepend-icon="mdi-camera" required></v-file-input>
      </v-col>
      <v-col cols="6" height="100%" class="d-flex flex-column">
        <v-file-input v-model="files" @change="changeFiles" label="More images" multiple accept="image/*" prepend-icon="mdi-file" required></v-file-input>
      </v-col>
      <v-col cols="12" class="text-center">
        <v-btn color="success" class="mr-4" @click="sendForm()">Submit</v-btn>
      </v-col>
    </v-row>
  </v-container>

  <v-container>
    <h1 class="text-center">Render</h1>
    <v-row class="mt-5" justify="center">
      <v-col cols="3">
        <submission :submission="s"/>
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
import submission from '@/components/SubmissionRender.vue'
import {backend} from '@/plugins/axios';

export default {
  name: 'HomePage',

  components: {
    submission
  },

  data: () => ({
    alert: {
      active:false,
      text: '',
    },
    user: {},
    s: {
      icon: {},

    },
    icon:undefined,
    files:undefined,
  }),

  created() {
    this.getUser().then((d) => {
      let author = "@" + d.RawData.username + "#" + d.RawData.discriminator;
      this.s.author = author;
    });
  },

  methods: {
    changeIcon() {
      this.s.icon = {
        uri: URL.createObjectURL(this.icon[0]),
      }
    },
    changeFiles() {
      this.s.files = [];
      for (let i = 0; i < this.files.length; i++) {
        this.s.files.push({
          name: this.files[i].name,
          uri: URL.createObjectURL(this.files[i]),
        });
      }
      console.log(this.s.files);
    },
    async getUser() {
      const response = await backend.get('/auth/status');
      return response.data.user;
    },
    sendForm() {
      let submission = this.s;
      submission.files = undefined;
      submission.icon = undefined;

      let data = new FormData()
      data.append("data", JSON.stringify(this.s));
      if (this.icon) {
        data.append("files", this.icon[0]);
      }
      if (this.files) {
        for (let i = 0; i < this.files.length; i++) {
          data.append("files", this.files[i]);
        }
      }

      backend.post('/submissions/create', data, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(() => {
        // Redirect to own submissions
        this.$router.push('/self');
      }).catch((e) => {
        this.alert.active = true;
        this.alert.text = e.response.data.error;
      });
    },
  },
  
}
</script>