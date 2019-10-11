<template>
  <div class="home">
    <div id="text-painel">TODAS AS ATIVIDADES</div>
    <q-card flat bordered class="my-card bg-grey-1" v-for="atividade of atividades" :key="atividade.id">
      <q-card-section>
        <div class="row items-center no-wrap">
          <div class="col">
            <q-badge color="red">{{ atividade.id }}</q-badge>
            <div class="text-titulo">{{ atividade.titulo }}</div>
            <div class="text-subtitulo">{{ atividade.subtitulo }}</div>
            <div class="text-descricao">{{ atividade.descricao }}</div>
            <div class="text-situacao"><q-btn @click="deletar(atividade.id)"  outline rounded color="red" label="Excluir"/></div>
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
  name: 'PageHome',

  data() {
    return {
      atividade:{
        id: '',
        titulo: '',
        subtitulo: '',
        descricao: '',
        situacao:''
      },
      atividades: [],
      errors: []
      
      
    }
  },

  mounted() {
    this.listar();
  },

  computed: {
    
  },

  methods: {
    listar() {
      Atividades.listar().then(response => {
        this.atividades = response.data.atividades;
      });
    },

    deletar(id) {
      Atividades.apagar(id).then(resposta => {
          this.deleteNotify()
          this.listar()
        }).catch(e => {
          this.errors = e.response.data.errors
        })
    },

    // Notificação de Deletado a Atividade
    deleteNotify() {
      this.$q.notify({
        message: 'Atividade Deletada Com Sucesso.',
        color: 'red'
      })  
    }
  }
}
</script>

<style scoped>
#home {
  background-color: rgb(107, 106, 106);
}

#text-painel {
  opacity: 0.5;
  font-size: 30px;
}

.my-card {
  width: 100%;
  max-width: 100%;
  margin-bottom: 10px;
}

.text-titulo {
  text-align: center;
  font-size: 40px;
  color: #028122;
}

.text-subtitulo {
  text-align: center;
  font-size: 20px;
}

.text-descricao {
  text-align: center;
  font-size: 10px;
}

.text-situacao {
  text-align: right;
  color: #be0202;
}
</style>