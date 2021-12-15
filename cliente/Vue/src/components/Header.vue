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
                    <router-link class="nav-link" to="/login">Iniciar sesión</router-link>
                </li>
            </ul>
        </div>
    </nav>
</template>

<script>

import { useRoute } from 'vue-router';

export default {
 
  name: 'Header',
  setup() {

    const route = useRoute();

    return {
      route
    }
  }

}
</script>

<style>

</style>