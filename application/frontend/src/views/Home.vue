<template>
  <div class="home">
    <div id="text-painel">
      Painel de Atividades
    </div>

    <div class="row">
      <div class="col" id="coluna-pendentes">
        <h6 id="text-colunas"> Pendente </h6>
            <!-- Cards de atividades -->
            <q-card class="my-card bg-purple text-white">
              <q-card-section>
                <div class="text-h6">{{  titulo }}</div>
                <q-separator dark inset/>
                <div class="descricao">{{ subTitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ msg }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
              </q-card-actions>
            </q-card><br>

            <!-- Cards de atividades -->
            <q-card class="my-card bg-purple text-white">
              <q-card-section>
                <div class="text-h6">{{  titulo }}</div>
                <q-separator dark inset/>
                <div class="descricao">{{ subTitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ msg }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
              </q-card-actions>
            </q-card><br>
      </div>

      <div class="col" id="coluna-andamento">
        <h6 id="text-colunas"> Andamento </h6>
            <!-- Cards de atividades -->
            <q-card class="my-card bg-secondary text-white">
              <q-card-section>
                <div class="text-h6">{{  titulo }}</div>
                <q-separator dark inset/>
                <div class="descricao">{{ subTitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ msg }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
              </q-card-actions>
            </q-card><br>
      </div>

      <div class="col" id="coluna-finalizada">
        <h6 id="text-colunas"> Finalizada </h6>
            <!-- Cards de atividades -->
            <q-card class="my-card bg-brown-5 text-white">
              <q-card-section>
                <div class="text-h6">{{  titulo }}</div>
                <q-separator dark inset/>
                <div class="descricao">{{ subTitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ msg }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
              </q-card-actions>
            </q-card><br>

            <!-- Cards de atividades -->
            <q-card class="my-card bg-brown-5 text-white">
              <q-card-section>
                <div class="text-h6">{{  titulo }}</div>
                <q-separator dark inset/>
                <div class="descricao">{{ subTitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ msg }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
              </q-card-actions>
            </q-card><br>

            <!-- Cards de atividades -->
            <q-card class="my-card bg-brown-5 text-white">
              <q-card-section>
                <div class="text-h6">{{  titulo }}</div>
                <q-separator dark inset/>
                <div class="descricao">{{ subTitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ msg }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
              </q-card-actions>
            </q-card><br>                    
      </div>
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
 
  color: #2c3e50;

}

#text-painel  {
  opacity: 0.2;
  font-size: 80px;
  text-align: center;
}

#text-colunas  {
   text-align: center;
   align-content: center;
}

#coluna-pendentes {
  border: 1px;
  border-style: dotted;
  height: 100%;
  width: 100%;
  margin-left: 10px;
  margin-right: 10px;
}

#coluna-andamento {
  border: 1px;
  border-style: dotted;
  height: 100%;
  width: 100%;
  margin-left: 10px;
  margin-right: 10px;
}

#coluna-finalizada {
  border: 1px;
  border-style: dotted;
  height: 100%;
  width: 100%;
  margin-left: 10px;
  margin-right: 10px;
  opacity: 0.5;
}

/* Cards de Atividades */
.my-card {
  width: 100%;
  max-width: 350px;
  margin-left: auto;
  margin-right: auto;
  text-align: left;
}
</style>