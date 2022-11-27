<template>
  <div>
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
  </div>
</template>

<script>
export default {
  data() {
    return {
      tickets: [],
    };
  },
  mounted() {
    this.getActiveTicket();
    sessionStorage.setItem("shopID", 1);
  },
  methods: {
    getActiveTicket() {
      this.axios
        .post("http://localhost:3000/shop/get-ticket-list")
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
        .post("http://localhost:3000/shop/accept-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.$router.push("/example2")
        });
    },
  },
};
</script>
