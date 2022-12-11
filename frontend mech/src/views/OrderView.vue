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
};
</script>
