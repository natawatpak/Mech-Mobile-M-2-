<template>
  <div class="text-center ma-4">
    <v-row no-gutters>
      <v-col cols="8">
        <v-badge
          color="#C0294A"
          :content="orders"
          :value="orders"
          class="text-h4 pa-4"
          overlap
        >
          Incoming order
        </v-badge>
        <v-card variant="tonal" class="text-left pa-4 ma-4">
          <v-card-title>Order {{details.id}} | {{details.name}}</v-card-title>
          <section>
            <v-card-text class="text-h7">
              Car {{details.type}} | {{details.brand}}
              <br> Problems {{details.problem}}
              <br> {{details.location}} km
            <v-row justify="end">
            <v-card-action>
              <v-btn
                class="mx-1"
                variant="tonal"
                color="error"
              >
                Decline
              </v-btn>

              <v-btn
                class="mx-1"
                color="blue"
                variant="tonal"
              >
                Accept
              </v-btn>
            </v-card-action>
          </v-row>
          </v-card-text>          
        </section>
      </v-card>
        
    
    <v-tabs
    v-model="tab"
      grow
      class="ma-4"
    >
      <v-tab value="one">On process</v-tab>
      <v-tab value="two">Finish</v-tab>
    </v-tabs>

    <v-card-text>
      <v-window v-model="tab">
        <v-window-item value="one">
          <v-card variant="tonal" class="text-left">
          <v-card-title>Order {{details.id}} | {{details.name}}</v-card-title>
          <section>
            <v-card-text class="text-h7">
              Car {{details.type}} | {{details.brand}}
              <br> Problems {{details.problem}}
              <br> {{details.location}} km
          </v-card-text>          
        </section>
      </v-card>
        </v-window-item>

        <v-window-item value="two">
          Two
        </v-window-item>

      </v-window>
    </v-card-text>

    </v-col>

    <v-col cols="4">
      <v-container class="bg-purple">
        <v-card></v-card>
      </v-container>
    </v-col>
    </v-row>
  </div>
</template>

<style>
.text-h4 {
  color: #000000;
  display: flex;
  justify-content: left;
  align-content: center;
}

</style>

<script>
export default {
  data() {
    return {
      items: [
          { order: '0123456', name: 'Nunnapat Kriengchaiyaprug' },
        ],
      orders: 1,
      details: { id:"12345", type: "SUV", brand: "MG", name: "nunnapat", problem: "2", location: "2.1" },
      tab: null, 
    };
  }
};
</script>