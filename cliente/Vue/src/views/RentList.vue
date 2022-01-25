<template>
    <div style="mt-2">
         <div class="row title">
            <div class="title_panelAdmin ml-4">
                <h1> Panel Administración Alquileres </h1>
            </div>
        </div>
        <div class="row">
            <div class="col p-3">
                 <h4>Acciones Disponibles:</h4>
                <router-link class="btn btn-success" to="/Rent/add">Unirme a la partida</router-link>
            </div>
        </div>
        <div class="row">
                        <RentItem v-for="rentitem in rentlist" :key="rentitem.id" :rentitem="rentitem" />
        </div>
    </div>
</template>

<script>
import Constant from '../Constant';
import RentItem from '../components/RentItem';
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    components : { RentItem },
    setup() {
        const store = useStore();
        const router = useRouter();

        store.dispatch(Constant.INITIALIZE_RENTITEM)
        
        const rentlist = computed(() => store.getters["rent/getRent"])        
        
        const goAddRent = () => {
            console.log("DENTRO ADD RENT");
            store.dispatch(Constant.INITIALIZE_RENTITEM);
            router.push({ name:"addRent" });
        }

        return { rentlist, goAddRent }
    }
}
</script>

<style>
.title{
    background-color: rgba(178, 178, 179, 0.788);
}
</style>