<template>
  <div id="app">
    <q-layout view="hHh lpR fFf">
      <q-header elevated class="bg-primary text-white">
        <q-toolbar>
          <q-toolbar-title>
              <div id="titulo">Todo List</div>
          </q-toolbar-title>
          <q-btn push color="white" text-color="primary" label="+" @click="createAt = true"/>
        </q-toolbar>
      </q-header>

      <q-dialog v-model="createAt" full-height>
        <q-card class="column full-height" style="width: 800px" >
          <q-card-section>
            <div class="text-h6">Informações da Atividade</div>
          </q-card-section>
            <div id="inputTitulo">
              <q-input v-model="text" label="Titulo"/>
              <q-input v-model="text" label="SubTitulo"/>
              <q-input square outlined v-model="text" label="Square outlined" />
            </div>
          <q-card-section class="col">
        
          </q-card-section>

          <q-card-actions align="right" class="bg-white text-teal">
            <q-btn flat label="Criar" v-close-popup />
            <q-btn flat label="Fechar" v-close-popup />
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

export default {
  name: 'app',
  data () {
    return {
      nomeUser: 'Witalo Matheus ',
      createAt: false,
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