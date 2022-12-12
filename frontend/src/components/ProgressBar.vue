<template>
  <div>
    <v-timeline direction="horizontal" line-inset="12">
      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 'Processing' || currentState == 'Accepted' || currentState == 'On the way' || currentState == 'Finish:Garage' || currentState == 'Finish:Completed' ? 'green':'white'" :icon=items[0]>
        <template v-slot:opposite>Accepted</template>
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 'Processing' || currentState == 'On the way' || currentState == 'Finish:Garage' || currentState == 'Finish:Completed'? 'green':'white'" :icon=items[1]>
        On the way
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 'Processing' || currentState == 'Finish:Garage' || currentState == 'Finish:Completed'? 'green':'white'" :icon=items[2]>
        <template v-slot:opposite>On process</template>
      </v-timeline-item>
      <v-timeline-item size="large" v-model="items" :dot-color="currentState == 'Finish:Garage' || currentState == 'Finish:Completed'? 'green':'white'" :icon=items[3]>
        Complete
      </v-timeline-item>
    </v-timeline>

    <v-dialog v-if="currentState == 'Finish:Garage' || currentState == 'Finish:Completed'" v-model="dialog" class="text-left pa-4" width="800">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5">Service summary</v-card-title>
        <v-card variant="outlined" class="text-left ma-2 pa-1">
          <section>
            <v-card-title>Shop name: {{shop.name}}</v-card-title>
            <v-card-text>{{shop.address}}, {{shop.ratings}}</v-card-text>
          </section>
          <div class="text-center">
            <!-- <v-rating
              v-model="rating"
              clearable
            ></v-rating> -->
          </div>
        </v-card>
        <v-row justify="end" class="pa-4">
          <!-- <v-btn class="mx-1 " variant="tonal" color="error" @click="dialog=false">Cancel</v-btn> -->
          <v-btn to="/" class="mx-1 " variant="tonal" color="blue" @click="dialog=false;clearTicket()">Ok</v-btn>
        </v-row>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
export default {
  props:{
    currentState:{
      type: String
    }
  },
  data() {
    return {
      dialog: true,
      rating: 0,
      items: ['mdi-thumb-up', 'mdi-car', 'mdi-wrench', 'mdi-check'],
      shop: {
        name: "A",
        address: "567 Kingston Rd.",
        ratings: "4.5"
      },
    };
  },
  methods: {
    clearTicket(){
      sessionStorage.removeItem("ticketID")
      sessionStorage.removeItem("shopID")
    }
  }
};
</script>