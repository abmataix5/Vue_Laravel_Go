<template>
  <div class="container">
    <div class="row title">
            <div class="title_panelAdmin ml-3">
                <h1> Panel Administración Actualizar Datos Usuario </h1>
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
            <label htmlFor="lastname">Apellidos :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemlocal.lastname"></textarea>  
        </div>
         <div class="form-group">
            <label htmlFor="email">Email :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemlocal.email"></textarea>  
        </div>
         <div class="form-group">
            <label htmlFor="phone">Telefono contacto :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemlocal.phone"></textarea>  
        </div>
        <div class="form-group">
          <span>Posición:</span><br>
          <select v-model="state.partneritemlocal.position">
            <option disabled value="">Seleccione una posición</option>
            <option>Drive</option>
            <option>Revés</option>
            <option>Ambas</option>
          </select>
        </div>
         <div class="form-group">
                <label htmlFor="image">Image:</label><br>
                <img class="h-75" v-bind:src="state.partneritemlocal.image" /> 
                <input type="file" class="form-control" id="image" ref="image"/>
          </div>
        <div class="form-group">
            <button type="button" class="btn btn-success m-1" @click="updatePartner">Update</button>
            <button type="button" class="btn btn-danger m-1" @click="cancel">Cancel</button>
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
import { ref } from 'vue';

export default {
    setup() {
        console.log("entra setup")
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();
        const image = ref(null);

        console.log(store.state.partner.partnerlist[0].id);
        console.log(currentRoute.params.id);
       
        const partneritem = store.state.partner.partnerlist.find((item)=> item.id == currentRoute.params.id)
        console.log(partneritem);

        const state = reactive({ 
            partneritemlocal : { ...partneritem } 
        });
     
        const updatePartner = () => {
          
          console.log("entra update partner");

          if(image.value.files[0]!= undefined){
            console.log("entra image.value");
            state.partneritemlocal.image = image.value.files[0];
            // console.log("valor final");
            console.log(state.partneritemlocal.image);
          }else{
            state.partneritemlocal.image = null;
          }

          console.log("UPDATE PARTNER VIEW");
            // router.push({ name:"partnerList" });
            store.dispatch(Constant.UPDATE_PARTNER, { partneritem: state.partneritemlocal });
        }

        const cancel = () => {
            router.push({ name:"partnerList"});
        }

        return { state, updatePartner, cancel, image };
    }
}
</script>

<style>

</style>