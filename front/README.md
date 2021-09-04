# HackU2021 front

## Project setup

1. firebaseの設定を環境変数に設定する

```bash
touch public/auth.json
```

#### `public/auth.json` にfirebaseの設定を書く

```json
{
  "apiKey": "",
  "authDomain": "",
  "projectId": "",
  "storageBucket": "",
  "messagingSenderId": "",
  "appId": ""
}
```

2. yarn install

```bash
yarn install
```

3. firebaseuiのライブラリにパッチを当てる

```bash
sed -i  -e "s#^(function() { var firebase=require('firebase/app');require('firebase/auth')#(function() { var firebase=require('firebase/compat/app');require('firebase/compat/auth')#" ./node_modules/firebaseui-ja/dist/npm__ja.js
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
