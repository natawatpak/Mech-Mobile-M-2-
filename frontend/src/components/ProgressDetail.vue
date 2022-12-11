
<template>
  <div>
    <v-card
      @click="locatorButtonPressed()"
      class="text-left pa-4 d-flex justify-space-between align-center rounded-lg elevation-4"
      color="blue-darken-2"
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
        Location: {{location.lat+', '+location.lng}}
        <br />
        Car brand: {{car.brand}}
        <br />
        Car license plate: {{car.plate}}
        <br />
        Problems
        <br />
        <p v-for="p in problems" :key="p">{{'- ' + p}}</p>
      </v-card-text>
    </v-card>

    <div class="text-center pa-4">
        <v-card-action>
          <v-btn :disabled="valid" color="red" variant="tonal" @click="checkStage()">Cancel</v-btn>
        </v-card-action>
    </div>

    <v-dialog v-model="dialog" width="1000" class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5">Cancel order</v-card-title>  
        <v-card-text>
          <p>Why are you cancelling?</p>   
          <v-radio-group>
            <v-radio label="The problem was already fixed" value="fixed"></v-radio>
            <v-radio label="Found another mechanical" value="found"></v-radio>
            <v-radio label="No need to fix anymore" value="no need"></v-radio>
            <v-radio label="Others" value="other"></v-radio>
          </v-radio-group>
        </v-card-text>
        <v-row justify="end" class="pa-4">
          <v-card-action>
            <v-btn class="mx-1" variant="tonal" color="error" @click="dialog = false">
              Cancel
            </v-btn>

            <v-btn class="mx-1" color="blue" variant="tonal" @click="submitForm()">
              Submit
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialog2" width="1000" class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5">Unable to cancelling oreder</v-card-title>  
        <v-card-text>
          <div class="font-weight-bold text-subtitle-1 pb-2">It may be possible to cancel if the mechanical status DOES NOT REFLECT these statused yet</div>   
          <v-row class="pa-3">
          <li class="font-weight-bold text-subtitle-1 pr-2">On the way</li> <p class="text-subtitle-1">This means that the mechanical has already on the way to help you</p>
          </v-row>
          <v-row class="pa-3 pb-6">
          <li class="font-weight-bold text-subtitle-1 pr-2">In process</li> <p class="text-subtitle-1">This means that the mechanical is fixing your car now</p>
          </v-row>
          <div class="font-weight-bold text-subtitle-1">If you still want to cancelling the order, please contact with your mechanical.</div>   
        </v-card-text>
        <v-row justify="end" class="pa-4">
          <v-card-action>
            <v-btn class="mx-1" variant="tonal" color="blue" @click="dialog2 = false">
              Ok
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>
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
    },
    currentStage: {
      type: Number
    }
  },
  data() {
    return {
      valid: false,
      dialog: false,
      dialog2: false,
    };
  },
  methods: {
    checkStage() {
      if (this.currentStage == 1) {
        this.dialog = true;
      }
      if (this.currentStage >= 2) {
        this.dialog2 = true;
      }
      if (this.currentStage == 4) {
        this.valid = true;
      }
    },
    submitForm(){
      this.dialog = false
      const data = new URLSearchParams({
        reason: this.value,
      })
      this.axios.post(this.$backendApi + "customer/create-profile",data).then((response)=>{
        console.log(response.data)
      })
    }
  }
};
</script>