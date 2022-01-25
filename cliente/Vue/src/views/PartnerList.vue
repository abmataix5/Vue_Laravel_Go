<template>
    <div style="mt-2">
        <div class="row title">
            <div class="title_panelAdmin ml-4">
                <h1> Panel Administración Socios </h1>
            </div>
        </div>
       <div class="row">
            <div class="col p-3">
                <h4>Acciones Disponibles:</h4>
                <router-link class="btn btn-success" to="/partner/add">Añadir Nuevo Socio</router-link>
            </div>
        </div>

        <hr>
        <div class="row">
            <PartnerItem v-for="partneritem in partnerlist" :key="partneritem.id" :partneritem="partneritem" />
        </div>
    </div>
</template>

<script>
import Constant from '../Constant';
import PartnerItem from '../components/PartnerItem';
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    components : { PartnerItem },
    setup() {
        const store = useStore();
        const router = useRouter();
        console.log("dentro setup!");
        store.dispatch(Constant.INITIALIZE_PARTNERITEM)
        
        const partnerlist = computed(() => store.getters["partner/getPartners"])        
        // console.log(partnerlist)
        const goAddPartner = () => {
            console.log("DENTRO ADD PARTNER");
            // store.dispatch(Constant.INITIALIZE_PARTNERITEM);
            router.push({ name:"addPartner" });
        }

        return { partnerlist, goAddPartner }
    }
}
</script>

<style>
.title{
    background-color: rgba(178, 178, 179, 0.788);
}
</style>