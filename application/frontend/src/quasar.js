import Vue from 'vue'

import './styles/quasar.styl'
import 'quasar/dist/quasar.ie.polyfills'
import lang from 'quasar/lang/pt-br.js'
import '@quasar/extras/roboto-font/roboto-font.css'
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/eva-icons/eva-icons.css'
import {
  Quasar, 
  QLayout,
  QHeader,
  QDrawer,
  QPageContainer,
  QPage,
  QToolbar,
  QToolbarTitle,
  QBtn,
  QIcon,
  QList,
  QItem,
  QItemSection,
  QItemLabel,
  QAvatar,
  QChip,
  QMarkupTable,
  QTable,
  QTh,
  QTr,
  QTd,
  QCard,
  QCardSection,
  QCardActions,
  QDialog,
  ClosePopup,
  QInput,
  QBadge,
  Notify,
  QFab,
  QFabAction
} from 'quasar'

Vue.use(Quasar, {
  plugins: {
    Notify
  },
  config: {
    notify: { /* Notify defaults */ }
  },
  components: {
    QLayout,
    QHeader,
    QDrawer,
    QPageContainer,
    QPage,
    QToolbar,
    QToolbarTitle,
    QBtn,
    QIcon,
    QList,
    QItem,
    QItemSection,
    QItemLabel,
    QAvatar,
    QChip,
    QMarkupTable,
    QTable,
    QTh,
    QTr,
    QTd,
    QCard,
    QCardSection,
    QCardActions,
    QDialog,
    QInput,
    QBadge,
    QFab,
    QFabAction,
  },
  directives: {
    ClosePopup
  },
  lang: lang
 })