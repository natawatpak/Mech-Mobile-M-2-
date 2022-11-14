<template>
    <div class="home">
      <CallCar/>
    </div>
  </template>
  
  <script>
  // @ is an alias to /src
  import CallCar from '@/components/CallCar.vue'
  
  export default {
    name: 'CallCarView',
    components: {
      CallCar
    }
  }
  </script>
  