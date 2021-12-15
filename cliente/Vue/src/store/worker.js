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
            } else {
                state.workerlist = {id: "",username: "",email: "", phone: "",address: "", worker_type: "",antiguedad: "", password:"",admin:"" };
            }
        },
        [Constant.LOGIN_WORKER]: (state, payload) => {
         
            setStore('token', payload.token); /* Guardamos token en localstorage */
            setStore('admin', payload.admin); /* Guardamos si es admin o no */

            state.user = payload;
            console.log(state.user.admin);  /* Guardamos en el state el usuario logueado */
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

          console.log(payload.id);
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
            console.log(payload)
            UserService.loginWorker(payload.workeritem)
                .then(function (worker_logued) {
                    console.log(worker_logued.data.response)
                    store.commit(Constant.LOGIN_WORKER, worker_logued.data.response);
                   
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
     
        isAuthenticated: (state) => {
      
            if(state.user){
                return true;
            }else{
                return false;
            }
        },
        isAdmin: (state) => {

            const isAdmin = state.user.admin;
            console.log(isAdmin);
            
            if(isAdmin == true){

                return true;

            }else{
                return false;
            }
        },
    }

}