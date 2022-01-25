import Constant from '../Constant';
import CourtService from '@/services/CourtService'

export const rentStore = {
    
    namespaced: true,
    state: {
      rentlist:[]
    },
    mutations: {
        [Constant.ADD_RENT]: (state, payload) => {
            console.log("MUTATION ADDCOURT");
            state.renttlist.push({ ...payload});
        },
        [Constant.UPDATE_RENT] :(state,payload)=> {
            let index = state.rentlist.findIndex((item)=>item.id === payload.rentitem.id);
            state.rentlist[index] = payload.rentitem;
          },
        [Constant.DELETE_RENT]: (state, payload) => {
            let index = state.rentlist.findIndex((item) => item.id === payload.id);
            state.rentlist.splice(index, 1);
        },
        [Constant.INITIALIZE_RENTITEM]: (state, payload) => {
          
            if (payload) {
                console.log("AAAAA");
                console.log(state);
                console.log("payload data: ")
                console.log(payload.data.data);
                state.rentlist = payload.data.data;
            } else {
                state.rentlist = { id:"", name: "",schedule: "",date: "", A_drive: "",A_reves: "", B_drive: "",B_reves: "" };
            }
        }
    },
    actions: {

        addRent(store,payload){
            console.log("ACTION ADDRENT");
            console.log(payload.rentitem);
            CourtService.add(payload.rentitem)
            .then(function (data) {
                console.log("ACTION ADDRENT .THEN");
                store.commit(Constant.ADD_COURT, data);
               
            })
            .catch(function (error) {
                console.log(error)
            })
        },

        initializeRentItem(store){

            CourtService.get("/rent")
                .then(function (data) {
                    console.log("INITIALIZE RENT ITEM!");
                    console.log(data.data.data)
                    store.commit(Constant.INITIALIZE_RENTITEM, data.data);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        deleteRent(store, payload){

            CourtService.delete(payload.id)
                    .then(function (res) {
                        console.log(res);
                        store.commit(Constant.DELETE_RENT, payload);
                    }
                    )
                    .catch(function (error) {
                        console.log(error)
                    })
        },
        updateRent(store, payload){
            console.log("ACTION UPDATE RENT");
            console.log(payload);
            CourtService.update(payload.rentitem.id, payload.rentitem)
                    .then(function (res) {
                        console.log(res);
                        console.log(res.data);
                        store.commit(Constant.UPDATE_RENT, payload);
                    }
                    )
                    .catch(function (error) {
                        console.log(error)
                    })
        },
    },
    getters: {
        getRent(state) {
            console.log("GETTERS!")
            console.log(state);
            return state.rentlist;
        }
    }

}