import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import router from './router'
import VueGoogleMaps from '@fawmi/vue-google-maps'

loadFonts()

createApp(App).use(router)
  .use(vuetify)
  .use(VueGoogleMaps, {
    load: {
        key: 'AIzaSyDa_vsioctDptTuaKkVjqqian4w-hkzpHY',
    },
})
  .mount('#app')
