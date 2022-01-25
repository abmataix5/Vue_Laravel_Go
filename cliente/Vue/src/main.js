import { createApp } from 'vue'
import 'bootstrap/dist/css/bootstrap.css';
import './main.css'; 
import App from './App.vue'


import router from "./router";
import { store } from "./store"

import { library } from "@fortawesome/fontawesome-svg-core";
import { faFacebook,faTwitter,faSnapchat,faInstagram } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(faFacebook,faTwitter,faSnapchat,faInstagram);

createApp(App)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(router)
  .use(store)
  .mount("#app");

