
<template>
  <div>
    <v-card
      variant="tonal"
      @click="locatorButtonPressed()"
      class="text-left pa-4 d-flex justify-space-between align-center"
    >
      <section class="d-flex justify-left align-center">
        <v-icon icon="mdi-store"></v-icon>
        <section>
          <v-card-title>{{shop.shopName}}</v-card-title>
          <v-card-text>{{shop.shopAddress}}, {{shop.ratings}}</v-card-text>
        </section>
      </section>
      <v-icon icon="mdi-phone" class="px-8"></v-icon>
    </v-card>
    <v-spacer class="mt-3"></v-spacer>
    <v-divider class="my-3"></v-divider>
    <v-card flat>
      <v-card-title>Detail</v-card-title>
      <v-card-text>
        ({{location.lat+', '+location.lng}})
        <br />
        {{car.plate+' '+car.brand}}
        <br />
        <p class="d-inline" v-for="p in problems" :key="p">{{p+' '}}</p>
        <br />
      </v-card-text>
      <div class="text-center">
        <v-card-action>
          <v-btn text color="red" variant="tonal" @click="reveal = false">Cancel</v-btn>
        </v-card-action>
      </div>
    </v-card>
  </div>
</template>
<script>
export default {
  props: {
    shop: {
      type: Object
    },
    car: {
      type: Object
    },
    location: {
      type: Object
    },
    problems: {
      type: Array
    }
  },
  data() {
    return {};
  }
};
</script>