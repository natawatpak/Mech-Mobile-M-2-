<template>
  <div class="pa-4 px-lg-16 px-xl-16 mx-lg-auto mx-xl-auto">
    <div class="text-left text-h4 my-2">Incoming order</div>
    
    <v-card-text class="pa-0 ma-0">
      <v-content class="justify-start pa-0">
        <v-container v-for="t in tickets" :key="t.ticketID" fluid class="pa-0" >
          <v-card @click="getDetails(t)" width="100%" v-if="show" class="text-left pa-4 rounded-lg elevation-4 my-4">
            <v-card-title class="text-h6 pa-0">
              {{t.cus.fName}} | {{t.car.plate}}
              <v-card-subtitle class="pa-0">{{t.date}}</v-card-subtitle>
            </v-card-title>
            <v-card-text class="text-h7 pa-0 py-2">
              Car {{ t.car.type }} | {{ t.car.brand }} <br />
              Problems {{ t.problem.length }} <br />
              {{ t.location.distance }} km
            </v-card-text>
          </v-card>
        </v-container>
      </v-content>
    </v-card-text>

    <v-dialog v-model="dialog" scrollable class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5 text-center pa-0"> Order details</v-card-title>
        <v-card-text  style="height: 100%;" class="ma-0 pa-2">
          <div class="py-2">
            <v-text class="text-h6 pa-0">Customer name</v-text>
            <br>
            {{(details.cus.fName+' '+details.cus.lName)}}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Car details</v-text>
            <br>
            Type: {{details.car.type}}
            <br>
            Brand: {{details.car.brand}}
            <br>
            License plate: {{details.car.plate}}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Location</v-text>
            <br>
            {{details.location.lat+', '+ details.location.lng}}
            <br>
            Distance: {{details.location.distance}} km
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Problems</v-text>
            <br>
            <p v-for="p in details.problem" :key="p">{{'- ' + p}}</p>
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Description</v-text>
            <br>
            {{details.description}}
          </div>
        </v-card-text>
        <v-divider></v-divider>
        <v-row class="justify-end pt-2">
          <v-card-action class="py-4 justify-end">
            <v-btn @click="dialog=false" class="mx-1" variant="tonal" color="error">
              Decline
            </v-btn>
            <v-btn @click="acceptTicket()" class="mx-1" color="blue" variant="tonal">
              Accept
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>
  </div>
</template>

<style>
.text-subtitle-1 {
  color: #8594A6;
}
</style>

<script>
export default {
  data() {
    return {
      dialog: false,
      show: true,
      details: [],
      tickets: [ { "car": { "brand": "dgdg", "carID": "fa96015c-8b1c-4f3d-8481-dfc6b54b3476", "plate": "1234", "type": "vfx" }, 
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
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam rutrum tincidunt arcu sed gravida. Suspendisse mollis ex sed magna viverra, eu tincidunt velit lobortis. Phasellus accumsan mauris est, in pretium orci lacinia in. Nulla in bibendum ex, eu condimentum mauris. Pellentesque quis nisl a justo ultrices molestie ac mattis libero. Aliquam vel sollicitudin quam, vel molestie nunc. Proin dolor dolor, vehicula sed nisl dapibus, semper scelerisque nisl. Sed id neque ac metus efficitur vestibulum. Phasellus quis nibh ac dui molestie fringilla sit amet bibendum sem. Quisque consectetur sit amet nunc ac elementum. Suspendisse vulputate mauris erat, nec tempus dui vulputate.", 
                    "date": '12 Nov 2022', 
                    "shopID": "1", "status": "Finish:Garage", "ticketID": "677fe4c6-447f-4d21-a0cd-c5dbe52f7fc8" }]
    };
  },
  mounted() {
    this.getActiveTicket();
    this.socket = new WebSocket("wss://ar20w151pd.execute-api.us-east-1.amazonaws.com/production?shopID="+sessionStorage.getItem("shopID"));
    console.log("Attempting Connection...");

    this.socket.onopen = () => {
      console.log("Successfully Connected");
      this.socket.send("Accepted");
    };

    this.socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
      this.getActiveTicket();
    });

    this.socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
      this.socket.send("Client Closed!");
    };

    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
    };
  },
  methods: {
    acceptTicket() {
      let ticket = this.details;
      sessionStorage.setItem("ticketID", ticket.ticketID);

      const data = new URLSearchParams({
        ticketID: ticket.ticketID,
        cusID: ticket.cus.cusID,
        carID: ticket.car.carID,
        problem: ticket.problem,
        shopID: sessionStorage.getItem("shopID"),
        lat: ticket.location.lat,
        lng: ticket.location.lng,
        description: ticket.description,
      });
      this.axios
        .post(this.$backendApi + "shop/accept-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.$router.push("/details")
        });
    },
    getDetails(t) {
      this.details = t;
      console.log(this.details)
      this.dialog = true;
    },
    getActiveTicket() {
      this.axios
        .post(this.$backendApi + "shop/get-ticket-list")
        .then((response) => {
          console.log(response.data);
          this.tickets = response.data;
        });
    },
  },
};
</script>