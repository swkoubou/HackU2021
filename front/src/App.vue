<template>
  <div id="app">
    <Header />
    <div id="nav">
      <router-link to="/">Home</router-link>
      |
      <router-link to="/about">About</router-link>
      |
      <router-link to="/components">Components</router-link>
      |
      <router-link to="/answerthequestion">問題回答ページ</router-link>
      |
      <router-link to="/score">スコア</router-link>
      |
      <router-link to="/loginsuccesspreviewpage">ログイン済み確認</router-link>
      |
      <router-link to="/login">Login</router-link>
    </div>
    <router-view />
  </div>
</template>

<script>
import Header from './components/Header.vue'
export default {
  name: 'App',
  components: {
    Header,
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
