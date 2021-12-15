<template>
<div class="row">
        <!-- <p>hola</p> -->
         <!-- <span :class="{ pointer:true, 'user-done':useritem.done }" :title="'Explanation : ' + useritem.desc">  -->
        <span> 
            ID: {{courtitem.id}}
            NAME: {{courtitem.name}}
            TURNO: {{courtitem.schedule}}
         </span> 
        <div class="float-right">
            <span class="badge badge-secondary pointer ml-1" @click.stop="editCourt(courtitem.id)">Edit</span> 
            <span class="badge badge-secondary pointer ml-1" @click.stop="deleteCourt(courtitem.id)">Delete</span>
        </div>
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
            console.log(id);
            store.dispatch(Constant.DELETE_COURT, { id });
        }
        const editCourt = (id) => {
            console.log("EDIT COURT");
            // store.dispatch(Constant.UPDATE_COURT, { courtitem: { ...props.courtitem } });
            router.push({ name: 'updateCourt', params: { id } })
        }
         return { deleteCourt, editCourt}
    }
}
</script>

<style>

</style>