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
    if (
      firebaseui.auth.AuthUI.getInstance() != null &&
      firebaseuiAuthContainer.innerHTML == ''
    ) {
      // インスタンスがあるのにUIが消えてたら復元する
      firebaseuiAuthContainer.replaceWith(
        firebaseui.auth.AuthUI.getInstance().D.a
      )
    } else {
      // UIがなければ作る
      const ui =
        firebaseui.auth.AuthUI.getInstance() ||
        new firebaseui.auth.AuthUI(firebase.auth())
      ui.start(firebaseuiAuthContainer, {
        signInOptions: [firebase.auth.EmailAuthProvider.PROVIDER_ID],
        callbacks: {
          signInSuccessWithAuthResult: async (response) => {
            const idToken = await response.user.getIdToken(true)
            localStorage.setItem('login_data', idToken.toString())
            this.$router.push('/loginsuccesspreviewpage')
            return true
          },
        },
      })
    }
  },
}
</script>
