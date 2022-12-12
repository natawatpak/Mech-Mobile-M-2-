
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
      <v-divider></v-divider>
      <v-card-text>
        <div class="d-inline">
          <v-text class="font-weight-medium text-subtitle-1">Location:</v-text> 
          {{shop.location.lat+', '+shop.location.lng}}
        </div>
        <br />
        <div class="d-inline">
          <v-text class="font-weight-medium text-subtitle-1">Car brand:</v-text> 
          {{car.brand}}
        </div>
        <br />
        <div class="d-inline">
          <v-text class="font-weight-medium text-subtitle-1">Car license plate:</v-text> 
          {{car.plate}}
        </div>
        <br />
        <div class="d-inline" v-if="ps.length">
          <v-text class="font-weight-medium text-subtitle-1">Problems:</v-text> 
          <li v-for="p in ps.toString().split(',')" :key="p">{{ p }}</li>
        </div>
      </v-card-text>
    </v-card>
  </div>
</template>
<script>
export default {
  data(){
    return{
      ps: this.problems
    }
  },
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
};
</script>