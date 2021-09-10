<template>
  <div class="notification-page">
    <div class="notification-list">
      <div class="notification-item">
        ２０２０年０８月１２日に問題集「社会史」を作成しました。
      </div>
      <div class="notification-item">
        ２０２０年０８月１２日に問題集「数学」を作成しました。
      </div>
      <div class="notification-item">
        ２０２０年０８月１２日に問題集「現代文」を作成しました。
      </div>
      <div class="notification-item">
        ２０２０年０８月１２日に問題集「現代文」を作成しました。
      </div>
      <div class="notification-item">
        ２０２０年０８月１２日に問題集「オリンピック」を作成しました。
      </div>
    </div>
    <!-- もっと見るは要素数に合わせて表示する予定なので一旦非表示 -->
    <!-- <button class="notification-button">もっと見る</button> -->
  </div>
</template>

<style scoped>
.notification-list {
  display: flex;
  flex-direction: column;
  max-width: 800px;
  width: 100%;
  margin: auto;
}

.notification-item {
  border: solid 1.2px #2c3e50;
  margin: 5px 0px;
  padding: 5px;
  text-align: center;
}

.notification-button {
  margin: 10px;
  padding: 10px 25px;
  border-radius: 25px;
  color: #2c3e50;
  border: solid 1.2px #2c3e50;
  background-color: white;
}

@media (max-width: 840px) {
  .notification-list {
    margin: 0px 20px;
    width: calc(100% - 40px);
  }
}
</style>
