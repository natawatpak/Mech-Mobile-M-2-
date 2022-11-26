<template>
  <div class="pa-5">
    <ProgressBar id='state' currentState="1" />
    <v-spacer class="my-5"></v-spacer>
    <ProgressDetail :shop=shop :car=car :location=location :problems=problems :currentStage="1"/>

    <v-dialog 
      v-model="dialog"
      style="width: 300px;">
      <v-card class="py-6">
        <v-card-text>
          Loading, please wait ...
          <v-progress-linear
            indeterminate
            color="deep-purple-accent-4"
            rounded
            height="6"
            class="my-4">
          </v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
// @ is an alias to /src
import ProgressBar from "@/components/ProgressBar.vue";
import ProgressDetail from "@/components/ProgressDetail.vue";

export default {
  data() {
    return {
      cusID: undefined,
      ticketID: undefined,
      carID: undefined,
      shopID: undefined,
      status: undefined,
      createdTime: undefined,
      acceptedTime: undefined,
      dialog: true,
      shop: {
        shopID: undefined,
        shopName: "A",
        shopAddress: "567 Kingston Rd.",
        shopTel: undefined,
        shopEmail: undefined,
        ratings: "4.5"
      },
      car: { id: "1", type: "SUV", brand: "MG", plate: "ก2113" },
      location: { lat:"20", lng:"90"},
      problems: ["no battery", "broken motor"]
    };
  },
  components: {
    ProgressBar,
    ProgressDetail
  },
  mounted () {
    this.ticketID = sessionStorage.getItem("ticketID")
    console.log(this.ticketID)
    this.cusID = sessionStorage.getItem("cusID")

    this.getTicket()
  },
  methods: {
    getTicket(){
      const data = new URLSearchParams({
        ticketID: this.ticketID,
      });
      this.axios
        .post("http://localhost:3000/customer/get-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.carID = response.data.carID;
          this.shopID = response.data.shopID;
          this.status = response.data.status;
          this.problems = response.data.problem;
          this.createdTime = response.data.createdTime;
          this.acceptedTime = response.data.acceptedTime;
        });
    },
    getShopProfile(){
      const data = new URLSearchParams({
        shopID: this.shopID,
      });
      this.axios
        .post("http://localhost:3000/customer/get-shop-profile", data)
        .then((response) => {
          this.shop = response.data;
        });
    },
  }
};
</script>
