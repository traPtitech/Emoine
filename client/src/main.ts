import { createApp } from 'vue'
import router from './router'
import store from './store'
import App from './App.vue'
import { setup } from '/@/lib/connect'

import 'ress'
import './index.scss'

const app = createApp(App)
app.use(router)
app.use(store)
app.mount('#app')

setup()
