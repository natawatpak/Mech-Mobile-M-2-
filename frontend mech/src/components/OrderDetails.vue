<template>
  <div class="text-center mx-16 justify-center">
    <v-card class="ma-6" flat>
        <div>
          <v-toolbar>
            <v-btn
              icon
              class="hidden-xs-only"
            >
              <v-icon>mdi-arrow-left</v-icon>
            </v-btn>

            <v-toolbar-title class="text-left text-h5 px-4">Order</v-toolbar-title>
          </v-toolbar>
        </div>
  </v-card>

  </div>
</template>

<style>

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