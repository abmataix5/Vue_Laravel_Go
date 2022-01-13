import Constant from "../Constant";
import { reactive } from "vue";
import { useStore } from "vuex";
import { useRouter, useRoute } from "vue-router";

export default function useUpdateWorkers() {
  const store = useStore();
  const router = useRouter();
  // var showPassword = reactive(true);
  const currentRoute = useRoute();

  const workeritem = store.state.worker.workerlist.find(
    (item) => item.id === currentRoute.params.id
  );

  const updateWorker = () => {
    router.push({ name: "workerList" });
    store.dispatch("worker/" + Constant.UPDATE_WORKER, {
      workeritem: state.workeritemlocal,
    });
  };

  

  const state = reactive({
    workeritemlocal: {
      name: "",
      email: "",
      phone: "",
      address: "",
      worker_type: "",
    },
    workeritemlocal2: { ...workeritem },

    showPassword: true,
  });




  return {  state, updateWorker };
}