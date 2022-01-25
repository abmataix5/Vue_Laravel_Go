<template>
    <div style="mt-2">
       <div class="row">
            <div class="col p-3">
                <h3>Ficha de {{partnerdetail.name}}</h3>
            </div>
        </div>
        <div class="card card-default card-borderless">
        <div class="card-body">
            <div class="row">
                 <div class="col-1">
                    <div class="row">
                        <img class="h-75" v-bind:src="partnerdetail.image" /> 
                    </div>
                    <div class="row">
                        <!-- <button type="button" class="btn btn-primary m-1" @click="add_img">Cambiar Avatar</button> -->
                        <!-- <button type="button" class="btn btn-primary m-1" @click="delete_img">Eliminar</button> -->
                    </div>
                </div>
                <div class="col">
                    <ul class="list-group">
                        Num Socio: {{partnerdetail.id}} <br>
                        Nombre: {{partnerdetail.name}}<br>
                        Apellidos: {{partnerdetail.lastname}}<br>
                        Tlf Contacto: {{partnerdetail.phone}}<br>
                        Email: {{partnerdetail.email}}<br>
                        Posición: {{partnerdetail.position}}<br>
                        <!-- rellenar datos -->
                    </ul>
                   
                </div>
            </div>
        </div>
        </div>
        <div class="form-group">
            <button type="button" class="btn btn-primary m-1" @click="cancel">Volver</button>
        </div>
    </div>
</template>

<script>
import Constant from '../Constant';
// import PartnerItem from '../components/PartnerItem';
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter , useRoute} from 'vue-router';

export default {

    setup() {
        console.log("dentro setup!");
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();
        // console.log(currentRoute.params.id);
        // console.log(store.state.partner);
        store.dispatch(Constant.DETAIL_PARTNER, currentRoute.params.id);

        const partnerdetail = computed(() => store.getters["partner/detailPartners"]) ;

    const cancel = () => {
        router.push({ name:"partnerList"});
    }       
        return {partnerdetail, cancel}
    }
}
</script>

<style>

</style>