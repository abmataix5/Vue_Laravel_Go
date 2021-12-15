<template>
    <div style="mt-2">
       <div class="row">
            <div class="col p-3">
                <router-link class="btn btn-primary" to="/Rent/add">Unirme a la partida</router-link>
            </div>
        </div>
        <div class="row">
            <div>Dias</div>

        </div>
        <div class="card card-default card-borderless">
        <div class="card-body">
            <div class="row">
                <div class="col">
                    <div class="list-group">
                        <RentItem v-for="rentitem in rentlist" :key="rentitem.id" :rentitem="rentitem" />
                    </div>
                </div>
            </div>
        </div>
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
.listgroup{
    /* display: flex;
    flex-direction: row; */
}
</style>