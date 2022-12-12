<template>
  <nav>
    <v-app-bar elevation = "4" flat app dark>
      <v-app-bar-nav-icon class="grey--text" @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title class="text-uppercase grey--text">
        <span class="font-weight-bold">M^2</span>
      </v-app-bar-title>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" absolute bottom temporary class="pa-4" style="position:fixed;">
      <v-list-item link>
        <v-list-item-content>
          <router-link to="/" class="text-decoration-none">
            <v-list-item-title class="text-h6">Dashboard</v-list-item-title>
          </router-link>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list-item link>
        <v-list-item-content>
          <router-link to="/order" class="text-decoration-none">
            <v-list-item-title class="text-h6">Incoming order</v-list-item-title>
          </router-link>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list-item link>
        <v-list-item-content>
          <router-link to="/history" class="text-decoration-none">
            <v-list-item-title class="text-h6">History</v-list-item-title>
          </router-link>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
    </v-navigation-drawer>
  </nav>
  
  
</template>

<style>
.text-h6 {
  padding: 20px;
  color: #000000;
}

</style>


<script>
 export default {
    data () {
      return {
        drawer: false,
        items: [
          { title: 'Dashboard', icon: 'mdi-view-dashboard' },
          { title: 'Order', icon: 'mdi-order' },
        ],
        right: null,
      }
    },
  }
</script>