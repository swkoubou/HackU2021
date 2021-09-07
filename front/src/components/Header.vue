<template>
  <header class="header">
    <button class="header-button" @click="notImplement">
      <FontAwesomeIcon icon="home" />
    </button>
    <input
      type="text"
      class="search-box"
      placeholder="問題や問題集を検索する"
    />
    <button class="header-button" @click="notImplement">
      <FontAwesomeIcon icon="pencil-alt" />
    </button>
    <button class="header-button" @click="notImplement">
      <FontAwesomeIcon icon="bell" />
    </button>
    <button class="header-button" @click="notImplement">
      <FontAwesomeIcon icon="user-circle" />
    </button>
  </header>
</template>

<script>
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faBell,
  faHome,
  faPencilAlt,
  faSearch,
  faUserCircle,
} from '@fortawesome/free-solid-svg-icons'

library.add(faHome, faSearch, faPencilAlt, faBell, faUserCircle)
export default {
  name: 'Header',
  components: {
    FontAwesomeIcon,
  },
  methods: {
    notImplement() {
      alert('まだないよ')
    },
  },
}
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  background-color: #ffdbdb;
  padding: 20px;
  margin: 0px;
}

.header-button {
  height: 100%;
  font-size: 1.5em;
  margin: 0px;
  flex: 0 0 auto;
  margin: 0px 5px;
  background-color: transparent;
  border: none;
}

.search-box {
  flex: 1 1 auto;
  margin: 0px 5px;
  border-radius: 10px;
  border: none;
}
</style>
