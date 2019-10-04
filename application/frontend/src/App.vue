<template>
  <div id="app">
    <q-layout view="hHh lpR fFf">
      <q-header elevated class="bg-primary text-white">
        <q-toolbar>
          <q-toolbar-title>
              <img src="./assets/logo.png" >
          </q-toolbar-title>

          <q-chip>
            <q-avatar>
              <img src="https://scontent.fbsb1-1.fna.fbcdn.net/v/t1.0-1/p160x160/18699989_1475430015854576_5533091456059580637_n.jpg?_nc_cat=110&_nc_oc=AQmTi4CA2E6sfplgrgosCFg7u6bIHtFOs3slwUkBHAmAv-rLIogEPWqPcMi3sfjaFas&_nc_ht=scontent.fbsb1-1.fna&oh=98b878a1b2b861fe0097d6adb86afbd1&oe=5E357B95">
            </q-avatar>
           {{ nomeUser }} |<a href="" id="link-sair">Sair</a>
          </q-chip>

        </q-toolbar>
      </q-header>

      <q-page-container>
        <router-view />
      </q-page-container>

    </q-layout>
  </div>
</template>

<script>
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
#link-sair {
  text-align: center;
  color: green;
  text-decoration:none;
  padding-left: 20px;
  padding-right: 5px;
}

.q-toolbar{
  background-color: white;
}

#info{
  font-size: 14px;
  color: #2c3e50;
}

#btnOptions{
  font-size: 14px;
  color: #2c3e50;
  text-align: center;
}

#socialMidia{
  font-size: 14px;
  color: #2c3e50;
  text-align: right;
}

</style>

      <!--
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
      -->