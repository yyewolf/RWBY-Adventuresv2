<template>
  <v-container>
    <v-row class="mb-3" justify="center">
      <h2>{{arena.title}}</h2>
    </v-row>
    <v-row class="mb-5" justify="center">
      <h4>There are {{arena.people}} people with you.</h4>
    </v-row>
    <v-row justify="center">
      <v-card width="20vw" @click="hit++">
        <v-card-text>
          <v-img src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/8bc5f32c-8fbf-4166-8729-b7d4d62cb35d/ddbvlct-86de9502-195a-4185-a02d-41432f9d19e9.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzhiYzVmMzJjLThmYmYtNDE2Ni04NzI5LWI3ZDRkNjJjYjM1ZFwvZGRidmxjdC04NmRlOTUwMi0xOTVhLTQxODUtYTAyZC00MTQzMmY5ZDE5ZTkucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.593heLTyYUe6rYeftcflrOjPyGVBnopJCEl_gfllkAo"/>
        </v-card-text>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'ArenaPage',
  props: ["arena"],
  data: () => ({
    hit:0,
  }),
}
</script>
