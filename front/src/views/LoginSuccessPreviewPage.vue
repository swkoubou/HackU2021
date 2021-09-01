<template>
  <div>
    <p>あなたはすでにログインしています。</p>
    <button @click="logout">ログアウト</button>
  </div>
</template>

<script>
import firebase from 'firebase/compat'
export default {
  name: 'LoginPage',
  methods: {
    async logout() {
      console.log('logout')
      await firebase.auth().signOut()
      this.$router.push('/login')
    },
  },
}
</script>
