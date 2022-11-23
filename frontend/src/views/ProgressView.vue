<template>
  <div class="pa-5">
    <ProgressBar currentState="4" />
    <v-spacer class="my-5"></v-spacer>
    <ProgressDetail :shop=shop :car=car :location=location :problems=problems />
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
    sessionStorage.put("ticketID", "1")

    this.ticketID = sessionStorage.getItem("ticketID")
    this.cusID = sessionStorage.getItem("cusID")

    // this.getActiveTicket()
  },
  methods: {
    getActiveTicket(){
      const data = new URLSearchParams({
        cusID: sessionStorage.getItem("cusID"),
      });
      console.log(sessionStorage.getItem("cusID"))
      this.axios
        .post("http://localhost:3000/customer/get-active-ticket", data)
        .then((response) => {
          console.log(response.data);
          this.carID = response.data.carID;
          this.shopID = response.data.shopID;
          this.status = response.data.status;
        });
    },
    getShopProfile(){
      const data = new URLSearchParams({
        shopID: this.shopID,
      });
      console.log(sessionStorage.getItem("cusID"))
      this.axios
        .post("http://localhost:3000/customer/get-shop-profile", data)
        .then((response) => {
          this.shop = response.data;
        });
    }
  }
};
</script>
