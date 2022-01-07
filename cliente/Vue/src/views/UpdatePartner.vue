<template>
  <div class="container">
    <div class="row">
      <div class="col p-3">
        <h2>Actualizar Datos Socio</h2>
      </div>
    </div>
    <div class="row">
      <div class="col">
         <div class="form-group">
          <label htmlFor="id"> ID:</label>
          <input type="text" class="form-control" v-model="state.partneritemlocal.id" readonly/>
        </div>

        <div class="form-group">
          <label htmlFor="name"> Nombre:</label>
          <input type="text" class="form-control" v-model="state.partneritemlocal.name" />
        </div>
        <div class="form-group">
            <label htmlFor="email">Email :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemlocal.email"></textarea>  
        </div>
        
        <div class="form-group">
            <button type="button" class="btn btn-primary m-1" @click="updatePartner">Update</button>
            <button type="button" class="btn btn-primary m-1" @click="cancel">Cancel</button>
        </div>
      </div>
    </div>
  </div> 
</template>

<script>
import Constant from '../Constant';
import { reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router';

export default {
    setup() {
        console.log("entra setup")
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();
        console.log(store.state.partner.partnerlist[0].id);
        console.log(currentRoute.params.id);
       
        const partneritem = store.state.partner.partnerlist.find((item)=> item.id == currentRoute.params.id)
        console.log(partneritem);
        const state = reactive({ 
            partneritemlocal : { ...partneritem } 
        });
     
        const updatePartner = () => {
          console.log("UPDATE PARTNER VIEW");
            router.push({ name:"partnerList" });
            store.dispatch(Constant.UPDATE_PARTNER, { partneritem: state.partneritemlocal });
        }

        const cancel = () => {
            router.push({ name:"partnerList"});
        }

        return { state, updatePartner, cancel };
    }
}
</script>

<style>

</style>