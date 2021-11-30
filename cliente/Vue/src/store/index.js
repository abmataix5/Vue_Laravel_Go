import { createStore } from 'vuex';
import { courtStore } from "./court";
import { workerStore } from './worker';

export const store = createStore({

modules:{
    court: courtStore,
    worker: workerStore
}

});