import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/Home';
import Rentings from '../views/Rentings';
import Workers from '../views/Workers';
import Courts from '../views/AddCourt';
import AddWorker from '../views/AddWorker';
import Login from '../views/Login'

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
        }
    },
    {
        path:'/courts',
        name: 'Courts',
        component:Courts,
        meta:{
            title:'Courts'
        }
    },
    {
        path:'/addworker',
        name: 'Add Worker',
        component:AddWorker,
        meta:{
            title:'Add worker'
        }
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



