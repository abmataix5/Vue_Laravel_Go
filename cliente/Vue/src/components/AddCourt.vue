<template>
    
<form>

  <div class="form-group row">
    <label for="pista" class="col-sm-2 col-form-label">Pista</label>
    <div class="col-sm-8">
      <input type="text" class="form-control" id="pista" placeholder="Pista Azul Central" >
    </div>
  </div>

  <div class="form-group row">
    <label for="horario" class="col-sm-2 col-form-label">Horario</label>
    <div class="col-sm-8">
      <input type="text" class="form-control" id="horario" placeholder="20:00-21:30" >
    </div>
  </div>

  <div class="form-group row">
    <label for="fecha" class="col-sm-2 col-form-label">Fecha de la partida:</label>
    <div class="col-sm-8">
      <input type="text" class="form-control" id="fecha" placeholder="21-12-2021">
    </div>
  </div>


  <div class="form-group row">
    <div class="col-sm-15">
      <button type="submit" class="btn btn-primary">Crear nueva pista</button>
    </div>
  </div>

</form>

</template>

<script>

import Constant from '../Constant';
import { reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {

    setup(){

        const store = useStore();
        const router = useRouter();

        const state = reactive({ 
            todoitemlocal : { 
              name:"", 
              email:"", 
              phone:"", 
              adresss:"",
              worker_type:"",
              antiguedad: 0 } 
        });


         const addWorker = () => {
          if (state.todoitemlocal.name.trim().length >= 2) {
            store.dispatch(Constant.ADD_WORKER, { todoitem : state.todoitemlocal })
            router.push({ name:"todoList" });
          } else {
            alert('Error en añadir un nuevo usuario');
          }
        }

        return {state, addWorker}

    }
    
}
</script>

<style>

</style>