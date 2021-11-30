<template>
    <nav class="navbar navbar-expand-md bg-dark navbar-dark">
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
                    <router-link class="nav-link" to="/rentings">Reservas</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/workers">Listado de trabajadores</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/courts">Crear nueva pista</router-link>
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