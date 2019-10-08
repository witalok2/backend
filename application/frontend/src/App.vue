<template>
  <div id="app">
    <q-layout view="hHh lpR fFf">
      <q-header elevated class="bg-primary text-white">
        <q-toolbar>
          <q-toolbar-title>
              <div id="titulo">Todo List</div>
          </q-toolbar-title>
          <q-btn push color="white" text-color="primary" label="Criar" @click="createAt = true"/>
        </q-toolbar>
      </q-header>

      <q-dialog v-model="createAt" full-height>
        <q-card class="column full-height" style="width: 800px" >
          <q-card-section>
            <div class="text-h6">Criação de Atividade</div>
          </q-card-section>

          <!-- formulario com os inputs de criação de atividades -->
          <form>
            <div id="inputTitulo">
                <q-input v-model="atividade.titulo" label="Titulo" />
                <q-input v-model="atividade.subtitulo" label="Subtitulo"/>
                <q-input v-model="atividade.descricao" filled type="textarea" placeholder="Descrição da Atividade"/>
            </div>
          </form>

          <q-card-actions align="right" class="bg-white text-teal">
            <q-btn v-on:click="salvar" flat label="Salvar" v-close-popup />
            <q-btn  flat label="Fechar" v-close-popup />
          </q-card-actions>
        </q-card>

      </q-dialog>

      <q-page-container>
        <router-view />
      </q-page-container>

    </q-layout>
  </div>
</template>

<script>
//Imports
import Atividades from './services/atividades'

export default {
  name: 'app',
  data () {
    return {
      createAt: false,

      // Dados dos inputs
      atividade:{
        titulo: '',
        subtitulo: '',
        descricao: '',
        situacao: 'Pendente'
      },
      atividades: [],
      errors: []

    }
  },

  methods:{
    //Função insert de Atividade
    salvar(){
        Atividades.salvar(this.atividade).then(resposta => { 
          this.atividade = {}
          this.errors = {}
          this.createNotify()
        }).catch(e => {
          this.errors = e.response.data.errors
        })
      }, 

    // Notificação de Sucesso na criação de Atividade
    createNotify() {
      this.$q.notify({
        message: 'Atividade Criada Com Sucesso.',
        color: 'green'
      })  
    },

    // Notificação de Erro na criação de Atividade
    erroNotify() {
      this.$q.notify({
        message: 'Erro na Criação da Atividade.',
        color: 'red'
      })  
    }
  }
}

</script>


<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  //text-align: center;
  color: #2c3e50;
  background-color: #ccc;
}

.q-toolbar{
  background-color: white;
}

#inputTitulo {
  margin-left: 20px;
  margin-right: 40px;
}

</style>