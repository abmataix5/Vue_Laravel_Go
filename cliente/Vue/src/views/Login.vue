<template>

<div class="container">
  <form @submit.prevent="onSubmit">
    <!-- Email -->
    <div class="form-group" :class="{ error: v$.form.email.$errors.length }">
      <label for="">Email</label>
      <input class="form-control" placeholder="Enter your username" type="email" v-model="v$.form.email.$model">
      <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- error message -->
      <div class="input-errors" v-for="(error, index) of v$.form.email.$errors" :key="index">
        <div class="error-msg">{{ error.$message }}</div>
      </div>
    </div>

    <!-- password -->
    <div class="form-group" :class="{ error: v$.form.password.$errors.length }">
      <label for="">Password</label>
      <input class="form-control" placeholder="Enter your password" type="password" v-model="v$.form.password.$model">
      <div class="pre-icon os-icon os-icon-fingerprint"></div>
          <!-- error message -->
      <div class="input-errors" v-for="(error, index) of v$.form.password.$errors" :key="index">
        <div class="error-msg">{{ error.$message }}</div>
      </div>
    </div>

    <!-- Submit Button -->
    <div class="buttons-w">
      <button :disabled="v$.form.$invalid" class="btn btn-primary">Login</button>
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


export default {

  setup() {

    const store = useStore();

    const state = reactive({
      form: {
     
        email: "",
        password: "",
      }
    });
    const rules = computed(() => {
      return {
        form: {
         
          email: { required, email },
          password: { required, min: minLength(6) },
        
        },
      };
    });

    const v$ = useVuelidate(rules, state);
    return { state, v$ , store};
  },
  methods: {

    onSubmit() {

      this.v$.$validate();
      console.log(this.v$.$error);
      if (!this.v$.$error) {
   
          this.store.dispatch("worker/" + Constant.LOGIN_WORKER, {
            workeritem: this.state.form,
          });
        alert("Usuario logueado");
     location.reload();
        router.push('/');
      } else {
      
        alert("Error login");
      }
    },
  },
};
</script>
<style>

.error-msg{
    color: red;
}
</style>