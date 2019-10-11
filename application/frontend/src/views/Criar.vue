<template>
<div>
  <div id="text-painel">CRIAR ATIVIDADES</div> <!-- Titulo -->
  <form>
  <q-page padding>
    <q-input v-model="atividade.titulo" label="Titulo" /><br> <!-- Input com o titulo da atividade -->
    <q-input v-model="atividade.subtitulo" label="SubTitulo" /><br> <!-- input com o subtitulo da atividade -->
    <q-input v-model="atividade.descricao" filled type="textarea" label="Descrição"/> <!-- Text area  -->
    <q-btn color="white" text-color="black" label="Criar"  v-on:click="salvar" />
  </q-page>
  </form>
</div>
</template>

<script>
//imports 
import Atividades from '../services/atividades'

export default {
  name: 'PageCriar',
  data() {
    return {
      atividade:{
        titulo: '',
        subtitulo: '',
        descricao: '',
        situacao:'Pendente',
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
    }
  }
}
</script>

<style scoped>

#text-painel {
  opacity: 0.5;
  font-size: 30px;
  text-align: center;
}
</style>