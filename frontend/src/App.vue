<template>
  <v-app :style="{background: '#343A4A'}">
    <Navbar/>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>



<script>
import Navbar from "@/components/NavBar.vue";

export default({
  components:{Navbar},
  data (){
    return{
      
    }
  }
});
</script>
