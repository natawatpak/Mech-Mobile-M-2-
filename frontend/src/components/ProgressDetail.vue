
<template>
  <div>
    <v-card
      @click="locatorButtonPressed()"
      class="text-left pa-4 d-flex justify-space-between align-center"
      color="blue-darken-1"
    >
      <v-row class="align-center justify-space-between px-4">
        <section class="d-flex justify-left align-center">
          <v-icon icon="mdi-store"></v-icon>
          <section>
            <v-card-title>{{shop.shopName}}</v-card-title>
            <v-card-text>{{shop.shopAddress}}, {{shop.ratings}}</v-card-text>
          </section>
        </section>
        <v-icon icon="mdi-phone" class="px-8"></v-icon>
      </v-row>
    </v-card>
    <v-spacer class="mt-3"></v-spacer>
    <v-divider class="my-3"></v-divider>
    <v-card flat class="rounded-lg elevation-4">
      <v-card-title>Detail</v-card-title>
      <v-card-text>
        Location: {{shop.location.lat+', '+shop.location.lng}}
        <br />
        Car brand: {{car.brand}}
        <br />
        Car license plate: {{car.plate}}
        <br />
        Problems
        <br />
        <p v-for="p in problems" :key="p">{{'- '+ p}}</p>
        <br />
      </v-card-text>
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