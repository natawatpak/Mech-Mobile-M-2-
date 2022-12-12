<template>
  <div>
    <v-dialog>dialog-transition</v-dialog>
    {{ shopID }}
    <v-card v-for="t in tickets" :key="t.ticketID">
      <v-card-title>
        {{ t.ticketID }}
      </v-card-title>
      <v-card-text>
        {{ t }}
      </v-card-text>
      <v-card-actions>
        <v-btn flat @click="acceptTicket(t)">accept</v-btn>
      </v-card-actions>
    </v-card>
    <RefreshDialog :dialog=dialog :key="dialog"/>
  </div>
</template>

<script>
import RefreshDialog from '@/components/RefreshDialog.vue';

export default {
  data() {
    return {
      tickets: [],
      shopID: "",
      dialog: true,
    };
  },
  components(){
    RefreshDialog
  },
  mounted() {
    this.getActiveTicket();
    console.log(sessionStorage.getItem("shopID"))
    this.shopID = sessionStorage.getItem("shopID")
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
      this.dialog = true
    };

    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
      this.dialog = true
    };
  },
  methods: {
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
          this.$router.push("/example2")
        });
    },
  },
};
</script>
