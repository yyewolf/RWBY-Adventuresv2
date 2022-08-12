<template>
  <v-app>
    <v-app-bar density="compact">
      <template v-slot:prepend>
        <v-tab to="/">RWBYzon</v-tab>
      </template>

      <v-row class="justify-center">
        <v-col cols="12" sm="6">
          <v-text-field 
            hide-details 
            append-icon="mdi-magnify" 
            single-line 
            outlined

            label="Search"
            density="compact"
            variant="outlined"
          ></v-text-field>
        </v-col>
      </v-row>

      <template v-slot:append>
        <v-btn prepend-icon="mdi-discord">Log in</v-btn>
      </template>
    </v-app-bar>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  components: {},

  data: () => ({
    //
  }),
}
</script>
