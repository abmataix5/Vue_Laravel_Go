import Constant from '../Constant';
import PartnerService from '@/services/PartnerService'

export const partnerStore = {
    
    namespaced: true,
    state: {
      partnerlist:[]
    },
    mutations: {
        [Constant.ADD_PARTNER]: (state, payload) => {
            console.log("MUTATION ADDPARTNER");
            state.partnerlist.push({ ...payload});
        },
        [Constant.UPDATE_PARTNER] :(state,payload)=> {
            let index = state.partnerlist.findIndex((item)=>item.id === payload.partneritem.id);
            state.partnerlist[index] = payload.partneritem;
          },
        [Constant.DELETE_PARTNER]: (state, payload) => {
            let index = state.partnerlist.findIndex((item) => item.id === payload.id);
            state.partnerlist.splice(index, 1);
        },
        [Constant.INITIALIZE_PARTNERITEM]: (state, payload) => {
            console.log("ENTRA MUTATIONS INITAILIZE");
            // if (payload) {
            //     console.log("AAAAA");
                console.log(state);
                // console.log("payload data: ")
                console.log(payload.data);
            //     state.partnerlist = payload.data;
            // } else {

            //     //CAMBIAR POR EL MODELO PARTNER
            //     state.partnerlist = { name: "",schedule: "",date: "", A_drive: "",A_reves: "", B_drive: "",B_reves: "" };
            // }
        }
    },
    actions: {

        addPartner(store,payload){
            console.log("ACTION ADDPARTNER");
            console.log(payload.partneritem);
            PartnerService.add(payload.partneritem)
            .then(function (data_partner) {
                console.log("ACTION ADDPARTNER .THEN");
                store.commit(Constant.ADD_PARTNER, data_partner);
               
            })
            .catch(function (error) {
                console.log(error)
            })
        },

        
        initializePartnerItem(store){
            console.log("INITIALIZE");
            PartnerService.get("/partner")
                .then(function (data_partner) {
                    console.log("INIT PARTNERITEM!");
                    
                    store.commit(Constant.INITIALIZE_PARTNERITEM, data_partner.data);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        deletePartner(store, payload){

            PartnerService.delete(payload.id)
                    .then(function (res) {
                        console.log(res);
                        store.commit(Constant.DELETE_PARTNER, payload);
                    }
                    )
                    .catch(function (error) {
                        console.log(error)
                    })
        },
        updatePartner(store, payload){
            console.log("ACTION UPDATE PARTNER");
            console.log(payload);
            PartnerService.update(payload.partneritem.id, payload.partneritem)
                    .then(function (res) {
                        console.log(res);
                        console.log(res.data);
                        store.commit(Constant.UPDATE_PARTNER, payload);
                    }
                    )
                    .catch(function (error) {
                        console.log(error)
                    })
        },
    },
    getters: {
        getPartner(state) {
            console.log(state.partnerlist);
            return state.partnerlist;
        }
    }

}