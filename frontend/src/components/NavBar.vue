<template>
  <nav>
    <v-app-bar elevation = "4" flat app>
        <v-app-bar-nav-icon class="grey--text" @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title class="text-uppercase grey--text">
        <span class="font-weight-bold" color="blue">M^2</span>
      </v-app-bar-title>
      <v-spacer></v-spacer>
      
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" absolute bottom temporary>
      <v-list nav dense>
        <v-list-item><router-link to="/"><v-btn flat> Home</v-btn></router-link></v-list-item>
        <v-list-item><router-link to="/about"><v-btn flat> About</v-btn></router-link></v-list-item>
        <v-list-item><v-btn flat @click="setting=true"> Setting</v-btn></v-list-item>
      </v-list>
    </v-navigation-drawer>
      <v-dialog v-model="setting" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
          <v-toolbar dark
            color="blue-lighten-3"
          >
            <v-btn
              icon
              dark
              @click="setting=false"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>Settings</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
              <v-btn
                dark
                text
                @click="setting = false"
              >
                Save
              </v-btn>
            </v-toolbar-items>
          </v-toolbar>
          
        </v-card>
      </v-dialog>
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