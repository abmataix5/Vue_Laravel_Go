import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/Home';
import Rentings from '../views/Rentings';
import Workers from '../views/Workers';
import Courts from '../views/AddCourt';
import AddWorker from '../views/AddWorker';
import Login from '../views/Login'
import guardAuth from "../guards/guardAuth";
/* import authGuard from '../guards/guardAuth' */
/* import store from '@/store'

const authGuard = (to, from, next) => {
    if (store.getters['worker/isAuthenticated']) {
        console.log(store.getters['worker/isAuthenticated']);
      next()
    } else {
        console.log(store.getters["worker/isAuthenticated"]);
      next('/login')
    }
  }
 */
const routes = [

    {
        path:'/',
        name: 'Home',
        component:Home,
        meta:{
            title:'Home'
        }
    },
    {
        path:'/rentings',
        name: 'Rentings',
        component:Rentings,
        meta:{
            title:'Rentings'
        }
    },
    {
        path:'/workers',
        name: 'Workers',
        component:Workers,
        meta:{
            title:'Workers'
        } ,
        beforeEnter: guardAuth.isAdmin 
    },
    {
        path:'/courts',
        name: 'Courts',
        component:Courts,
        meta:{
            title:'Courts'
        },
        beforeEnter: guardAuth.isAuthenticated 
    },
    {
        path:'/addworker',
        name: 'Add Worker',
        component:AddWorker,
        meta:{
            title:'Add worker'
        } ,
        beforeEnter: guardAuth.isAdmin 
    },
    {
        path:'/login',
        name: 'Login',
        component:Login,
        meta:{
            title:'Login'
        }
    }


]





const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
  });


  export default router;



