<template>

    <li >
        <span :class="{ pointer:true}">
           ID del usuario : {{partneritem.id}}<br>
           Nombre del usuario : {{partneritem.name}}<br>
           Email: {{partneritem.email}}<br>
        </span>
        <div class="float-right">
            <span class="btn btn-success   pointer ml-1" @click.stop="editPartner(partneritem.id)">Editar</span> 
            <span class="btn btn-danger pointer ml-1" @click.stop="deletePartner(partneritem.id)">Eliminar</span>
        </div>
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

   const deletePartner = (id) => {
       console.log("delete");
            
       console.log(id);
       store.dispatch(Constant.DELETE_PARTNER, { id });
    };
    const editPartner = (id) => {
            console.log("EDIT PARTNER");
            // console.log(id)
            // store.dispatch(Constant.UPDATE_COURT, { courtitem: { ...props.courtitem } });
            router.push({ name: 'updatePartner', params: { id } })
        }
    return { deletePartner, editPartner }; 
  },
};
</script>

<style></style>