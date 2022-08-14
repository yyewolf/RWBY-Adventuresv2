import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import router from './router'
import Notifications from '@kyvg/vue3-notification'

loadFonts()

createApp(App)
  .use(router)
  .use(vuetify)
  .use(Notifications)
  .mount('#app')
