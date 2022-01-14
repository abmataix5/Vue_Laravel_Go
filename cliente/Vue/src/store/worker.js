import Constant from '../Constant';
import UserService from '@/services/UserService'
import { setStore } from '@/services/jwt.service'

export const workerStore = {
    namespaced: true,
    state: {
    },
    mutations: {
        [Constant.ADD_WORKER]: (state, payload) => {
            state.workerlist.push({ ...payload});
        },
        [Constant.DELETE_WORKER]: (state, payload) => {
            let index = state.workerlist.findIndex((item) => item.id === payload.id);
            state.workerlist.splice(index, 1);
        },
        [Constant.INITIALIZE_WORKERITEM]: (state, payload) => {
          
            if (payload) {
                state.workerlist = payload;
                console.log(state.worker);
            } else {
                state.workerlist = {id: "",username: "",email: "", phone: "",address: "", worker_type: "",antiguedad: "", password:"",admin:"" };
            }
        },
        [Constant.LOGIN_WORKER]: (state, payload) => {
         
            setStore('token', payload.token); /* Guardamos token en localstorage */
            setStore('admin', payload.admin); /* Guardamos si es admin o no */
            setStore('username', payload.username); /* Guardamos si es admin o no */
            state.user = payload;
            console.log(state.user.admin);  /* Guardamos en el state el usuario logueado */
        },
        [Constant.LOGIN_LARAVEL_WORKER]: (state, payload) => {
            console.log("ENTRA MUTATIONS LOGIN LARAVEL WORKERS");

            console.log(state)
            console.log(payload)
            console.log(payload.access_token)
   
            setStore('token', payload.access_token); /* Guardamos token en localstorage */
            // setStore('admin', payload.admin); /* Guardamos si es admin o no */

            state.user = payload;
            console.log(state.user);  /* Guardamos en el state el usuario logueado */
        }
    },
    actions: {

        [Constant.ADD_WORKER]: (store, payload) => {
        
            UserService.createWorker(payload.workeritem)
                .then(function (worker) {
                    store.commit(Constant.ADD_WORKER, worker.data);
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        [Constant.INITIALIZE_WORKERITEM]: (store) => {

            UserService.getAll()
                .then(function (data_worker) {
                    console.log(data_worker.data.response)
                    store.commit(Constant.INITIALIZE_WORKERITEM, data_worker.data.response);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        [Constant.DELETE_WORKER]: (store, payload) => {

          console.log(payload);
            UserService.deleteWorker(payload.id)
                .then(function (res) {
                    if (res.statusText !== "OK") {
                        throw new Error()
                    }
                    console.log(res)
                    store.commit(Constant.DELETE_WORKER, payload);
                }
                )
                .catch(function (error) {
                    console.log(error)
                })
        },
        [Constant.LOGIN_WORKER]: (store,payload) => {
            console.log(payload.workeritem)
            UserService.loginWorker(payload.workeritem)
                .then(function (worker_logued) {
                    console.log(worker_logued.data.response)
                    store.commit(Constant.LOGIN_WORKER, worker_logued.data.response);
                    alert("Usuario logueado");
                    location.reload();
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        [Constant.LOGIN_LARAVEL_WORKER]: (store,payload) => {
            console.log("ENTRA LOGIN LARAVEL WORKERS");
            console.log(payload.workeritem)
            UserService.loginLaravelWorker(payload.workeritem)
                .then(function (data) {
                    console.log("vuelve LoginWorkerLaravel")
                    console.log(data.data)
                    store.commit(Constant.LOGIN_LARAVEL_WORKER, data.data);
                    location.reload();
                })
                .catch(function (error) {
                    console.log(error)
                })
        }
    },
    getters: {
        getWorkers(state) {

            return state.workerlist;
        },
     
        isWorkerAuthenticated: () => {

            if(localStorage.getItem('token')){
                return true;
            }else{
                return false;
            }
        },
        isAdmin: () => {

            if(localStorage.getItem('admin') === null){
                return null;
            }

            if(localStorage.getItem('admin') === 'true'){
                return true;
            }else{
                return false
            }
        
        },
        userLogued(){
            if(localStorage.getItem('token')){
                return true;
            }else{
                return false;
            }
        },
    }

}