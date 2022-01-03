import { createStore } from 'vuex';
import { courtStore } from "./court";
import { workerStore } from './worker';
import { rentStore } from './rent';
import { partnerStore } from './partner';
export const store = createStore({

modules:{
    court: courtStore,
    worker: workerStore,
    rent: rentStore,
    partner: partnerStore
}

});