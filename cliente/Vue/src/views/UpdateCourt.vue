<template>
  <div class="container">
    <div class="row">
      <div class="col p-3">
        <h2>Actualizar Datos Pista</h2>
      </div>
    </div>
    <div class="row">
      <div class="col">
         <div class="form-group">
          <label htmlFor="id" > ID:</label>
          <input type="text" class="form-control" v-model="state.courtitemlocal.id" readonly />
        </div>

       <div class="form-group">
          <span>Pista:</span><br>
             <select v-model="state.courtitemlocal.name">
            <option disabled value="">seleccione pista</option>
            <option>Kapitalia</option>
            <option>Padelitis</option>
            <option>La baronia</option>
            <option>Maxichina</option>
            <option>Carnicas Levante</option>

          </select>
        </div>
        <div class="form-group">
          <span>Horario:</span><br>
             <select v-model="state.courtitemlocal.schedule">
            <option disabled value="">Selecciona un Horario</option>
            <option>08:00 / 9:30</option>
            <option>09:30 / 11:00</option>
            <option>11:00 / 12:30 </option>
            <option>12:30 / 14:00 </option>
            <option>17:00 / 18:30 </option>
            <option>18:30 / 20:00 </option>
            <option>20:00 / 21:30 </option>
          </select>
        </div>
        <div class="form-group">
            <label htmlFor="date">FECHA :</label>
            <textarea class="form-control" rows="1" id="date" v-model="state.courtitemlocal.date"></textarea>  
        </div>
        <div class="form-group">
            <label htmlFor="A_drive">JUGADOR A_DRIVE :</label>
            <textarea class="form-control" rows="1" id="A_drive" v-model="state.courtitemlocal.A_drive"></textarea>  
        </div>
        <div class="form-group">
            <label htmlFor="A_reves">JUGADOR A_REVES :</label>
            <textarea class="form-control" rows="1" id="A_reves" v-model="state.courtitemlocal.A_reves"></textarea>  
        </div>
        <div class="form-group">
            <label htmlFor="B_drive">JUGADOR B_DRIVE :</label>
            <textarea class="form-control" rows="1" id="B_drive" v-model="state.courtitemlocal.B_drive"></textarea>  
        </div>
        <div class="form-group">
            <label htmlFor="B_reves">JUGADOR B_REVES :</label>
            <textarea class="form-control" rows="1" id="B_reves" v-model="state.courtitemlocal.B_reves"></textarea>  
        </div>
        <div class="form-group">
            <button type="button" class="btn btn-success m-1" @click="updateCourt">Update</button>
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

export default {
    setup() {
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();
        console.log(store.state.court.courtlist[0].id);
        console.log(currentRoute.params.id);
       
        const courtitem = store.state.court.courtlist.find((item)=> item.id == currentRoute.params.id)
        console.log(courtitem);
        const state = reactive({ 
            courtitemlocal : { ...courtitem } 
        });
     
        const updateCourt = () => {
          console.log("UPDATE COURT VIEW");
            router.push({ name:"courtList" });
            store.dispatch(Constant.UPDATE_COURT, { courtitem: state.courtitemlocal });
        }

        const cancel = () => {
            router.push({ name:"courtList"});
        }

        return { state, updateCourt, cancel };
    }
}
</script>

<style>

</style>