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
                
                state.partnerlist = payload.data;
                
            } else {
                state.partnerlist = { id: "", name: "",email: ""};
            }
        },
        [Constant.DETAIL_PARTNER]: (state, payload) => {
            console.log("ENTRA MUTATIONS DETAIL");
            if (payload) {
                
                state.partnerdetail = payload;
                
            } else {

            
                state.partnerdetail = { id: "", name: "",lastname:"", email: ""};
            }
        }
    },
    actions: {

        addPartner(store,payload){
            console.log("ACTION ADDPARTNER");
            //EL valor password está en default. Ya que los socios son registrados por personal autorizado.
            console.log(payload.partneritem);
            
            let formData = new FormData();
            formData.append("name", payload.partneritem.name);
            formData.append("lastname", payload.partneritem.lastname);
            formData.append("email", payload.partneritem.email);
            formData.append("position", payload.partneritem.position);
            formData.append("image", payload.partneritem.image);
            formData.append("phone", payload.partneritem.phone);
            formData.append("password", "default");
            formData.append("admin", "false");


            PartnerService.add(formData)
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
         
            let formData = new FormData();
            formData.append("name", payload.partneritem.name);
            formData.append("lastname", payload.partneritem.lastname);
            formData.append("email", payload.partneritem.email);
            formData.append("position", payload.partneritem.position);
            formData.append("image", payload.partneritem.image);
            formData.append("phone", payload.partneritem.phone);
            formData.append("password", "default");
            formData.append("admin", "false");

            PartnerService.update(payload.partneritem.id,formData)
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