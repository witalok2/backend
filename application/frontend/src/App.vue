<template>
  <div id="app">
    <q-layout view="hHh lpR fFf">
      <q-header elevated class="bg-primary text-white">
        <q-toolbar>
          <q-toolbar-title>
              <div id="titulo">Todo List</div>
          </q-toolbar-title>

          <q-btn to="/Listar" label="Atividades" outline color="purple" />

        </q-toolbar>
      </q-header>

      <q-page-container>
        <router-view />
      </q-page-container>

    </q-layout>
  </div>
</template>

<script>
//Imports

export default {
  name: 'app',
  data () {
    return {
      nomeUser: 'Witalo Matheus '
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



</style>