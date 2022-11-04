
<template>
  <v-timeline direction="horizontal" line-inset="12">
    <v-timeline-item :dot-color="currentState == 1? 'green':'white'">
      <template v-slot:opposite>Opposite</template>
      Content
    </v-timeline-item>

    <v-timeline-item :dot-color="currentState == 2? 'green':'white'">
      <template v-slot:opposite>Opposite</template>
      Content
    </v-timeline-item>

    <v-timeline-item :dot-color="currentState == 3? 'green':'white'">
      <template v-slot:opposite>Opposite</template>
      Content
    </v-timeline-item>
  </v-timeline>
</template>
<script>
export default {
  props:{
    currentState:{
      type: String
    }
  },
  data() {
    return {
    };
  }
};
</script>