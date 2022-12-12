<template>
  <div class="text-center ma-4">
    <v-list class="herosection justify-center">
      <div class="d-flex justify-center ma-4">
        <v-img :aspect-ratio="16/9" :max-width="200" src="@/assets/car-repair.png" class="img-fluid"></v-img>
      </div>
      <v-list-item class="justify-center">
        <div class="ma-4 text-center"> 
          <h2>Mech Mobile</h2>
          <h4>Are you facing some car problems?</h4>
          <h4>Let us help!</h4>
        </div>
      </v-list-item>

      <v-list-item>
        <router-link to="/callmech" class="text-decoration-none">
          <v-btn color="green-lighten-1" rounded="lg" size="large">Call Mechanic!!</v-btn>
        </router-link>
      </v-list-item>
      <v-list-item>
        <router-link to="/callcar" class="text-decoration-none">
          <v-btn color="blue-darken-2" size="large">Call Slide car!!</v-btn>
        </router-link>
      </v-list-item>
    </v-list>
  </div>
</template>

<style>
.herosection {
  background-color: transparent;
}
</style>
<script>
</script>