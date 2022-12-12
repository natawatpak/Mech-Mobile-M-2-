<template>
  <div>
    {{ ticket }}

    <v-dialog v-model="dialog" style="width: 400px">
      <v-card class="justify-center">
        <v-card-title class="pa-4">Connection error</v-card-title>
        <div class="d-flex justify-center align-row ma-4">
          <v-img
            :aspect-ratio="16 / 9"
            :max-width="200"
            src="@/assets/connection-error.png"
            class="img-fluid"
          ></v-img>
        </div>
        <v-card-text class="pb-2"
          >Please click refresh to reload the page.</v-card-text
        >
        <v-btn @click="reloadPage()" variant="tonal" color="blue"
          >Refresh</v-btn
        >
      </v-card>
    </v-dialog>

    <v-divider></v-divider>
    <v-btn @click="updateTicketStatus('On the way')">On the way</v-btn>
    <v-btn @click="updateTicketStatus('Processing')">Processing</v-btn>
    <v-btn @click="updateTicketStatus('Finish:Garage')">Finish:Garage</v-btn>
    <v-btn @click="updateTicketStatus('Finish:Completed')"
      >Finish:Completed</v-btn
    >
  </div>
</template>

<script>
export default {
  data() {
    return {
      ticket: null,
      dialog: false,
    };
  },
  mounted() {
    this.getTicket(sessionStorage.getItem("ticketID"));
    this.socket = new WebSocket(
      this.$wsApi + "shop/ws/" + sessionStorage.getItem("ticketID")
    );
    console.log("Attempting Connection...");
    this.socket.onopen = () => {
      console.log("Successfully Connected");
      this.socket.send("Accepted");
    };
    this.socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
    });
    this.socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
      this.socket.send("Client Closed!");
      this.dialog = true
    };
    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
      this.dialog = true
    };
  },
  methods: {
    getTicket(id) {
      const data = new URLSearchParams({
        ticketID: id,
      });
      this.axios
        .post(this.$backendApi + "shop/get-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.ticket = response.data;
          Object.keys(response.data).forEach((prop) => console.log(prop));
        });
    },
    updateTicketStatus(status) {
      this.socket.send(status);
      const data = new URLSearchParams({
        ticketID: this.ticket.ticketID,
        cusID: this.ticket.cus.cusID,
        carID: this.ticket.car.carID,
        problem: this.ticket.problem,
        shopID: sessionStorage.getItem("shopID"),
        status: status,
        lat: this.ticket.location.lat,
        lng: this.ticket.location.lng,
        description: this.ticket.description,
      });
      this.axios
        .post(this.$backendApi + "shop/update-ticket", data)
        .then((response) => {
          console.log(response);
          this.ticket.status = response.data.status;
        });
    },
    reloadPage() {
      window.location.reload();
      this.dialog = false;
    },
  },
};
</script>
