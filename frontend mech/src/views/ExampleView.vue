<template>
  <div class="pa-4 px-lg-16 px-xl-16 mx-lg-auto mx-xl-auto">

    <v-dialog v-model="dialog" style="width: 400px;">
    <v-card class="justify-center">
      <v-card-title class="pa-4">Connection error</v-card-title>
      <div class="d-flex justify-center align-row ma-4">
        <v-img :aspect-ratio="16/9" :max-width="200" src="@/assets/connection-error.png" class="img-fluid"></v-img>
      </div>
      <v-card-text class="pb-2">Please click refresh to reload the page.</v-card-text>
      <v-btn @click="reloadPage()" variant="tonal" color="blue">Refresh</v-btn>
    </v-card>
  </v-dialog>

   
    <div class="text-left text-h4 my-2">Incoming order</div> 

    <!-- <div v-if="tickets.length > 0"> -->
    <v-card-text class="pa-0 ma-0">
      <v-content class="justify-start pa-0">
        <v-container v-for="t in tickets" :key="t.ticketID" class="pa-0">
          <v-card
            @click="getDetails(t)"
            width="100%"
            class="text-left pa-4 rounded-lg elevation-4"
          >
            <v-card-title class="text-h6 pa-0">
              {{ t.cus.fName }} | {{ t.car.plate }}
            </v-card-title>
            <v-card-text class="text-subtitle-2 pa-0 py-2">
              Car type: {{ t.car.type }}
              <br />
              Car brand: {{ t.car.brand }}
              <br />
              Total problems: {{ splitProblem(t.problem).length }}
            </v-card-text>
          </v-card>
        </v-container>
      </v-content>
    </v-card-text>

    <v-dialog v-model="dialog" scrollable class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5 text-center pa-0"
          >Order details</v-card-title
        >
        <v-card-text style="height: 100%" class="ma-0 pa-0">
          <div class="py-2">
            <v-text class="text-h6 pa-0">Customer name</v-text>
            <br />
            {{ details.cus.fName + " " + details.cus.lName }}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Car details</v-text>
            <br />
            Type: {{ details.car.type }}
            <br />
            Brand: {{ details.car.brand }}
            <br />
            License plate: {{ details.car.plate }}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Location</v-text>
            <br />
            Latitude: {{ details.location.lat }}
            <br />
            Longtitude: {{ details.location.lng }}
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Problems</v-text>
            <br /> 
            <li v-for="p in splitProblem(details.problem)" :key="p">{{ p }}</li>
          </div>
          <v-divider></v-divider>
          <div class="py-2">
            <v-text class="text-h6 pa-0">Description</v-text>
            <br />
            {{ details.description }}
          </div>
        </v-card-text>
        <v-divider></v-divider>
        <v-row class="justify-end pt-2">
          <v-card-action class="py-4 justify-end">
            <v-btn
              @click="dialog = false"
              class="mx-1"
              variant="tonal"
              color="error"
            >
              Decline
            </v-btn>
            <v-btn
              flat
              @click="acceptTicket(details)"
              class="mx-1"
              color="blue"
              variant="tonal"
            >
              Accept
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>
    <!-- </div> -->
  </div>
</template>

<script>

export default {
  data() {
    return {
      tickets: [],
      shopID: "",
      problems: [],
      dialog: false,
      details: [],
    };
  },
  mounted() {
    this.getActiveTicket();
    console.log(sessionStorage.getItem("shopID"));
    this.shopID = sessionStorage.getItem("shopID");
    this.socket = new WebSocket(this.$wsApi + "shop/ws/active-ticket");
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
      this.dialog = true;
    };

    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
      this.dialog = true;
    };
  },
  methods: {
    splitProblem(p) {
      return (this.problems = p.split(","));
    },
    getDetails(t) {
      this.details = t;
      console.log(this.details);
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

    acceptTicket(ticket) {
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
          this.$router.push("/example2");
        });
    },
    reloadPage(){
          window.location.reload()
          this.dialog = false
    },
  },
};
</script>
