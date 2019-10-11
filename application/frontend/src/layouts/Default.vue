<template>
<div id="home">
  <q-layout view="hHh Lpr lff">
    <q-header elevated class="bg-black">
        <q-toolbar>
          <q-btn flat @click="drawer = !drawer" round dense icon="menu" />
          <q-toolbar-title style="text-align: center;">Atividades</q-toolbar-title>
        </q-toolbar>
      </q-header>

      <q-drawer
        v-model="drawer"
        show-if-above

        :mini="miniState"
        @mouseover="miniState = false"
        @mouseout="miniState = true"

        :width="200"
        :breakpoint="500"
        bordered
        content-class="bg-grey-3"
      >
        <q-scroll-area class="fit">
          <q-list padding>
            <q-item clickable v-ripple>
              <q-item-section avatar>
                <q-icon name="today" />
              </q-item-section>

              <q-item-section>
                <router-link to="/criar" style="text-decoration:none;">Criar</router-link>
              </q-item-section>
            </q-item>

            <q-item clickable v-ripple>
              <q-item-section avatar>
                <q-icon name="list" />
              </q-item-section>

              <q-item-section>
                <router-link to="/" style="text-decoration:none;">Atividades</router-link>
              </q-item-section>
            </q-item>
          
          </q-list>
        </q-scroll-area>
      </q-drawer>

    <q-page-container>
      <q-page padding>
        
           <router-view />
          
        </q-page>
     
    </q-page-container>
  </q-layout>
</div>
</template>

<script>
//imports
import MenuIcon from 'vue-material-design-icons/Menu.vue';

export default {
  name: 'LayoutDefault',
  data () {
    return {
      leftDrawerOpen: this.$q.platform.is.desktop,
      drawer: false,
      miniState: true
    }
  }
}
</script>

<style scoped>
#home {
  background-color: azure;
}
</style>