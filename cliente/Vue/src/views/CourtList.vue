<template>
    <div class="mt">
        <div class="row title">
            <div class="title_panelAdmin ml-4">
                <h1> Panel Administración Pistas </h1>
            </div>
        </div>
       <div class="row">
            <div class="col p-3">
                <h4>Acciones Disponibles:</h4>
                <router-link class="btn btn-success" to="/Court/add">Añadir Pista</router-link>
            </div>
        </div>
        <hr>
        <div class="row">
                <CourtItem v-for="courtitem in courtlist" :key="courtitem.id" :courtitem="courtitem" />
        </div>
    </div>
</template>

<script>
import Constant from '../Constant';
import CourtItem from '../components/CourtItem';
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    components : { CourtItem },
    setup() {
        const store = useStore();
        const router = useRouter();

        store.dispatch(Constant.INITIALIZE_COURTITEM)
        
        const courtlist = computed(() => store.getters["court/getCourts"])        
        
        const goAddCourt = () => {
            console.log("DENTRO ADD COURT");
            store.dispatch(Constant.INITIALIZE_COURTITEM);
            router.push({ name:"addCourt" });
        }

        return { courtlist, goAddCourt }
    }
}
</script>

<style>
.title{
background-color: rgba(178, 178, 179, 0.788);
}
</style>