<template>

    <li >
        <span :class="{ pointer:true}">
           <strong>Num Socio : </strong>{{partneritem.id}}
           <strong>Nombre : </strong>{{partneritem.name}}
           <strong>Apellidos: </strong> {{partneritem.lastname}}
        </span>
        <div class="float-right">
            <span class="btn btn-info   pointer ml-1" @click.stop="detailPartner(partneritem.id)">Ver Detalles</span> 
            <span class="btn btn-success   pointer ml-1" @click.stop="editPartner(partneritem.id)">Editar</span> 
            <span class="btn btn-danger pointer ml-1" @click.stop="deletePartner(partneritem.id)">Eliminar</span>
        </div>
         <hr>
  
    </li>

    
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

<style></style>