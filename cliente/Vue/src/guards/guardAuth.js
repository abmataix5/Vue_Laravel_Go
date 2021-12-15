import {store} from '@/store'


  export default {

    workerAuthenticated(from,to,next){

      if (store.getters['worker/isWorkerAuthenticated']) {
        next()
      } else {
        next('/login')
      }

    },

    isAdmin(from,to,next) {

      if (store.getters['worker/isAdmin']) {
        next()
      } else {
        next('/login')
      }

    }

  }