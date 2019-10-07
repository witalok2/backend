<template>
  <div class="home">
    <div id="text-painel">
      TODAS AS ATIVIDADES
    </div>
     
    <div class="q-pa-md">
      <q-markup-table flat bordered>
        <thead class="bg-teal">
          <tr>
            <th class="text-left">Titulo</th>
            <th class="text-left">Subtitulo</th>
            <th class="text-left">Descrição</th>
            <th class="text-left">Situação</th>
          </tr>
        </thead>
        <tbody class="bg-grey-3">
          <tr v-for="(lista, index) in listaAll" :key="index">
            <td class="text-left">{{ lista.titulo }}</td>
            <td class="text-left">{{ lista.subtitulo }}</td>
            <td class="text-left">{{ lista.descricao }}</td>
            <td class="text-left">{{ lista.situacao }}</td>
          </tr>
        </tbody>
      </q-markup-table>
    </div>
    

    
   <q-dialog v-model="infoAt" full-height>
        <q-card class="column full-height" style="width: 800px" >
          <q-card-section>
            <div class="text-h6">Informações da Atividade</div>
          </q-card-section>

          <q-card-section class="col">
        
          </q-card-section>

          <q-card-actions align="right" class="bg-white text-teal">
            <q-btn flat label="Criar" v-close-popup />
            <q-btn flat label="Fechar" v-close-popup />
          </q-card-actions>
        </q-card>
      </q-dialog>


    <div class="row">
      <div class="col" id="coluna-pendentes">
        <h6 id="text-colunas"> Pendente </h6>
            <!-- Cards de atividades -->
            <q-card class="my-card bg-purple text-white" v-for="(lista, index) in filteListaPendente" :key="index">
              <q-card-section>
                <div class="text-h6">{{ lista.titulo }}</div>
                 <q-separator />
                <div class="descricao">{{ lista.subtitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ lista.descricao }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
          
              <q-btn round color="deep-orange" icon="edit_location" @click="infoAt = true"/>
              </q-card-actions>

            </q-card><br>
      </div>

      <div class="col" id="coluna-andamento">
        <h6 id="text-colunas"> Em Andamento </h6>
            <!-- Cards de atividades -->
            <q-card class="my-card bg-purple text-white" v-for="(lista, index) in filteListaAndamento" :key="index">
              <q-card-section>
                <div class="text-h6">{{ lista.titulo }}</div>
                 <q-separator />
                <div class="descricao">{{ lista.subtitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ lista.descricao }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
          
<q-btn round color="deep-orange" icon="edit_location" @click="infoAt = true"/>
              </q-card-actions>

            </q-card><br>
      </div>

      <div class="col" id="coluna-finalizada">
        <h6 id="text-colunas"> Finalizada </h6>
            <!-- Cards de atividades -->
            <q-card class="my-card bg-purple text-white" v-for="(lista, index) in filteListaFinalizada" :key="index">
              <q-card-section>
                <div class="text-h6">{{ lista.titulo }}</div>
                 <q-separator />
                <div class="descricao">{{ lista.subtitulo }}</div>
              </q-card-section>

              <q-card-section>
                {{ lista.descricao }}
              </q-card-section>

              <q-separator dark />

              <q-card-actions>
          
              <q-btn round color="deep-orange" icon="edit_location" @click="infoAt = true"/>
              </q-card-actions>

            </q-card><br>      
      </div>
    </div>

  </div>
</template>

<script>
//Imports
import Atividades from '../services/atividades'

export default {
  name:'listar',
  data() { 
    return { 
      listaAll: [],
      infoAt: false,
      searchP:'Pendente',
      searchA:'Em Andamento',
      searchF:'Finalizada',
      text: '',
    } 
  },
  computed: {
    filteListaPendente:function(){
      return this.listaAll.filter((lista) => {
        return lista.situacao.match(this.searchP)
        })
      },

    filteListaAndamento:function(){
      return this.listaAll.filter((lista) => {
        return lista.situacao.match(this.searchA)
        })
      },

  
  filteListaFinalizada:function(){
      return this.listaAll.filter((lista) => {
        return lista.situacao.match(this.searchF)
        })
      },
  },
   mounted(){
    Atividades.lista().then(response => { this.listaAll = response.data.atividades }) 
  },
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


th {
  background-color: black;
  color: black;
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
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
  margin-bottom: 10px;
}

#inputTitulo {
  margin-left: 20px;
  margin-right: 40px;
}

</style>