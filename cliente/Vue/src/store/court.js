import Constant from '../Constant';
import CourtService from '@/services/CourtService'

export const courtStore = {
    
    namespaced: true,
    state: {
      courtlist:[]
    },
    mutations: {
        [Constant.ADD_COURT]: (state, payload) => {
            console.log("MUTATION ADDCOURT");
            state.courtlist.push({ ...payload});
        },
        [Constant.UPDATE_COURT] :(state,payload)=> {
            let index = state.courtlist.findIndex((item)=>item.id === payload.courtitem.id);
            state.courtlist[index] = payload.courtitem;
          },
        [Constant.DELETE_COURT]: (state, payload) => {
            let index = state.courtlist.findIndex((item) => item.id === payload.id);
            state.courtlist.splice(index, 1);
        },
        [Constant.INITIALIZE_COURTITEM]: (state, payload) => {
          
            if (payload) {
                console.log("AAAAA");
                console.log(state);
                console.log("payload data: ")
                console.log(payload);
                console.log(payload.data);
                state.courtlist = payload.data.data;
            } else {
                state.courtlist = { name: "",schedule: "",date: "", A_drive: "",A_reves: "", B_drive: "",B_reves: "" };
            }
        }
    },
    actions: {

        addCourt(store,payload){
            console.log("ACTION ADDCOURT");
            console.log(payload.courtitem);
            CourtService.add(payload.courtitem)
            .then(function (data_court) {
                console.log("ACTION ADDCOURT .THEN");
                store.commit(Constant.ADD_COURT, data_court);
               
            })
            .catch(function (error) {
                console.log(error)
            })
        },

        initializeCourtItem(store){

            CourtService.get("/courts")
                .then(function (data_court) {
                    console.log("INITIALIZE!");
                    store.commit(Constant.INITIALIZE_COURTITEM, data_court.data);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        deleteCourt(store, payload){

            CourtService.delete(payload.id)
                    .then(function (res) {
                        console.log(res);
                        store.commit(Constant.DELETE_COURT, payload);
                    }
                    )
                    .catch(function (error) {
                        console.log(error)
                    })
        },
        updateCourt(store, payload){
            console.log("ACTION UPDATE COURT");
            console.log(payload);
            CourtService.update(payload.courtitem.id, payload.courtitem)
                    .then(function (res) {
                        console.log(res);
                        console.log(res.data);
                        store.commit(Constant.UPDATE_COURT, payload);
                    }
                    )
                    .catch(function (error) {
                        console.log(error)
                    })
        },
    },
    getters: {
        getCourts(state) {
            console.log(state.courtlist);
            return state.courtlist;
        }
    }

}