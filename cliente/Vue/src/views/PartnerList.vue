<template>
    <div style="mt-2">
       <div class="row">
            <div class="col p-3">
                <router-link class="btn btn-primary" to="/partner/add">Añadir Socio</router-link>
            </div>
        </div>
        <div class="card card-default card-borderless">
        <div class="card-body">
            <div class="row">
                <div class="col">
                    <ul class="list-group">
                        <PartnerItem v-for="partneritem in partnerlist" :key="partneritem.id" :partneritem="partneritem" />
                    </ul>
                </div>
            </div>
        </div>
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
            store.dispatch(Constant.INITIALIZE_PARTNERITEM);
            router.push({ name:"addPartner" });
        }

        return { partnerlist, goAddPartner }
    }
}
</script>

<style>

</style>