<template>
    <div class="container">
    <form @submit.prevent="onSubmit">
    <!-- Nombre del trabajdor-->
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Nombre del trabajador:</label><input class="form-control" placeholder="Marc García Márquez" type="text" v-model="state.workeritemlocal.name">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
      
        </div>
      </div>

  <!-- Dirección -->

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Dirección del trabajador:</label><input class="form-control" placeholder="Calle Gran Via nº3" type="text" v-model="state.workeritemlocal.name">

        </div>
      </div>

  <!-- Número del trabajador -->

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Número de telefono:</label><input class="form-control" placeholder="644926667" type="text" v-model="state.workeritemlocal.name">
       
        </div>
      </div>

   <!-- Puesto del trabajador -->

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Puesto del trabajador:</label><input class="form-control" placeholder="Professor, gerente..." type="text" v-model="state.workeritemlocal.name">
      
        </div>
      </div>

    </div>


    <!-- Email  -->
    <div class="form-group">
      <label for="">Email:</label><input class="form-control" placeholder="Introduzca elemail de contacto del trabajador" type="email" v-model="state.workeritemlocal.name">
      <div class="pre-icon os-icon os-icon-email-2-at2"></div>
 
    </div>


    <!-- Contraseña-->
    <div class="row">

      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Contraseña</label><input class="form-control"  type="password" v-model="state.workeritemlocal.name">
          <div class="pre-icon os-icon os-icon-fingerprint"></div>
        </div>
      </div>

      <div class="col-sm-6">

        <div class="form-check">
          <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault" v-model="state.workeritemlocal.name">
          <label class="form-check-label" for="flexCheckDefault"> Concedir permisos de administrador </label>
        </div>

      </div>


    </div>

    <!-- Submit Button -->
    <div class="buttons-w">
      <button class="btn btn-primary" :disabled="v$.form.$invalid" >Actualizar</button> <!-- No dejara enviar si todos los inputs no son correctos -->
    </div>
    
  </form>
  </div>
</template>

<script>

import useAddWorkers from "../composable/updateWorkers";
export default {
  setup() {
 const { state, updateWorker, cancel } = useAddWorkers();
    return { state, updateWorker, cancel };
  },
};
</script>



<style>

</style>
