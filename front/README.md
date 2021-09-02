# HackU2021 front

## Project setup

1. firebaseの設定を環境変数に設定する

```
export API_KEY=
export AUTH_DOMAIN=
export PROJECT_ID=
export STORAGE_BUCKET=
export MESSAGING_SENDER_ID=
export APP_ID=
```

2. yarn install

```bash
yarn install
```

3. firebaseuiのライブラリにパッチを当てる

```bash
sed -i -e "s#^import firebase from 'firebase/app'#import firebase from 'firebase/compat/app'#" ./node_modules/firebaseui/dist/esm.js
```

### Compiles and hot-reloads for development

```
yarn serve
```

### Compiles and minifies for production

```
yarn build
```

### Lints and fixes files

```
yarn lint
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).
