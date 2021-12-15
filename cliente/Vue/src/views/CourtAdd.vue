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
          <label htmlFor="user">Nombre :</label>
          <input type="text" class="form-control" v-model="state.courtitemvalue.name" />
        </div>
        <div class="form-group">
            <label htmlFor="email">Schedule :</label>
             <input type="text" class="form-control" v-model="state.courtitemvalue.schedule" />
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