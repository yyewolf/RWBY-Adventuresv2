<template>
    <v-dialog v-model="activate">
        <v-card width="90vw">
            <v-card-title>
                <span class="headline">Advanced Search</span>
            </v-card-title>
            <v-card-text>
                <v-container>
                <v-text-field v-model="filters.name_has" label="Name (Must contain)"></v-text-field>

                <p class="mb-2">Value :</p>
                <v-range-slider v-model="value_range" :max="100" :min="0" :step="0.01" hide-details class="mb-2 align-center" @update:modelValue="slideSlides('value', $event)">
                    <template v-slot:prepend>
                    <v-text-field :model-value="value_range[0]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('value_range', 0, $event)"></v-text-field>
                    </template>
                    <template v-slot:append>
                    <v-text-field :model-value="value_range[1]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('value_range', 1, $event)"></v-text-field>
                    </template>
                </v-range-slider>

                <p class="mb-2">Level :</p>
                <v-range-slider v-model="level_range" :max="500" :min="1" :step="1" hide-details class="mb-2 align-center" @update:modelValue="slideSlides('level', $event)">
                    <template v-slot:prepend>
                    <v-text-field :model-value="level_range[0]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('level_range', 0, $event)"></v-text-field>
                    </template>
                    <template v-slot:append>
                    <v-text-field :model-value="level_range[1]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('level_range', 1, $event)"></v-text-field>
                    </template>
                </v-range-slider>

                <p class="mb-2">Arms / Minions :</p>
                <v-range-slider v-model="buffs_range" :max="2" :min="0" :step="1" hide-details class="mb-2 align-center" @update:modelValue="slideSlides('buffs', $event)">
                    <template v-slot:prepend>
                    <v-text-field :model-value="buffs_range[0]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('buffs_range', 0, $event)"></v-text-field>
                    </template>
                    <template v-slot:append>
                    <v-text-field :model-value="buffs_range[1]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('buffs_range', 1, $event)"></v-text-field>
                    </template>
                </v-range-slider>

                <p class="mb-2">Rarity :</p>
                <v-range-slider v-model="rarity_range" :max="5" :min="0" :step="1" hide-details class="mb-2 align-center" @update:modelValue="slideSlides('rarity', $event)">
                    <template v-slot:prepend>
                    <v-text-field :model-value="rarity_range[0]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('rarity_range', 0, $event)"></v-text-field>
                    </template>
                    <template v-slot:append>
                    <v-text-field :model-value="rarity_range[1]" hide-details type="number" variant="outlined" style="width: 20vw" density="compact" @change="update('rarity_range', 1, $event)"></v-text-field>
                    </template>
                </v-range-slider>

                <p class="mb-2">Order by :</p>
                <v-row>
                    <v-col cols="12" sm="6">
                    <v-select v-model="order_by" :items="['Price', 'Value', 'Level', 'Rarity', 'Date', 'Type']" label="I will order by date by default." @update:modelValue="updateSelect"></v-select>
                    </v-col>
                    <v-col cols="12" sm="6">
                    <v-switch v-model="orderState" hide-details inset :label="orderState ? 'Ascending (Lowest first)' : 'Descending (Highest first)'" @change="updateOrderType"></v-switch>
                    </v-col>
                </v-row>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" @click="$emit('close')">Close</v-btn>
                <v-btn color="primary" @click="search();">Search</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
  name: 'App',

  components: {},

  props:["open", "default_name"],

  data: () => ({
    value_range: [0, 100],
    level_range: [1, 500],
    buffs_range: [0, 2],
    rarity_range: [0, 5],
    order_by: undefined,
    orderState: false,
    filters: {
      name_has: '',
      value_above: '0',
      value_below: '100',
      level_above: '1',
      level_below: '500',
      buffs_above: '0',
      buffs_below: '2',
      rarity_above: '0',
      rarity_below: '5',
      order_by: '',
      order_type: '',
    },
    login_link: undefined,
    logged_in: false,

    activate: false,
  }),
  
  watch: {
    activate: function (n, o) {
        o;
        if (n == false) {
            this.$emit('close');
        }
    },
    open: function (n, o) {
        o;
        this.activate = n;
    },
  },

  mounted() {
    if (this.default_name) {
      this.filters.name_has = this.default_name;
    }
  },

  methods : {
    update: function(val, i, e) {
      this[val][i] = e.target.value;
    },
    updateSlides: function(val, i, e) {
      let filter = val.split('_')[0] + (i ? '_below' : '_above');
      this.filters[filter] = e.target.value;
    },
    slideSlides: function(val, e) {
      this.filters[val+"_above"] = e[0].toString();
      this.filters[val+"_below"] = e[1].toString();
    },
    updateSelect: function(e) {
      switch(e) {
        case 'Price':
          this.filters.order_by = 'price';
          break;
        case 'Value':
          this.filters.order_by = 'value';
          break;
        case 'Level':
          this.filters.order_by = 'level';
          break;
        case 'Rarity':
          this.filters.order_by = 'rarity';
          break;
        case 'Date':
          this.filters.order_by = 'id';
          break;
        case 'Type':
          this.filters.order_by = 'type';
          break;
      }
    },
    updateOrderType: function() {
      this.filters.order_type = this.orderState ? 'asc' : 'desc';
    },
    search: function() {
      this.$emit('close');
      localStorage.setItem('filters', JSON.stringify(this.filters))
      this.$router.push({
        name: 'Search',
      });
    },
  }
}
</script>
