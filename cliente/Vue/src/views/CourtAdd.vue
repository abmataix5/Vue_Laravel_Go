<template>
  <div class="container">
    <div class="row">
      <div class="col p-3">
        <h2>Añadir Pista</h2>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div class="form-group">
          <span>Pista:</span><br>
             <select v-model="state.courtitemvalue.name">
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
             <select v-model="state.courtitemvalue.schedule">
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
            <label htmlFor="email">Fecha :</label>
             <input type="text" class="form-control" v-model="state.courtitemvalue.date" />
        </div>
        <div class="form-group">
            <label htmlFor="email">A_drive :</label>
             <input type="text" class="form-control" v-model="state.courtitemvalue.A_drive" />
        </div>
        <div class="form-group">
            <label htmlFor="email">A_reves :</label>
             <input type="text" class="form-control" v-model="state.courtitemvalue.A_reves" />
        </div>
        <div class="form-group">
            <label htmlFor="email">B_drive :</label>
             <input type="text" class="form-control" v-model="state.courtitemvalue.B_drive" />
        </div>
        <div class="form-group">
            <label htmlFor="email">B_reves :</label>
             <input type="text" class="form-control" v-model="state.courtitemvalue.B_reves" />
        </div>
        <div class="form-group">
            <button type="button" class="btn btn-primary m-1" @click="addCourt">Añadir Pista</button>
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
        const store = useStore();
        const router = useRouter();

        const state = reactive({ 
            courtitemvalue : { 
                name: "",
                schedule: "",
                date: "",
                A_drive: "",
                A_reves: "",
                B_drive: "",
                B_reves: "" }  
        });

        const addCourt = () => {
          
        console.log("entra addCourt");
        console.log(state.courtitemvalue);

          if (state.courtitemvalue.name != undefined) {
            store.dispatch(Constant.ADD_COURT, { courtitem : state.courtitemvalue })
            router.push({ name:"courtList" });
          } else {
            alert('Por favor Inserta Campos');
          }
        }
        const cancel = () => {
            router.push({ name:"courtList"});
        }

        return { state, addCourt, cancel }
    }
}
</script>

<style>

</style>