import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/Home';
import Workers from '../views/Workers';
import AddWorker from '../views/AddWorker';
import Login from '../views/Login'
import guardAuth from "../guards/guardAuth";
import CourtList from '../views/CourtList';
import CourtAdd from '../views/CourtAdd';
import UpdateCourt from '../views/UpdateCourt';
import PartnerList from '../views/PartnerList';
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
        } ,
        beforeEnter: guardAuth.isAdmin 
    },
    {
        path:'/partner',
        name: 'partnerList',
        component:PartnerList,
        meta:{
            title:'Partner'
        },
        // beforeEnter: guardAuth.workerAuthenticated 
    },
    {
        path:'/court',
        name: 'courtList',
        component:CourtList,
        meta:{
            title:'Courts'
        },
        beforeEnter: guardAuth.workerAuthenticated 
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
    },
    {
        path:'/loginLaravel',
        name: 'LoginLaravel',
        component:Login,
        meta:{
            title:'LoginLaravel'
        }
    },


]





const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
  });


  export default router;



