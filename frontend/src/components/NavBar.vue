<template>
  <nav>
    <v-app-bar elevation="8" flat app>
      <v-app-bar-nav-icon class="grey--text" @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title class="text-uppercase grey--text">
        <span class="font-weight-bold">M^2</span>
      </v-app-bar-title>
      <v-spacer></v-spacer>   
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" absolute bottom temporary class="pa-4" style="position:fixed;">
        <v-list-item link>
          <v-list-item-content>
            <router-link to="/" class="text-decoration-none">
              <v-list-item-title class="text-h6">Home</v-list-item-title>
            </router-link>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item link>
          <v-list-item-content>
            <router-link to="/about" class="text-decoration-none">
              <v-list-item-title class="text-h6">About</v-list-item-title>
            </router-link>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
    </v-navigation-drawer>
  </nav>
  
</template>

<style>
nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #18139c;
}

.text-h6 {
  padding: 20px;
  color: #000000;
}

</style>


<script>
export default {
    data(){
        return{
            drawer: false,
            setting: false,
        }
    }
}
</script>