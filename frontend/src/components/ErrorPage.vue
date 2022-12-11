<template>
    <div class="pa-16 text-center align-center mb-6">
      <div class="d-flex justify-center ma-4">
        <v-img :aspect-ratio="16/9" :max-width="420" src="@/assets/404.png" class="img-fluid"></v-img>
      </div>
      <br />
      <h1>Page Not Found</h1>
      <v-spacer class="my-5"></v-spacer>
      <h3>We're sorry, the page you requested could not be found</h3>
      <h3>Please go back to home page</h3>
      <v-spacer class="my-5"></v-spacer>
      <section class="text-center">
        <router-link to="/home" class="text-decoration-none"><v-btn class="mx-1" variant="tonal" color="blue">Go home</v-btn></router-link>
      </section>
    </div>
  </template>