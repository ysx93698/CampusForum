import { createApp } from 'vue'
import './index.css'
import App from './App.vue'
import router from "@/router";
import {store} from "@/pinia";

const app = createApp(App)

app
    .use(router)
    .use(store)

app.mount('#app')