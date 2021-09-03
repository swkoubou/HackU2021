<template>
  <div>
    <p>あなたはすでにログインしています。</p>
    <button @click="logout">ログアウト</button>
    <button @click="accsessPrivateData">APIアクセスを試す</button>
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
      localStorage.removeItem('login_data')
      this.$router.push('/login')
    },
    async accsessPrivateData() {
      const response = await fetch('http://localhost:8080/api/private/mydata', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('login_data')}`,
        },
      })
      const text = await response.text()
      alert(text)
    },
  },
}
</script>
