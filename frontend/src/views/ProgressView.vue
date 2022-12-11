<template>
  <div class="pa-5">
    <ProgressBar id='state' :currentState=status />
    <v-spacer class="my-5"></v-spacer>
    <ProgressDetail :shop=shop :car=car :location=location :problems=problems :currentState=status />

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
      socket: undefined,
      dialog: true,
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
  created () {
    this.ticketID = sessionStorage.getItem("ticketID")
    console.log(this.ticketID)
    this.cusID = sessionStorage.getItem("cusID")

    this.getTicket()

    this.socket = new WebSocket("wss://7c2mohris1.execute-api.us-east-1.amazonaws.com/production?ticketID=" + this.ticketID);

    console.log("Attempting Connection...");

    this.socket.onopen = () => {
        console.log("Successfully Connected");
        this.socket.send("Hi From the Client!")
        this.dialog = false;
    };

    this.socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
      if(event.data == "Accepted"){
        this.getTicket();
      }else{
        this.status = event.data
      }
      
    });
    
    this.socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        this.socket.send("Client Closed!")
    };

    this.socket.onerror = error => {
        this.console.log("Socket Error: ", error);
    };

  },
  methods: {
    getTicket(){
      const data = new URLSearchParams({
        ticketID: this.ticketID
      });
      this.axios
        .post(this.$backendApi + "customer/get-ticket", data,{
        headers:{
            Authorization: sessionStorage.getItem("jwt")
        }})
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
        .post(this.$backendApi + "customer/get-shop-profile", data)
        .then((response) => {
          this.shop = response.data;
        });
    },
    getCarDetail(){
      const data = new URLSearchParams({
        carID: this.carID,
      });
      this.axios
        .post(this.$backendApi + "customer/get-car", data,
        {
          headers:{
              Authorization: sessionStorage.getItem("jwt")
          }
        })
        .then((response) => {
          this.car = response.data;
        });
    },
    
  }
};
</script>
