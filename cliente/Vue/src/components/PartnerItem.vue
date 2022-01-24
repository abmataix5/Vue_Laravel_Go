<template>

<div class="col-3 partner d-flex justify-content-around">
    <div class="card p-2 m-2">
      <div class="img-card d-flex justify-content-center">
       <img class="h-75" v-bind:src="partneritem.image" />
       </div>
      <div class="card‐body">
        <h3 class="card‐title">
          <strong>Num Socio : </strong>{{partneritem.id}}<br>
          <strong>Nombre : </strong>{{partneritem.name}}<br>
          <strong>Apellidos: </strong> {{partneritem.lastname}}
        </h3>
        <div class="card‐text d-flex justify-content-center">
          <span class="btn btn-info   pointer ml-1" @click.stop="detailPartner(partneritem.id)">Ver Detalles</span> 
          <span class="btn btn-success   pointer ml-1" @click.stop="editPartner(partneritem.id)">Editar</span> 
          <span class="btn btn-danger pointer ml-1" @click.stop="deletePartner(partneritem.id)">Eliminar</span>
        </div>
      </div>
    </div>
</div>

    
</template>

<script>

import Constant from "../Constant";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

export default {
  props: {
    partneritem: Object,
  },
  setup() {

   const store = useStore();
   const router = useRouter();

   const detailPartner = (id) => {
 
      router.push({ name: 'detailPartner', params: { id } })
    };

   const deletePartner = (id) => {

       store.dispatch(Constant.DELETE_PARTNER, { id });
    };
    const editPartner = (id) => {
           
          
            router.push({ name: 'updatePartner', params: { id } })
      }
    return { deletePartner, editPartner, detailPartner }; 
  },
};
</script>

<style>

</style>