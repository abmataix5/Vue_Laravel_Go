import {store} from '@/store'


  export default {

    isAuthenticated(from,to,next){

      console.log('entra');
      if (store.getters['worker/isAuthenticated']) {
        next()
      } else {
        next('/login')
      }

    },

    isAdmin(from,to,next) {

      console.log('entra');
      if (store.getters['worker/isAdmin']) {

        const admin = store.getters['worker/isAdmin'];
        console.log(admin);
        next()
      } else {
        next('/login')
      }

    }

  }