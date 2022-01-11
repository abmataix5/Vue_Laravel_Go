<template>
  <div class="container">
    <div class="row">
      <div class="col p-3">
        <h2>Añadir Socio</h2>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div class="form-group">
          <label htmlFor="name">Nombre :</label>
          <input type="text" class="form-control" v-model="state.partneritemvalue.name" />
        </div>
        <div class="form-group">
            <label htmlFor="lastname">Apellidos :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemvalue.lastname"></textarea>  
        </div>
         <div class="form-group">
            <label htmlFor="email">Email :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemvalue.email"></textarea>  
        </div>
         <div class="form-group">
            <label htmlFor="phone">Telefono contacto :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemvalue.phone"></textarea>  
        </div>
        <!-- <div class="form-group">
            <label htmlFor="position">Posición :</label>
            <textarea class="form-control" rows="1" v-model="state.partneritemvalue.position"></textarea>  
        </div> -->
        <div class="form-group">
          <span>Posición:</span><br>
          <select v-model="state.partneritemvalue.position">
            <option disabled value="">Seleccione una posición</option>
            <option>Drive</option>
            <option>Revés</option>
            <option>Ambas</option>s
          </select>
        
        </div>
         <!-- <div class="form-group">
            <label htmlFor="position">Position :</label>
             <input type="text" class="form-control" v-model="state.partneritemvalue.date" />
        </div> -->
       
        <div class="form-group">
            <button type="button" class="btn btn-primary m-1" @click="addPartner">Añadir Socio</button>
            <button type="button" class="btn btn-primary m-1" @click="cancel">Cancelar</button>
        </div>
      </div>
    </div>
  </div> 
</template>

<script>
import Constant from '../Constant';
import { reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    setup() {
        console.log("entra setup")
        const store = useStore();
        const router = useRouter();

        const state = reactive({ 
            partneritemvalue : { 
                name: "",
                lastname:"",
                email: "",
                phone:"",
                position:""
                }  
        });

        const addPartner = () => {
          
        console.log("entra addPartner");
        console.log(state.partneritemvalue);

          if (state.partneritemvalue.name != undefined) {
            store.dispatch(Constant.ADD_PARTNER, { partneritem : state.partneritemvalue })
            // router.push({ name:"partnerList" });
          } else {
            alert('Por favor Inserta Campos');
          }
        }
        const cancel = () => {
            router.push({ name:"partnerList"});
        }

        return { state, addPartner, cancel }
    }
}
</script>

<style>

</style>