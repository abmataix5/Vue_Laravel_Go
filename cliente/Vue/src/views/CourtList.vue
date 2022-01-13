<template>
    <div class="mt-2 ml-3">
        <div class="row ml-2">
            <h3>PISTAS DISPONIBLES</h3>
        </div>
       <div class="row">
            <div class="col p-3">
                <router-link class="btn btn-primary" to="/Court/add">Añadir Pista</router-link>
            </div>
        </div>
        <div class="card card-default card-borderless">
        <div class="card-body">
            <div class="row">
                <div class="col">
                    <ul class="list-group">
                        <CourtItem v-for="courtitem in courtlist" :key="courtitem.id" :courtitem="courtitem" />
                    </ul>
                </div>
            </div>
        </div>
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

</style>