import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/Home';

import Workers from '../views/Workers';
import AddWorker from '../views/AddWorker';

import Login from '../views/Login';

import CourtList from '../views/CourtList';
import CourtAdd from '../views/CourtAdd';
import UpdateCourt from '../views/UpdateCourt';

import RentList from '../views/RentList';

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
        path:'/workers',
        name: 'Workers',
        component:Workers,
        meta:{
            title:'Workers'
        }
    },
    {
        path:'/court',
        name: 'courtList',
        component:CourtList,
        meta:{
            title:'Courts'
        }
    },
    { 
        path:'/court/add', 
        name:'courtAdd', 
        component: CourtAdd,
        meta:{
            title:'CourtsAdd'
        } 
    },
    { 
        path:'/court/update/:id', 
        name:'updateCourt', 
        component: UpdateCourt,
        meta:{
            title:'UpdateCourts'
        } 
    },
    {
        path:'/rent',
        name: 'rentList',
        component:RentList,
        meta:{
            title:'Rentings'
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



