<template>

    <li >
        <span :class="{ pointer:true}">
           Nombre del usuario : {{workeritem.username}}<br>
           Email: {{workeritem.email}}<br>
           Numero de telefono: {{workeritem.phone}}<br>
           Cargo del trabajador: {{workeritem.worker_type}}<br>
        </span>
        <div class="float-right">
            <span class="badge badge-secondary pointer ml-1" @click.stop="editTodo(workeritem.id)">Editar</span> 
            <span class="badge badge-secondary pointer ml-1" @click.stop="deleteW(workeritem.id)">Eliminar</span>
        </div>
    </li>

    
</template>

<script>

import Constant from "../Constant";
import { useStore } from "vuex";
/* import { useRouter } from "vue-router"; */

export default {
  props: {
    workeritem: Object,
  },
  setup() {

   const store = useStore();
/*    const router = useRouter(); */

   const deleteW = (id) => {
       console.log("delete");
            
       console.log(id);
       store.dispatch("worker/"+Constant.DELETE_WORKER, { id });
    };
    return { deleteW }; 
  },
};
</script>

<style></style>