<template>
  <div class="text-left ma-4">

    <v-card variant="tonal" color="teal" class="pa-4">
      <v-row class="justify-space-between align-center">
      <v-card-title>Incoming order</v-card-title>
      <router-link to="/order" class="text-decoration-none">
        <v-btn color="teal-darken-1" variant="outlined">
          Go check
        </v-btn>
      </router-link>
    </v-row>
    </v-card>

    <div class="mt-4">
      <v-text class="text-h5">Today summary</v-text>
    </div>

    <v-row no-gutters>
      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>Total order</v-card-title>
          <v-card-text class="text-center text-h6">{{total}}</v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>On the way</v-card-title>
          <v-card-text class="text-center text-h6">{{otw}}</v-card-text>
        </v-card>
      </v-col>

      <v-responsive width="100%"></v-responsive>

      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>In process</v-card-title>
          <v-card-text class="text-center text-h6">{{onprocess}}</v-card-text>
        </v-card>
      </v-col>

      <v-col>
        <v-card class="pa-2 ma-2" variant="tonal">
          <v-card-title>Finish</v-card-title>
          <v-card-text class="text-center text-h6">{{finish}}</v-card-text>
        </v-card>
      </v-col>
    </v-row>


        <v-tabs v-model="tab" grow class="ma-4">
          <v-tab class="text-h5" value="one">On process</v-tab>
          <v-tab class="text-h5" value="two">Finish</v-tab>
        </v-tabs>

        <v-card-text>
          <v-window v-model="tab">
            <v-window-item value="one">
              <v-content>
                <v-container v-for="item in incoming" :key="item.id" class="pa-0 my-4">
                  <v-card width="100%" variant="tonal" class="text-left pa-4 ">
                    <v-row class="pa-2">
                      <v-card-title>
                        {{item.cus.fName}} | {{item.car.plate}}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{
                        item.status
                      }}</v-chip>
                    </v-row>
                    <v-card-text class="text-h7">
                      Car {{ item.car.type }} | {{ item.car.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.location.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>

            <v-window-item value="two">
              <v-content>
                <v-container v-for="item in incoming" :key="item.id" class="pa-0 my-4">
                  <v-card width="100%" variant="tonal" class="text-left pa-4 ">
                    <v-row class="pa-2">
                      <v-card-title>
                        {{item.cus.fName}} | {{item.car.plate}}
                      </v-card-title>
                      <v-chip class="ma-2" color="yellow-darken-3">{{
                        item.status
                      }}</v-chip>
                    </v-row>
                    <v-card-text class="text-h7">
                      Car {{ item.car.type }} | {{ item.car.brand }} <br />
                      Problems {{ item.problem }} <br />
                      {{ item.location.distance }} km
                    </v-card-text>
                  </v-card>
                </v-container>
              </v-content>
            </v-window-item>
          </v-window>
        </v-card-text>
  </div>
</template>

<style>
.text-h4 {
  display: flex;
  justify-content: left;
  align-content: center;
}
</style>

<script>
export default {
  data() {
    return {
      dialog: false,
      orders: 1,
      total: 5,
      otw: 1,
      onprocess: 1,
      finish: 1,
      description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.",
      incoming: [ { "car": { "brand": "dgdg", "carID": "fa96015c-8b1c-4f3d-8481-dfc6b54b3476", "plate": "1234", "type": "vfx" }, 
                    "cus": { "cusID": "f67efc77-629d-4672-a753-558b1c0dd250", "fName": "ggg", "lName": "gg" }, 
                    "location": { "lat": 13.726849, "lng": 100.770309, "distance": 2.34 }, 
                    "problem": "1,4,3", 
                    "date": "12 Nov 2022",
                    "shopID": "1", "status": "Finish:Garage", "ticketID": "677fe4c6-447f-4d21-a0cd-c5dbe52f7fc8"},
                    { "car": { "brand": "honda", "carID": "fa96015c-8b1c-4f3d-8481-dfc6b54b3476", "plate": "SN 4727", "type": "sedan" }, 
                    "cus": { "cusID": "f67efc77-629d-4672-a753-558b1c0dd250", "fName": "Nunnapat", "lName": "Kriengchaiyaprug" }, 
                    "location": { "lat": 13.726849, "lng": 100.770309, "distance": 1.23 }, 
                    "problem": "1,4,3", 
                    "date": "12 Nov 2022",
                    "shopID": "1", "status": "Finish:Garage", "ticketID": "677fe4c6-447f-4d21-a0cd-c5dbe52f7fc8" }],
      tab: null,
    };
  },
};
</script>