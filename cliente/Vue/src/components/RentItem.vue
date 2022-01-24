<template>
<div class="row col-3 d-flex justify-content-around">
        <figure>
                <img src="../assets/court.png" class="img-fluid ">
            <figcaption>   
                <div>ID: {{rentitem.id}}</div>
                <div>NAME: {{rentitem.name}}</div>
                <div>TURNO: {{rentitem.schedule}}</div>
            </figcaption>
            <div class="float-right">
                <span class="btn btn-success   pointer ml-1 mt-2" @click.stop="addUser(useritem.id)">Añadir Jugador</span><br>
                <span class="btn btn-info   pointer ml-1 mt-1" @click.stop="editRent(rentitem.id)">Editar</span>
                <span class="btn btn-danger pointer ml-1 mt-1" @click.stop="deleteCourt(courtitem.id)">Borrar</span>
            </div>
        </figure>     
</div>
</template>

<style>
/* .row{
    display: flex;
    flex-direction: row;
} */
figure{
    width: 350px;
    height: 350px;
    border: 1px solid black;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    justify-items: center;
   
}
figure caption div{
    padding-left: 15px;
}
img{
    width: 120px;
    height: 100px;
}
</style>

<script>
import Constant from '../Constant';
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    props: {
        rentitem: Object
    },
    setup() {
        const store = useStore();
        const router = useRouter();

        const deleteRent = (id) => {
            console.log(id);
            store.dispatch(Constant.DELETE_RENT, { id });
        }
        const editRent = (id) => {
            console.log("EDIT RENT");
            // store.dispatch(Constant.UPDATE_COURT, { rentitem: { ...props.item } });
            router.push({ name: 'updateRent', params: { id } })
        }
         return { deleteRent, editRent}
    }
}
</script>

