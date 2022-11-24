<template>
  <div class="pa-5">
    <ProgressBar :currentState=status />
    <v-spacer class="my-5"></v-spacer>
    <ProgressDetail :shop=shop :car=car :problems=problems />
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
      shop: {
        shopID: undefined,
        shopName: "",
        shopAddress: "",
        shopTel: "",
        shopEmail: "",
        ratings: "",
        location: { lat:"20", lng:"90"},
      },
      car: { id: "1", type: "SUV", brand: "MG", plate: "ก2113" },
      
      problems: []
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
          this.getCarDetail()
          if(this.status!="Active"){
            this.getShopProfile()
          }
        })
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
    getCarDetail(){
      const data = new URLSearchParams({
        carID: this.carID,
      });
      this.axios
        .post("http://localhost:3000/customer/get-car", data)
        .then((response) => {
          this.car = response.data;
        });
    },
  }
};
</script>
