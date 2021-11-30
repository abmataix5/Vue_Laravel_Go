<template>
    
   <form>

  <div class="form-group row">
    <label for="name" class="col-sm-2 col-form-label">Nombre</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="name" placeholder="Nombre del trabajador" v-model="state.workeritemvalue.name">
    </div>
  </div>

  <div class="form-group row">
    <label for="email" class="col-sm-2 col-form-label">Email</label>
    <div class="col-sm-10">
      <input type="email" class="form-control" id="email" placeholder="trabajador1@gmail.com" v-model="state.workeritemvalue.email">
    </div>
  </div>

  <div class="form-group row">
    <label for="phone" class="col-sm-2 col-form-label">Número de telefono:</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="phone" placeholder="655489768" v-model="state.workeritemvalue.phone">
    </div>
  </div>

  <div class="form-group row">
    <label for="location" class="col-sm-2 col-form-label">Localización</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="location" placeholder="Localización" v-model="state.workeritemvalue.adress">
    </div>
  </div>

  <div class="form-group row">
    <label for="type" class="col-sm-2 col-form-label">Puesto del trabajador</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="type" placeholder="Localización" v-model="state.workeritemvalue.worker_type">
    </div>
  </div>

  <div class="form-group row">
    <label for="antiguedad" class="col-sm-2 col-form-label">Antiguedad</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="antiguedad" placeholder="Localización" v-model="state.workeritemvalue.antiguedad">
    </div>
  </div>


  <div class="form-group row">
    <div class="col-sm-10">
      <button type="submit" class="btn btn-primary" @click="newWorker">Crear nuevo usuario</button>
    </div>
  </div>
</form>

</template>

<script>

import Constant from "../Constant";
import { reactive } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

export default {
    
  setup() {
    const store = useStore();
    const router = useRouter();
    const state = reactive({
      workeritemvalue: {
        username: "",
        email: "",
        phone: "",
        address: "",
        worker_type: "",
        antiguedad: "0 años" ,
      },
    });

    const newWorker = () => {

      if (state.workeritemvalue.name != undefined) {
        store.dispatch("worker/" + Constant.ADD_WORKER, {
          workeritem: state.workeritemvalue,
        });
        router.push({ name: "workerList" });
      } else {
        alert("Error");
      }
    };
   
    return { state, newWorker};
  }
    
    
}
</script>

<style>

</style>
