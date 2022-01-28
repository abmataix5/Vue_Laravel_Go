<template>
    <nav class="navbar navbar-expand-lg bg-primary navbar-dark text-white py-4">
        
     <router-link href="#" class="navbar-brand title" to="/" >MyPadelClub</router-link>

        <div :class="navClass">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <router-link class="nav-link" to="/">Home</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/rent">Reservas</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/partner">Listado de Socios</router-link>
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
                    <router-link class="nav-link" to="/loginLaravel" v-if="state.userLogued === false">Acceso STAFF </router-link>
                </li>
                 <li class="nav-item">
                    <router-link class="nav-link" to="/login" v-if="state.userLogued === true" @click="checkOut">Cerrar sesión</router-link>
                </li>
                 <li class="nav-item text-white" v-if="state.userAdmin === true">
                    Cuenta Administradora
                </li>
                 <ul class="nav-item text-white" v-if="state.userAdmin === false">
                    Cuenta Trabajador {{state.userAdmin}}
                </ul>
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

        userLogued: computed(() => store.getters["worker/userLogued"]),
        userAdmin: computed(() => store.getters["worker/isAdmin"]),
        username : localStorage.getItem('username')
      });

      
        const checkOut = () => {
            localStorage.removeItem('token');
            localStorage.removeItem('admin');
             localStorage.removeItem('username');
            location.reload();
        }


    return { route,state,checkOut }
    

  }

}
</script>

<style>

</style>