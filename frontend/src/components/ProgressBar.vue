<template>
  <v-timeline direction="horizontal" line-inset="12">
    <v-timeline-item size="large" v-model="items" :dot-color="currentState == 1? 'green':'white'" :icon=items[0]>
      <template v-slot:opposite>Accepted</template>
    </v-timeline-item>

    <v-timeline-item size="large" v-model="items" :dot-color="currentState == 2? 'green':'white'" :icon=items[1]>
      On the way
    </v-timeline-item>

    <v-timeline-item size="large" v-model="items" :dot-color="currentState == 3? 'green':'white'" :icon=items[2]>
      <template v-slot:opposite>On process</template>
    </v-timeline-item>
    <v-timeline-item size="large" v-model="items" :dot-color="currentState == 4? 'green':'white'" :icon=items[3]>
      Complete
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
      items: ['mdi-thumb-up', 'mdi-car', 'mdi-wrench', 'mdi-check'],
    };
  }
};
</script>