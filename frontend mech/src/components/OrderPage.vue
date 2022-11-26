<template>
  <div class="text-center mx-16 justify-center">
    <div class="text-left text-h4 my-4 pa-4">Order</div>
      <v-container fluid v-for="item in items" :key="item.id" class="pa-0 mx-4">
      <v-card :to="'/details'" variant="tonal" class="text-left pa-4 my-4 mr-6">
        <v-row class="pa-2">
          <v-card-title class="text-h6">
            Order {{item.id}} | {{item.username}}
          </v-card-title>
          <v-chip class="ma-2" color="yellow-darken-3">{{item.status}}</v-chip>
        </v-row>      
        <v-card-subtitle>{{item.date}}</v-card-subtitle>
        <v-card-text class="text-h7">
          Car {{ item.type }} | {{ item.brand }} <br />
          Problems {{ item.problem }} <br />
          {{ item.distance }} km
        </v-card-text>
      </v-card>
    </v-container>
  </div>
</template>

<style>

</style>

<script>
export default {
  mounted(){
    this.getActiveTicket()
  },
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
          problem: "2",
          status: "On-precess",
        },
        {
          id: "22222",
          username: "natawat",
          date: "18 Nov 2022",
          distance: "0.4",
          type: "Sedan",
          brand: "MG",
          problem: "2",
          status: "On the way",
        },
        {
          id: "33333",
          username: "kiitiphum",
          date: "22 Nov 2022",
          distance: "1.3",
          type: "Van",
          brand: "MG",
          status: "Accept",
        },
      ],
    };
  },
  methods:{
    getActiveTicket(){
      this.axios
        .post("http://localhost:3000/shop/get-ticket-list")
        .then((response) => {
          console.log(response.data);
          this.items = response.data
        });
    },
  }
};
</script>