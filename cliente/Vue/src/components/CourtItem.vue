<template>

<div class="row col-3 d-flex justify-content-around ml-2">
        <figure>
                <img src="../assets/court.png" class="img-fluid ">
            <figcaption>   
                <div>ID: {{courtitem.id}}</div>
                <div>NAME: {{courtitem.name}}</div>
                <div>TURNO: {{courtitem.schedule}}</div>
            </figcaption>
            <div class="float-right">
                <span class="btn btn-success   pointer ml-1 mt-2" @click.stop="editCourt(courtitem.id)">Editar Pista</span> 
                <span class="btn btn-danger pointer ml-1 mt-1" @click.stop="deleteCourt(courtitem.id)">Eliminar Pista</span>
            </div>
        </figure>     
</div>
</template>

<script>
import Constant from '../Constant';
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    props: {
        courtitem: Object
    },
    setup() {
        const store = useStore();
        const router = useRouter();

        const deleteCourt = (id) => {
           
            store.dispatch(Constant.DELETE_COURT, { id });
        }
        const editCourt = (id) => {
        
            router.push({ name: 'updateCourt', params: { id } })
        }
         return { deleteCourt, editCourt}
    }
}
</script>

<style>

</style>