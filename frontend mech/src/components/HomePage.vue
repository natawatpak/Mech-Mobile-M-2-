<template>
  <div class="text-center ma-4">
    <v-list>
      <v-list-item>
        <v-list-title class="text-h5">
          Incoming Order
        </v-list-title>
      </v-list-item>
    </v-list>  
        <v-card
      class="text-left pa-4 d-flex justify-left align-center"
      outlined
      >
    <v-card-text>
      <div class="text-h6 text--primary">
        #123456
      </div>
      <p>adjective</p>
      <div class="text--primary">
        relating to or dependent on charity; charitable.<br>
        "an eleemosynary educational institution."
      </div>
    </v-card-text>
    <v-card-actions>
      <v-btn
        text
        color="teal accent-4"
        @click="reveal = true"
      >
        Learn More
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <v-card
        v-if="reveal"
        class="transition-fast-in-fast-out v-card--reveal"
        style="height: 100%;"
      >
        <v-card-text class="pb-0">
          <p class="text-h4 text--primary">
            Origin
          </p>
          <p>late 16th century (as a noun denoting a place where alms were distributed): from medieval Latin eleemosynarius, from late Latin eleemosyna ‘alms’, from Greek eleēmosunē ‘compassion’ </p>
        </v-card-text>
        <v-card-actions class="pt-0">
          <v-btn
            text
            color="teal accent-4"
            @click="reveal = false"
          >
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-expand-transition>
  </v-card>
  </div>
</template>

<style>
.text-h5 {
  display: flex;
  justify-content: left;
}
v-card--reveal {
  bottom: 0;
  opacity: 1 !important;
  position: absolute;
  width: 100%;
}
</style>

<script>
export default {
  data() {
    return {
      items: [
          { order: '0123456', name: 'Nunnapat Kriengchaiyaprug' },
        ],
    };
  }
};
</script>