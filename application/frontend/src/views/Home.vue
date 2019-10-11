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
            <div class="text-situacao"><q-btn @click="deletar(atividade.id)"  outline rounded color="red" label="Excluir" /></div>
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

    deletar(id){
      Atividades.apagar(id).then(resposta => {
          alert(id)
          this.listar()
        }).catch(e => {
          this.errors = e.response.data.errors
        })
    }
  }
}
</script>

<style scoped>
#home {
  background-color: rgb(107, 106, 106);
}

.my-card {
  width: 100%;
  max-width: 100%;
  margin-bottom: 10px;
}

.text-titulo {
  font-size: 40px;
  color: #028122;
}

.text-subtitulo {
  font-size: 20px;
}

.text-descricao {
  font-size: 10px;
}

.text-situacao {
  text-align: right;
  color: #be0202;
}
</style>