<template>

  <div class="container">

      <div style="mt-2">
        <div class="row">
            <div class="col p-3">
                <router-link class="btn btn-primary" to="/addworker">Añadir nuevo trabajador</router-link>
            </div>
        </div>
        <div class="card card-default card-borderless">
        <div class="card-body">
            <div class="row">
                <div class="col">
                    <ul class="list-group">
                                <WorkerItem v-for="workeritem in state.workerlist" :key="workeritem.id"  :workeritem="workeritem"/>
                    </ul>
                </div>
            </div>
        </div>
        </div>
    </div>


  </div>

</template>

<script>

import Constant from "../Constant";
import WorkerItem from "../components/WorkerItem.vue";
import { reactive, computed } from "vue";
import { useStore } from "vuex";

export default {
    
    name: 'workers',
    components:{
        WorkerItem
    },
    setup() {

      const store = useStore();
      const state = reactive({

        workerlist: computed(() => store.getters["worker/getWorkers"]),
      });

      if (!state.workerlist) {
    
        store.dispatch("worker/" + Constant.INITIALIZE_WORKERITEM);
      }
      return { state };
  },
    
}
</script>

<style>

</style>
