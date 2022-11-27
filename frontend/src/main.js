import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import router from './router'
// import { createStore } from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueGoogleMaps from '@fawmi/vue-google-maps'

loadFonts()

createApp(App).use(router)
  .use(vuetify)
  .use(VueAxios, axios)
  .use(VueGoogleMaps, {
    load: {
        key: 'AIzaSyBzVgz4sJ1dYzVOOybCqLRV1p7sX34HuuQ',
    },})
  .mount('#app')
