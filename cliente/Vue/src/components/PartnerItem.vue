<template>

    <li >
        <span :class="{ pointer:true}">
           ID del usuario : {{partneritem.id}}<br>
           Nombre del usuario : {{partneritem.name}}<br>
           Email: {{partneritem.email}}<br>
        </span>
        <div class="float-right">
            <span class="btn btn-success   pointer ml-1" @click.stop="editPartner(partneritem.ID)">Editar</span> 
            <span class="btn btn-danger pointer ml-1" @click.stop="deletePartner(partneritem.ID)">Eliminar</span>
        </div>
    </li>

    
</template>

<script>

import Constant from "../Constant";
import { useStore } from "vuex";
/* import { useRouter } from "vue-router"; */

export default {
  props: {
    partneritem: Object,
  },
  setup() {

   const store = useStore();
/*    const router = useRouter(); */

   const deletePartner = (id) => {
       console.log("delete");
            
       console.log(id);
       store.dispatch("partner/"+Constant.DELETE_PARTNER, { id });
    };
    return { deletePartner }; 
  },
};
</script>

<style></style>