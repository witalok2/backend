import Vue from "vue";
import Vuex from "vuex";
import { api } from "@/services.js";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    listar_tarefas: null
  },
  mutations: {
    UPDATE_TAREFAS(state, payload) {
      state.listar_tarefas = payload;
    },
    ADD_TAREFAS(state, payload) {
      state.listar_tarefas.unshit(payload);
    },

  },
  actions: {
    getTarefasUp(context) {
      api.get("/tarefas")
          .then(response => {
            context.commit("UPDATE_TAREFAS", response.data);
          });
    },
  }
})
