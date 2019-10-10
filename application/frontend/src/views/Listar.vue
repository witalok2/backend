<template>
  <div class="home">
    <div id="text-painel">TODAS AS ATIVIDADES</div>

    <!-- Dialog Atividade pendente -->
    <q-dialog v-model="infoAtPen">
      <q-card style="width: 700px; max-width: 80vw;">
        <q-card-section>
          <div class="text-h6">Informações da Atividade</div>
        </q-card-section>

        <q-card-section class="col">
          <q-badge color="red">{{all.id}}</q-badge>
          <p></p>
          <div id="txtDialoghTitulo">{{all.titulo}}</div>
          <div id="txtDialoghSubTitulo">{{all.subtitulo}}</div>
          <div id="txtDialoghDescricao">{{all.descricao}}</div>
        </q-card-section>

        <q-card-actions align="right" class="bg-white text-teal">
          <q-btn flat label="Iniciar" v-close-popup @click="editar" />
          <q-btn flat label="Fechar" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Dialog Atividade Em Andamento -->
    <q-dialog v-model="infoAtEnd">
      <q-card style="width: 700px; max-width: 80vw;">
        <q-card-section>
          <div class="text-h6">Informações da Atividade</div>
        </q-card-section>

        <q-card-section class="col">
          <q-badge color="red">{{all.id}}</q-badge>
          <p></p>
          <div id="txtDialoghTitulo">{{all.titulo}}</div>
          <div id="txtDialoghSubTitulo">{{all.subtitulo}}</div>
          <div id="txtDialoghDescricao">{{all.descricao}}</div>
        </q-card-section>

        <q-card-actions align="right" class="bg-white text-teal">
          <q-btn flat label="Finalizar" v-close-popup />
          <q-btn flat label="Fechar" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <div class="row">
      <div class="col" id="coluna-pendentes">
        <h6 id="text-colunas">Pendente</h6>
        <!-- Cards de atividades -->

        <q-card
          class="my-card bg-secondary text-white"
          v-for="(lista, index) in filteListaPendente"
          :key="index"
        >
          <q-btn @click="infoAtPen = true, all= lista">
            <q-card-section>
              <q-badge color="red">#{{ lista.id }}</q-badge>
              <div class="text-h6">{{ lista.titulo }}</div>
              <q-separator />
              <div class="descricao">{{ lista.subtitulo }}</div>
            </q-card-section>
          </q-btn>
        </q-card>
      </div>

      <div class="col" id="coluna-andamento">
        <h6 id="text-colunas">Em Andamento</h6>
        <!-- Cards de atividades -->
        <q-card
          class="my-card bg-purple text-white"
          v-for="(lista, index) in filteListaAndamento"
          :key="index"
        >
          <q-btn @click="infoAtEnd = true, all= lista">
            <q-card-section>
              <q-badge color="red">#{{ lista.id }}</q-badge>
              <div class="text-h6">{{ lista.titulo }}</div>
              <q-separator />
              <div class="descricao">{{ lista.subtitulo }}</div>
            </q-card-section>
          </q-btn>
        </q-card>
      </div>

      <div class="col" id="coluna-finalizada">
        <h6 id="text-colunas">Finalizada</h6>
        <!-- Cards de atividades -->
        <q-card
          class="my-card bg-brown-5 text-white"
          v-for="(lista, index) in filteListaFinalizada"
          :key="index"
        >
          <q-btn @click="infoAt = true, all= lista">
            <q-card-section>
              <q-badge color="red">#{{ lista.id }}</q-badge>
              <div class="text-h6">{{ lista.titulo }}</div>
              <q-separator />
              <div class="descricao">{{ lista.subtitulo }}</div>
            </q-card-section>
          </q-btn>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script>
//Imports
import Atividades from "../services/atividades";

export default {
  name: "listar",
  data() {
    return {
      listaAlls: [],
      infoAtPen: false,
      infoAtEnd: false,
      searchP: "Pendente", searchA: "Em Andamento", searchF: "Finalizada",
      all: {}
    };
  },

  mounted() {
    this.listar();
  },

  computed: {
    // Refatorar esse filtro que esta triplicado....
    filteListaPendente: function() {
      return this.listaAlls.filter(lista => {
        return lista.situacao.match(this.searchP);
      });
    },

    filteListaAndamento: function() {
      return this.listaAlls.filter(lista => {
        return lista.situacao.match(this.searchA);
      });
    },

    filteListaFinalizada: function() {
      return this.listaAlls.filter(lista => {
        return lista.situacao.match(this.searchF);
      });
    }
  },

  methods: {
    listar() {
      Atividades.listar().then(response => {
        this.listaAlls = response.data.atividades;
      });
    }
  }
};
</script>

<style>
.home {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#text-painel {
  opacity: 0.2;
  font-size: 80px;
  text-align: center;
}

#tabela {
  padding-left: 30px;
  padding: 30px;
}

th {
  background-color: black;
  color: black;
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

#text-colunas {
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
  width: auto;
  max-width: 300px;
  margin-left: auto;
  margin-right: auto;
  text-align: left;
  margin-bottom: 10px;
}

#inputTitulo {
  margin-left: 20px;
  margin-right: 40px;
}

#txtDialoghId {
  font-size: 35px;
  text-align: center;
}

#txtDialoghTitulo {
  font-size: 50px;
  text-align: center;
}

#txtDialoghSubTitulo {
  font-size: 18px;
  text-align: center;
  margin-bottom: 25px;
  color: brown;
}

#txtDialoghDescricao {
  font-size: 14px;
  padding-left: 15px;
}
</style>