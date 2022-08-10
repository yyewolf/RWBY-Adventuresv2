<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12" sm="4">
        <v-card elevation="2" outlined height="100%">
          <v-card-text>
            <p class="text-h4">
              Latest auctions <v-btn variant="plain" href="/about">(View all)</v-btn>
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="4">
        <v-card elevation="2" outlined height="100%">
          <v-card-text>
            <p class="text-h4">
              Random Characters
            </p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="4">
        <v-card elevation="2" outlined height="100%">
          <v-card-text>
            <p class="text-h4">
              Latest listings <v-btn variant="plain" href="/about">(View all)</v-btn>
            </p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'HomePage',

  data: () => ({}),
}
</script>
