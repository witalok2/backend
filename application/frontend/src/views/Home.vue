<template>
  <div class="home">
    <div id="text-painel">TODAS AS ATIVIDADES</div>

    <q-dialog v-model="infoEdit">
      <q-card style="width: 700px; max-width: 80vw;">
        <q-card-section>
          <div class="text-h6">Alteração de Atividade</div>
        </q-card-section>
        <form id="formEdicao">
          <q-input outlined v-model="dadosEdit.titulo" label="Titulo" />
          <br />
          <!-- Input com o titulo da atividade -->
          <q-input outlined v-model="dadosEdit.subtitulo" label="SubTitulo" />
          <br />
          <!-- input com o subtitulo da atividade -->
          <q-input v-model="dadosEdit.descricao" filled type="textarea" label="Descrição" />
          <!-- Text area  -->
        </form>

        <q-card-actions align="right" class="bg-white text-teal">
          <q-btn @click="Update(atividade)" flat label="OK" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <q-card
      flat
      bordered
      class="my-card bg-grey-1"
      v-for="(atividade, index) of atividades"
      :key="index"
    >
      <q-card-section>
        <div class="row items-center no-wrap">
          <div class="col">
            <q-badge color="red">{{ index + 1 }}</q-badge>
            <div class="text-titulo">{{ atividade.titulo }}</div>
            <div class="text-subtitulo">{{ atividade.subtitulo }}</div>
            <div class="text-descricao">{{ atividade.descricao }}</div>
            <br />
            <div class="btns">
              <q-btn
                @click="infoEdit = true, dadosEdit = atividade"
                id="btnEditar"
                outline
                rounded
                color="blue"
                label="Editar"
              />
              <q-btn @click="Delete(atividade.id)" outline rounded color="red" label="Excluir" />
            </div>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<style>
</style>

<script>
//Imports
import Atividades from "../services/atividades";

export default {
  name: "PageHome",

  data() {
    return {
      atividade: {
        id: "",
        titulo: "",
        subtitulo: "",
        descricao: ""
      },
      atividades: [],
      errors: [],
      infoEdit: false,
      dadosEdit: {}
    };
  },

  mounted() {
    this.List();
  },

  computed: {},

  methods: {
    List() {
      Atividades.list().then(response => {
        this.atividades = response.data.atividades;
      });
    },

    Delete(id) {
      Atividades.apagar(id)
        .then(resposta => {
          this.deleteNotify();
          this.List();
        })
        .catch(e => {
          this.errors = e.response.data.errors;
        });
    },

    Update(dado) {
      Atividades.update(this.dadosEdit)
        .then(resposta => {
          this.dadosEdit = {};
          this.errors = {};
          this.List();
          this.updateNotify();
        })
        .catch(e => {
          this.errors = e.response.data.errors;
        });
    },

    // Notificação de Deletado Atividade
    deleteNotify() {
      this.$q.notify({
        message: "Atividade Deletada Com Sucesso.",
        color: "red"
      });
    },

    // Notificação de atualização de Atividade
    updateNotify() {
      this.$q.notify({
        message: "Atividade Alterada Com Sucesso.",
        color: "purple"
      });
    }

  }
};
</script>

<style scoped>
#home {
  background-color: rgb(107, 106, 106);
}

#text-painel {
  opacity: 0.5;
  font-size: 30px;
  text-align: center;
}

.my-card {
  width: 100%;
  max-width: 100%;
  margin-bottom: 10px;
}

.text-titulo {
  text-align: center;
  font-size: 35px;
  color: #028122;
}

.text-subtitulo {
  text-align: center;
  font-size: 20px;
}

.text-descricao {
  text-align: center;
  font-size: 12px;
}

.btns {
  text-align: right;
  color: #be0202;
  margin-right: 3px;
}

#btnEditar {
  margin-right: 5px;
}

#formEdicao {
  margin-left: 10px;
  margin-right: 10px;
}
</style>