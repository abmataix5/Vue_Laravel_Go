export const courtStore = {
    namespaced: true,
    state: {
      courts: []
    },
    mutations: {
      SETCOURTS(state, value) {
        state.courts = value;
      }
    },
    actions: {
      setCourts(state, value) {
        state.commit("SETCOURTS", value);
      }
    },
    getters: {
      getCourts(state) {
        return state.cars;
      }
    }
  }


