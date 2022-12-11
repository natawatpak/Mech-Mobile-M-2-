<template>
  <div class="pa-5">
    <div class="text-left text-h4 my-2">Incoming order</div>

    <v-card-text class="pa-0 ma-0">
      <v-window v-model="tab">
        <v-window-item value="one">
          <v-content class="justify-start pa-0">
            <v-container v-for="item in incoming" :key="item.id" class="pa-0" >
              <v-card @click="dialog=true" width="100%" variant="tonal" v-if="show" class="text-left pa-0 my-4">
                <v-card-title class="text-h6">
                  {{item.cus.fName}} | {{item.car.plate}}
                  <v-card-subtitle class="pa-0">{{item.date}}</v-card-subtitle>
                </v-card-title>
                <v-card-text class="text-h7 pa-0 px-4">
                  Car {{ item.car.type }} | {{ item.car.brand }} <br />
                  Problems {{ item.problem }} <br />
                  {{ item.location.distance }} km
                </v-card-text>
                <v-row justify="end" class="pa-6">
                </v-row>
              </v-card>
            </v-container>
          </v-content>
        </v-window-item>
      </v-window>
    </v-card-text>

    <v-dialog v-model="dialog" scrollable class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5"> Order {{details.id}}</v-card-title>
        <v-divider></v-divider>
        <v-card-text  style="height: 100%;" class="ma-0 pa-0">
          <v-card class="text-left ma-2 pa-1" variant="outlined">
            <v-title class="text-h6">Car Details</v-title>
              <v-card-text class="text-h7">
                Type: {{details.type}} | Brand: {{details.brand}}
                <br />
                License plate: {{details.plate}}
              </v-card-text>
          </v-card>
          <v-card
            variant="outlined"
            class="text-left ma-2 pa-1 d-flex justify-left align-center"
            height="100%"
          >
            <section>
              <v-card-title>Current Location</v-card-title>
              <v-card-text>{{location.lat+', '+ location.lng}}</v-card-text>
            </section>
          </v-card>
          <v-card class="text-left ma-2 pa-1" variant="outlined">
            <v-title class="text-h6">Problems</v-title>
            <v-card-text class="text-h7">
              <p v-for="p in problems" :key="p">{{'-' + p}}</p>
            </v-card-text>
          </v-card>
          <v-card class="text-left ma-2 pa-1" variant="outlined">
            <v-title class="text-h6">Description</v-title>
              <v-card-text class="text-h7">
                <p>{{description}}</p>
              </v-card-text>
          </v-card>
        </v-card-text>
        <v-divider></v-divider>
        <v-row class="justify-end">
          <v-card-action class="py-4 justify-end">
            <v-btn @click="dialog=false" class="mx-1" variant="tonal" color="error">
              Decline
            </v-btn>
            <v-btn to="/details" class="mx-1" color="blue" variant="tonal">
              Accept
            </v-btn>
          </v-card-action>
        </v-row>
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
      items: [
        {
          id: "11111",
          username: "nunnapat",
          date: "12 Nov 2022",
          distance: "2.2",
          type: "SUV",
          brand: "MG",

        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
        },
      ],
      details: {
        id: "12345",
        type: "SUV",
        brand: "MG",
        plate: "ศง 4727",
        name: "nunnapat",
        problem: "2",
        location: "2.1",
        status: "Finish",
      },
      tab: null,
      dialog: false,
      dialog2: false,
      location: {
        lat: "20",
        lng: "90"
      },
      chips: [ ],
      show: true,
      problems: ["no battery", "broken motor"],
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
                    "shopID": "1", "status": "Finish:Garage", "ticketID": "677fe4c6-447f-4d21-a0cd-c5dbe52f7fc8" }]
    };
  },
};
</script>