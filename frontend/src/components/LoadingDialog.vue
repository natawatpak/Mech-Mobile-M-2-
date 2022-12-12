<template>
  <v-dialog 
    v-model="dialog"
    style="width: 300px;"
  >
    <v-card class="py-6">
      <v-card-text>
        Loading, please wait ...
        <v-progress-linear
          indeterminate
          color="blue-darken-1"
          rounded
          height="6"
          class="my-4">
        </v-progress-linear>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data() {
    return {
      dialog: true
    }
  }
}
</script>

<style>

</style>