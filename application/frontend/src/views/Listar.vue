<template>
  <div class="home"> <!-- v-for="(lista, index) in listaAlls" :key="index" -->
    <div id="text-painel">TODAS AS ATIVIDADES</div>

  
    
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
      all: {}
    };
  },

  mounted() {
    this.listar();
  },

  computed: {
    
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

<style scoped>
.home {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* background-color: beige; */
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