<template>
    <nav class="navbar navbar-expand-lg bg-dark navbar-dark">
        <span class="navbar-brand">MyPadelClub</span>
        <button class="navbar-toggler" type="button" @click="changeIsNavShow">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div :class="navClass">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <router-link class="nav-link" to="/">Home</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/rent">Reservas</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/workers">Listado de trabajadores</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/court">Menu Pistas</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/login" v-if="state.userLogued === false">Iniciar sesión</router-link>
                </li>
                 <li class="nav-item">
                    <router-link class="nav-link" to="/login" v-if="state.userLogued === true" @click="checkOut">Cerrar sesión</router-link>
                </li>
            </ul>
        </div>
    </nav>
</template>

<script>

import { useRoute } from 'vue-router';
import { reactive, computed } from "vue";
import { useStore } from "vuex";

export default {
 
  name: 'Header',
  setup() {

      const route = useRoute();
      const store = useStore();

      const state = reactive({

        userLogued: computed(() => store.getters["worker/userLogued"])

      });

      
        const checkOut = () => {
            localStorage.removeItem('token');
            localStorage.removeItem('admin');
            location.reload();
        }


    return { route,state,checkOut }
    

  }

}
</script>

<style>

</style>