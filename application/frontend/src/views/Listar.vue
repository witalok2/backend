<template>
  <div class="home">
    <div id="text-painel">
      TODAS AS ATIVIDADES
    </div>

  </div>
</template>

<script>
//Imports
import axios from 'axios'

export default {
  name:'listar',
  data(){
    return {
      atividades: null
      }
  },
  mounted(){
    axios.get('http://localhost:1323/atividade/listar').then(response => (this.atividades = response))}
}

</script>

<style>
.home {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#text-painel  {
  opacity: 0.2;
  font-size: 80px;
  text-align: center;
  margin-bottom: 30px;
}

#tabela {
  padding-left: 30px;
  padding: 30px;
}

</style>