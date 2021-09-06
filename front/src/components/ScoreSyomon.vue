<template>
  <div class="score-syomon">
    <h4 class="score-syomon-no">問{{ SyomonNo }}</h4>
    <div class="score-compare-answers">
      <div class="score-compare-answers-user">
        あなたの回答 : {{ UserAnswer == null ? '未回答' : UserAnswer }}
      </div>
      <div class="score-compare-answers-question">
        問題の答え : {{ QuestionAnswer }}
      </div>
    </div>
    <div class="isCorrentAnswer-true" v-if="isCorrentAnswer == true">正解</div>
    <div class="isCorrentAnswer-false" v-else>不正解</div>
  </div>
</template>

<script>
export default {
  name: 'ScoreSyomon',
  props: {
    SyomonNo: {
      type: Number,
      required: true,
    },
    UserAnswer: {
      type: String,
      required: true,
    },
    QuestionAnswer: {
      type: String,
      required: true,
    },
  },
  computed: {
    isCorrentAnswer: function () {
      return this.UserAnswer === this.QuestionAnswer
    },
  },
}
</script>

<style scoped>
.score-syomon {
  justify-content: space-evenly;
  text-align: center;
  display: flex;
  align-items: center;
  border: 1px solid black;
}

.score-compare-answers {
  text-align: start;
}

.isCorrentAnswer-true {
  color: green;
}

.isCorrentAnswer-false {
  color: red;
}
</style>
