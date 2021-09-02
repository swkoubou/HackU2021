<template>
  <div id="login-page">
    <p>Login page</p>
    <div id="firebaseui-auth-container"></div>
  </div>
</template>

<script>
import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'
import * as firebaseui from 'firebaseui'
import 'firebaseui/dist/firebaseui.css'
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
    let ui =
      firebaseui.auth.AuthUI.getInstance() ||
      new firebaseui.auth.AuthUI(firebase.auth())
    ui.start(firebaseuiAuthContainer, {
      signInOptions: [firebase.auth.EmailAuthProvider.PROVIDER_ID],
      callbacks: {
        signInSuccessWithAuthResult: () => {
          this.$router.push('/loginsuccesspreviewpage')
          return false
        },
      },
    })
  },
}
</script>
