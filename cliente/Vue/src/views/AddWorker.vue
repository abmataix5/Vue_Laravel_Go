<template>
    <div class="container">
    <form @submit.prevent="onSubmit">
    <!-- Nombre del trabajdor-->
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Nombre del trabajador:</label><input class="form-control" placeholder="Marc García Márquez" type="text" v-model="v$.form.username.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.username.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

  <!-- Dirección -->

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Dirección del trabajador:</label><input class="form-control" placeholder="Calle Gran Via nº3" type="text" v-model="v$.form.address.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.address.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

  <!-- Número del trabajador -->

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Número de telefono:</label><input class="form-control" placeholder="644926667" type="text" v-model="v$.form.phone.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.phone.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

   <!-- Puesto del trabajador -->

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Puesto del trabajador:</label><input class="form-control" placeholder="Professor, gerente..." type="text" v-model="v$.form.worker_type.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.worker_type.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

    </div>


    <!-- Email  -->
    <div class="form-group">
      <label for="">Email:</label><input class="form-control" placeholder="Introduzca elemail de contacto del trabajador" type="email" v-model="v$.form.email.$model">
      <div class="pre-icon os-icon os-icon-email-2-at2"></div>
      <!-- Error Message -->
        <div class="input-errors" v-for="(error, index) of v$.form.email.$errors" :key="index">
          <div class="error-msg">{{ error.$message }}</div>
        </div>
    </div>


    <!-- Contraseña-->
    <div class="row">

      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Contraseña</label><input class="form-control"  type="password" v-model="v$.form.password.$model">
          <div class="pre-icon os-icon os-icon-fingerprint"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.password.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

      <div class="col-sm-6">

        <div class="form-check">
          <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault" v-model="v$.form.admin.$model">
          <label class="form-check-label" for="flexCheckDefault"> Concedir permisos de administrador </label>
        </div>

      </div>


    </div>

    <!-- Submit Button -->
    <div class="buttons-w">
      <button class="btn btn-primary" :disabled="v$.form.$invalid" >Añadir nuevo trabajador</button> <!-- No dejara enviar si todos los inputs no son correctos -->
    </div>
    
  </form>
  </div>
</template>

<script>


import useVuelidate from "@vuelidate/core";
import { reactive, computed } from "vue";
import { required, email, minLength } from "@vuelidate/validators";
import router from '@/router'
import { useStore } from "vuex";
import Constant from "../Constant";


export function validName(name) {
 let validNamePattern = new RegExp("^[a-zA-Z]+(?:[-'\\s][a-zA-Z]+)*$");
 if (validNamePattern.test(name)){
   return true;
 }
 return false;
}
 
export function validPhone(phone) {
 let validPhonePattern = new RegExp("^[679]{1}[0-9]{8}$");
 if (validPhonePattern.test(phone)){
   return true;
 }
 return false;
}



export default {

  setup() {

    const store = useStore();

    const state = reactive({
      form: {
        username: "",
        email: "",
        password: "",
        address: "",
        phone:"",
        worker_type:"",
        antiguedad:"0 años",
        admin:""
      }
    });
    const rules = computed(() => {
      return {
        form: {
          username: {

          required, name_validation: {
           $validator: validName,
           $message: 'Nombre del trabajador incorrecto, el nombre no puede contener guiones ni numeros.'
         }

          },
          phone: {
          required, phone_validation: {
            $validator: validPhone,
            $message: 'El formato del número de telefono es incorrecto'
          }
          },
          worker_type:{required},
          email: { required, email },
          password: { required, min: minLength(6) },
          address: {required},
          admin:{}
        },
      };
    });

    const v$ = useVuelidate(rules, state);
    return { state, v$ , store};
  },
  methods: {

    onSubmit() {
      this.v$.$validate();

      if (!this.v$.$error) {

        if(this.state.form.admin == true){
          this.state.form.admin = 'true'
        }else{
          this.state.form.admin = 'false'
        }
   
        this.store.dispatch("worker/" + Constant.ADD_WORKER, {
            workeritem: this.state.form,
        });
        alert("Usuario registrado");
   router.push('/workers');
      } else {
      
        alert("Error");
      }
    },
  },
};
</script>



<style>

</style>
