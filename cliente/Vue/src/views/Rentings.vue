<template>
    
    <div>
         AQUI PODRÁS ALQUILAR PISTAS DISPONIBLES! <CourtList/>
    </div>

</template>

<script>
import CourtList from '../components/CourtList.vue'

export default {
    
    name: 'rentings',
    components:{
        CourtList
    }
    
}
</script>

<style>

</style>
