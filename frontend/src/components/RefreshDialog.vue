<template>
    <v-dialog v-model="dialog" style="width: 400px;">
    <v-card class="justify-center">
      <v-card-title class="pa-4">Connection error</v-card-title>
      <div class="d-flex justify-center align-row ma-4">
        <v-img :aspect-ratio="16/9" :max-width="200" src="@/assets/connection-error.png" class="img-fluid"></v-img>
      </div>
      <v-card-text class="pb-2">Please click refresh to reload the page.</v-card-text>
      <v-btn @click="reloadPage()" variant="tonal" color="blue">Refresh</v-btn>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data() {
    return {
      dialog: true
    }
  },
  methods: {
        reloadPage(){
          window.location.reload()
          this.dialog = false;
        }
      },
};  
</script>
