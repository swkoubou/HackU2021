# HackU2021
HackU2021にて製作するアプリのリポジトリ

## 開発の流れ

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
