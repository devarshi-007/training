import { createApp } from 'vue'
import App from './App.vue'
import NewComp from './components/NewComp.vue'
const app = createApp(App);
app.mount('#app')
app.component('NewComp',NewComp);