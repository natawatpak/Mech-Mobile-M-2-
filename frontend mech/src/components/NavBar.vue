<template>
  <nav>
    <v-app-bar elevation = "4" flat app>
      <v-app-bar-title class="text-uppercase grey--text">
        <span class="font-weight-bold" color="blue">M^2</span>
      </v-app-bar-title>
    </v-app-bar>
    <v-navigation-drawer permanent>
      <v-list
        dense
        nav
      >
        <v-list-item
          v-for="item in items"
          :key="item.title"
          link
        >
          <v-list-item-content>
            <v-list-item-title class="text-h6">{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
      </v-list>
    </v-navigation-drawer>
  </nav>
  
  
</template>

<style>
.text-h6{
  padding: 20px;
}
</style>


<script>
 export default {
    data () {
      return {
        items: [
          { title: 'Dashboard', icon: 'mdi-view-dashboard' },
          { title: 'Order', icon: 'mdi-order' },
        ],
        right: null,
      }
    },
  }
</script>