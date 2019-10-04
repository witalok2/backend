<template>
  <div class="home">
    <div id="text-painel">
      Painel de Atividades
    </div>
  </div>
</template>

<script>

export default {
  name: 'home',
  data () {
    return {
      titulo:     'API - Pagamento',
      subTitulo:  'Meio de Pagamento',
      msg:      'Inserir Botão de pagamento na plataforma'

    }
  }
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
}

</style>