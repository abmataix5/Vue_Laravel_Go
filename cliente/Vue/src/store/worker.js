import Constant from '../Constant';
import UserService from '@/services/UserService'

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
                state.workerlist = {id: "",name: "",email: "", phone: "",address: "", worker_type: "",antiguedad: "" };
            }
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
                    console.log(data_worker.data)
                    store.commit(Constant.INITIALIZE_WORKERITEM, data_worker.data);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        [Constant.DELETE_WORKER]: (store, payload) => {
    
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
    },
    getters: {
        getWorkers(state) {
            return state.workerlist;
        }
    }

}