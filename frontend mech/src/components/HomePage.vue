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
          <v-content>
            <v-container>
              <v-layout column wrap>
                <v-flex v-for="item in items" :key="item.id">
                  <v-card width="100%" variant="tonal" class="text-left pa-4">
                    <v-row class="pa-2">
                      <v-card-title>
                        Order {{item.id}} | {{item.username}}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow">{{item.status}}</v-chip>
                    </v-row>
                    <v-card-text class="text-h7">
                      Car {{details.type}} | {{details.brand}}
                      <br> Problems {{details.problem}}
                      <br> {{details.location}} km
                    </v-card-text>  
                  </v-card>
                </v-flex>
              </v-layout>
          </v-container>
        </v-content>
        </v-window-item>

        <v-window-item value="two">
          <v-card variant="tonal" class="text-left my-4 pa-4">
            <v-row class="pa-2">
              <v-card-title>Order {{details.id}} | {{details.name}}</v-card-title>
              <v-chip class="ma-2" color="green">{{details.status}}</v-chip>
            </v-row>
          <section>
            <v-card-text class="text-h7">
              Car {{details.type}} | {{details.brand}}
              <br> Problems {{details.problem}}
              <br> {{details.location}} km
          </v-card-text>          
        </section>
      </v-card>
        </v-window-item>

      </v-window>
    </v-card-text>

    </v-col>
    <v-col cols="4">
      <v-container class="bg-purple" height="100%">
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
      orders: 1,
      details: { id:"12345", type: "SUV", brand: "MG", name: "nunnapat", problem: "2", location: "2.1"},
      items: [
        { id:"11111", username: "nunnapat", distance: "2.2", type: "SUV", brand: "MG", status: "On-process"},
        { id:"22222", username: "natawat", distance: "0.4", type: "Sedan", brand: "MG", status: "On the way"},
        { id:"33333", username: "kiitiphum", distance: "1.3", type: "Van", brand: "MG", status: "Accept"}
      ],
      tab: null, 
    };
  }
};
</script>