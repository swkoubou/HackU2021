<template>
  <div>
    <p>あなたはすでにログインしています。</p>
    <button @click="logout">ログアウト</button>
  </div>
</template>

<script>
import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'
export default {
  name: 'LoginPage',
  methods: {
    async logout() {
      await firebase.auth().signOut()
      this.$router.push('/login')
    },
  },
}
</script>
