<template>
  <div class="pa-5">
    <ProgressBar currentState="2" />
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
      shop: {
        name: "A",
        address: "567 Kingston Rd.",
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
  }
};
</script>
