<template>
  <div class="home">
    <OrderPage />
  </div>
</template>

<script>
// @ is an alias to /src
import OrderPage from "@/components/OrderPage.vue";

export default {
  name: "OrderView",
  components: {
    OrderPage,
  },
  data() {
    return {
      socket: undefined,
    };
  },
  mounted() {
    this.socket = new WebSocket("ws://127.0.0.1:3000/shop/ws");
    console.log("Attempting Connection...");

    this.socket.onopen = () => {
      console.log("Successfully Connected");
      this.socket.send("Hi From the Client!");
    };

    this.socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
    });

    this.socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
      this.socket.send("Client Closed!");
    };

    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
    };
  },
};
</script>
