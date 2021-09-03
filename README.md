# HackU2021

HackU2021にて製作するアプリのリポジトリ


## ビルドの準備

1. firebase の設定をjsonに設定する

```bash
touch front/public/auth.json
```

#### `front/public/auth.json`に firebase の設定を書く

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

#### `./go`に firebase の秘密鍵を書く

```bash
cp ~/Downloads/hogehoge-firebase-adminsdk-hogehoge.json ./go/firebase-adminsdk.json
```

## ビルドの方法

```bash
docker-compose up --build
```

## 開発の流れ

1. issueを切る

   必ずAssignees,Labels,Projects,を設定すること

1. develop branchに移動する

   `git checkout develop`

1. 新しいbranchを切る

   `git checkout -b (front|back|common)/${issue番号}`

1. 開発を行う

   `git add`

   `git commit -m "${やったこと}"`

1. 進捗をGitHubに上げる

   `git push origin ${作業を行ったbranch名}`

1. develop branchにpull requestを出す

### pull requestのReviewersになったら

1. pull requestの内容をチェックし動作確認を行う

1. 問題が無ければpull requestをApproveする

1. Approveした本人が対象のbranchへMergeする

### Labelの種類

    common: front,back共通、または割り切れないタスク

    front: frontendが受け持つタスク

    back: backendが受け持つタスク
