<template>
 <form @submit.prevent="onSubmit">
  <div class="container">
    <div class="row">
      <div class="col p-3">
        <h2>Añadir Socio</h2>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div class="form-group">
          <label for=""> Nombre:</label><input class="form-control" placeholder="Enter first name" type="text" v-model="v$.form.name.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.name.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
        <div class="form-group">
            <label for="">Apellidos:</label><input class="form-control" placeholder="apellidos" type="text" v-model="v$.form.lastname.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.lastname.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
         <div class="form-group">
            <label for=""> Email:</label><input class="form-control" placeholder="ejemplo@ejemplo.com" type="text" v-model="v$.form.email.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.email.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
         <div class="form-group">
            <label for=""> Telefono:</label><input class="form-control" placeholder="telefono" type="text" v-model="v$.form.phone.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.phone.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>

        <div class="form-group">
          <span>Posición:</span><br>
             <select v-model="v$.form.position.$model">
            <option disabled value="">Seleccione una posición</option>
            <option>Drive</option>
            <option>Revés</option>
            <option>Ambas</option>
          </select>
        
        </div>
        <div class="form-group">
                <label htmlFor="image">Image:</label>
                <input type="file" class="form-control" id="image" ref="image"/>
        </div>
        <div class="form-group">
            <button class="btn btn-primary"  >Añadir Socio</button>
            <button type="button" class="btn btn-primary m-1" @click="cancel">Cancelar</button>
        </div>
      </div>
    </div>
  </div> 
 </form>
</template>

<script>
import Constant from '../Constant';
import { reactive, computed } from 'vue'
import { required, email } from "@vuelidate/validators";
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';
import router from '@/router'
import { ref } from 'vue';
import useVuelidate from "@vuelidate/core";

export function validName(name) {
 let validNamePattern = new RegExp("^[a-zA-Z]+(?:[-'\\s][a-zA-Z]+)*$");
 if (validNamePattern.test(name)){
   return true;
 }
 return false;
}

export function validLastname(lastname) {
 let validLastnamePattern = new RegExp("^[a-zA-Z]+(?:[-'\\s][a-zA-Z]+)*$");
 if (validLastnamePattern.test(lastname)){
   return true;
 }
 return false;
}
 
export function validPhone(phone) {
 let validPhonePattern = new RegExp("^[0-9]{8}$");
 if (validPhonePattern.test(phone)){
   return true;
 }
 return false;
}

export default {
    setup() {
        console.log("entra setup")
        const store = useStore();
        const router = useRouter();

        const image = ref();

        const state = reactive({ 
            form : { 
                name: "",
                lastname:"",
                email: "",
                phone:"",
                position:"",
                image:""
                }  
        });
        
        const rules = computed(() => {
          return {
            form: {
              name: {

              required, name_validation: {
              $validator: validName,
              $message: 'Nombre Socio incorrecto, el nombre no puede contener guiones ni numeros.'
              }

              },
              lastname: {

              required, lastname_validation: {
              $validator: validLastname,
              $message: 'Apellidos Socio incorrectos, los apellidos no pueden contener guiones ni numeros.'
              }

              },
              phone: {
              required, phone_validation: {
                $validator: validPhone,
                $message: 'El formato del número de telefono es incorrecto'
              }
              },
              position:{required},
              email: { required, email },
              password: {},
              admin: {},
              image: {}
            },
          };
        });

        const v$ = useVuelidate(rules, state);

        
        const cancel = () => {
            router.push({ name:"partnerList"});
        }

        return { state, image ,cancel, v$ , store}
    },
    methods:{
       onSubmit() {
        console.log("entra on submit");
        this.v$.$validate();

        if (!this.v$.$error) { //si no se produce ningun error de validación.

          console.log(router);
          if(this.image.value != undefined){
              
              this.state.form.image = this.image.files[0]; //guardamos los datos del archivo.
      
          }else{
              this.state.partneritemvalue.image = null; // valor por defecto.
          }
              console.log("realiza dispatch");
              this.store.dispatch(Constant.ADD_PARTNER, { partneritem : this.state.form })
              alert("Usuario registrado");
          
        } else {
          alert("Error");
        }
    },
    },
};
</script>

<style>

</style>