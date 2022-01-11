import Constant from '../Constant';
import PartnerService from '@/services/PartnerService'

export const partnerStore = {
    
    namespaced: true,
    state: {
      partnerlist:[],
      partnerdetail:[]
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
            if (payload) {
                console.log("AAAAA");
             
                console.log("payload data: ")
                console.log(payload.data)
                state.partnerlist = payload.data;
                
            } else {
                state.partnerlist = { id: "", name: "",email: ""};
            }
        },
        [Constant.DETAIL_PARTNER]: (state, payload) => {
            console.log("ENTRA MUTATIONS DETAIL");
            if (payload) {
                // console.log("AAAAA");
             
                // console.log("payload data: ");
                // console.log(payload);
                state.partnerdetail = payload;
                // console.log(state.partnerdetail);
            } else {

            
                state.partnerdetail = { id: "", name: "",lastname:"", email: ""};
            }
        }
    },
    actions: {

        addPartner(store,payload){
            console.log("ACTION ADDPARTNER");
            //EL valor password está en default. Ya que los socios son registrados por personal autorizado.
            payload.partneritem.password= "default"; //valor por defecto para los socios.
            payload.partneritem.admin="false"; //valor por defecto para los socios.
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
            PartnerService.get()
                .then(function (data_partner) {
                    console.log("INIT PARTNERITEM!");
                    console.log(data_partner)
                    console.log(data_partner.data.data)
                    store.commit(Constant.INITIALIZE_PARTNERITEM, data_partner.data.data);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        detailPartner(store,payload){
            console.log("DETAIL PARTNER");
            console.log(payload);

            PartnerService.getOne(payload)
                .then(function (data_partner) {

                    console.log(data_partner.data.data)
                  
                    store.commit(Constant.DETAIL_PARTNER, data_partner.data.data);
                   
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        deletePartner(store, payload){

            PartnerService.delete(payload.id)
                    .then(function (res) {
                        console.log(res);
                        console.log(res.message)
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
        getPartners(state) {
            console.log(state.partnerlist);
            return state.partnerlist;
        },
        detailPartners(state) {
            console.log(state.partnerdetail);
            return state.partnerdetail;
        }
    }

}