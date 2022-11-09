
<template>
  <v-timeline direction="horizontal" line-inset="12">
    <v-timeline-item :dot-color="currentState == 'accepted' ? 'green':'white'">
      <template v-slot:opposite>Accepted</template>
    </v-timeline-item>

    <v-timeline-item :dot-color="currentState == 'processing' ? 'green':'white'">
      <template v-slot:opposite></template>
      Processing
    </v-timeline-item>

    <v-timeline-item :dot-color="currentState == 'completed'? 'green':'white'">
      <template v-slot:opposite>Completed</template>
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