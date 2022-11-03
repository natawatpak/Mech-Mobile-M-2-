<template>
  <div class="text-center ma-4">
    <v-list>
      <v-list-item>
        <router-link to="/callmech" class="text-decoration-none">
          <v-btn variant="tonal" color="green" size="large">Call Mechanic!!</v-btn>
        </router-link>
      </v-list-item>
      <v-list-item>
        <v-btn variant="tonal" color="blue" size="large">Call Slide car!!</v-btn>
      </v-list-item>
    </v-list>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  }
};
</script>