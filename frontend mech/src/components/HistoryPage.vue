<template>
  <div class="text-center ma-4 justify-center">
    <div class="text-left text-h4 my-2">History</div>

    <v-card-text class="pa-0 ma-0">
          <v-content class="justify-start pa-0">
            <v-container v-for="item in incoming" :key="item.id" class="pa-0" >
              <v-card @click="dialog=true" width="100%" class="text-left pa-0 my-4 rounded-lg elevation-4">
                <v-row class="px-4 pt-4 d-flex justify-space-between align-center">
                  <v-card-title class="text-h6">
                    {{item.cus.fName}} | {{item.car.plate}}
                  </v-card-title>
                  <v-chip class="ma-2" color="green">{{item.status}}</v-chip>
                </v-row>
                <v-card-subtitle>{{item.date}}</v-card-subtitle>
                <v-card-text class="text-h7">
                  Car {{ item.car.type }} | {{ item.car.brand }} <br />
                  Total problems: {{ item.problem.length }} <br />
                  {{ item.location.distance }} km
                </v-card-text>
              </v-card>
            </v-container>
          </v-content>
    </v-card-text>

    <v-dialog v-model="dialog" scrollable class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-row class="pa-5">
          <v-btn
            icon
            dark
            @click="dialog = false"
          >
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-card-title class="text-h5 text-center"> Order details</v-card-title>
        </v-row>
        <v-card-text  style="height: 100%;" class="ma-0 pa-2">
          <div class="py-2">
            <v-text class="text-h6 pa-0">Customer name</v-text>
            <br>
            {{(incoming[0].cus.fName+' '+incoming[0].cus.lName)}}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Car details</v-text>
            <br>
            Type: {{incoming[0].car.type}}
            <br>
            Brand: {{incoming[0].car.brand}}
            <br>
            License plate: {{incoming[0].car.plate}}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Location</v-text>
            <br>
            {{incoming[0].location.lat+', '+ incoming[0].location.lng}}
            <br>
            Distance: {{incoming[0].location.distance}} km
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Problems</v-text>
            <br>
            <p v-for="p in incoming[0].problem" :key="p">{{'- ' + p}}</p>
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Description</v-text>
            <br>
            {{incoming[0].description}}
          </div>
        </v-card-text>
        <v-divider></v-divider>
      </v-card>
    </v-dialog>
  </div>
</template>

<style>

</style>

<script>
export default {
  data() {
    return {
      dialog: false,
      incoming: [ { "car": { "brand": "dgdg", "carID": "fa96015c-8b1c-4f3d-8481-dfc6b54b3476", "plate": "1234", "type": "vfx" }, 
                    "cus": { "cusID": "f67efc77-629d-4672-a753-558b1c0dd250", "fName": "ggg", "lName": "gg" }, 
                    "location": { "lat": 13.726849, "lng": 100.770309, "distance": 2.34 }, 
                    "problem": ["Lubricating oil overdue that should be changed", "Out of Brake lining"],
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.", 
                    "date": '12 Nov 2022',
                    "shopID": "1", "status": "Finish:Garage", "ticketID": "677fe4c6-447f-4d21-a0cd-c5dbe52f7fc8" },
                    { "car": { "brand": "honda", "carID": "fa96015c-8b1c-4f3d-8481-dfc6b54b3476", "plate": "SN 4727", "type": "sedan" }, 
                    "cus": { "cusID": "f67efc77-629d-4672-a753-558b1c0dd250", "fName": "Nunnapat", "lName": "Kriengchaiyaprug" }, 
                    "location": { "lat": 13.726849, "lng": 100.770309, "distance": 2.34 }, 
                    "problem": ["Lubricating oil overdue that should be changed", "Out of Brake lining"],
                    "date": '12 Nov 2022', 
                    "shopID": "1", "status": "Finish:Garage", "ticketID": "677fe4c6-447f-4d21-a0cd-c5dbe52f7fc8" }]
    };
  },
};
</script>