<template>
  <div id="login-page">
    <p>Login page</p>
    <div id="firebaseui-auth-container"></div>
  </div>
</template>

<script>
import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'
import firebaseui from 'firebaseui-ja/dist/npm__ja'
import 'firebaseui-ja/dist/firebaseui.css'
export default {
  name: 'LoginPage',
  mounted() {
    const firebaseuiAuthContainer = document.querySelector(
      '#firebaseui-auth-container'
    )
    if (firebaseuiAuthContainer == null) {
      // domの生成が出来ていない時はuiの生成をやめる
      return
    }
    const ui =
      firebaseui.auth.AuthUI.getInstance() ||
      new firebaseui.auth.AuthUI(firebase.auth())
    ui.start(firebaseuiAuthContainer, {
      signInOptions: [
        firebase.auth.EmailAuthProvider.PROVIDER_ID,
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
      ],
      signInSuccessUrl: 'http://localhost:8080/#/loginsuccesspreviewpage',
      callbacks: {
        signInSuccessWithAuthResult: async (response) => {
          const idToken = await response.user.getIdToken(true)
          localStorage.setItem('login_data', idToken.toString())
          this.$router.push('/loginsuccesspreviewpage')
          ui.reset()
          return false
        },
      },
    })
  },
}
</script>
