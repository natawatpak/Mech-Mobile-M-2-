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

let app = createApp(App).use(router)
app.config.globalProperties.$backendApi = 'https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/'
app.config.globalProperties.$wsApi = 'wss://6eblsxltxc.execute-api.us-east-1.amazonaws.com/production'
app.use(vuetify)
  .use(VueAxios, axios)
  .use(VueGoogleMaps, {
    load: {
        key: 'AIzaSyBzVgz4sJ1dYzVOOybCqLRV1p7sX34HuuQ',
    },})
  .mount('#app')

