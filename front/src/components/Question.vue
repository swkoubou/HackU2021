<template>
  <div id="question-parent" class="question-card">
    <div v-if="questionData.questionType === '4taku'">
      <Question4taku :question="questionData" v-model="userAnswers[0]" />
    </div>
    <div v-else-if="questionData.questionType === 'anaume'">
      <QuestionAnaume :question="questionData" v-model="userAnswers[0]" />
    </div>
    <div v-else-if="questionData.questionType === 'collection'">
      <!-- QuestionCollectionは、QuestionParentの親の代わりになるので、全体をあげる -->
      <QuestionCollection :question="questionData" v-model="userAnswers" />
    </div>
    <div v-else>
      <h3>unknown : {{ questionData.questionType }}</h3>
    </div>
    <br />
    <button @click="getUserAnswersAndGoScorePage()">スコアを確認する</button>
    <div>ユーザー回答Debug : {{ this.userAnswers }}</div>
  </div>
</template>

<script>
import Question4taku from '@/components/Question4taku.vue'
import QuestionAnaume from '@/components/QuestionAnaume.vue'
import QuestionCollection from '@/components/QuestionCollection.vue'

export default {
  name: 'QuestionParent',
  props: ['questionData'],
  components: {
    Question4taku,
    QuestionAnaume,
    QuestionCollection,
  },
  data() {
    return {
      userAnswers: [],
      answerData: [],
    }
  },
  created() {
    // ユーザーの回答と、問題の回答をスコア画面に持っていくために配列を作ります
    if (this.questionData.questionType === 'collection') {
      for (
        let questionIndex = 0;
        questionIndex < this.questionData.questions.length;
        questionIndex++
      ) {
        this.userAnswers.push([])
        this.answerData[questionIndex] = []
        for (
          let answerIndex = 0;
          answerIndex <
          this.questionData.questions[questionIndex].answers.length;
          answerIndex++
        ) {
          this.answerData[questionIndex].push({
            questionAnswer:
              this.questionData.questions[questionIndex].answers[answerIndex],
            userAnswer: null,
          })
        }
      }
    } else {
      this.userAnswers.push([])
      this.answerData[0] = []
      for (
        let answerIndex = 0;
        answerIndex < this.questionData.answers.length;
        answerIndex++
      ) {
        this.answerData[0].push({
          questionAnswer: this.questionData.answers[answerIndex],
          userAnswer: null,
        })
      }
    }
  },
  methods: {
    getUserAnswersAndGoScorePage() {
      for (
        let userAnswersIndex = 0;
        userAnswersIndex < this.userAnswers.length;
        userAnswersIndex++
      ) {
        for (
          let answerIndex = 0;
          answerIndex < this.userAnswers[userAnswersIndex].length;
          answerIndex++
        ) {
          if (this.userAnswers[userAnswersIndex][answerIndex] == null) {
            this.answerData[userAnswersIndex][answerIndex].userAnswer = null
          } else {
          this.answerData[userAnswersIndex][answerIndex].userAnswer =
            this.userAnswers[userAnswersIndex][answerIndex]}
        }
      }
      this.$router.push({
        name: 'ScorePage',
        params: { answerData: this.answerData },
      })
    },
  },
}
</script>

<style scoped>
.question-card {
  background: #e2e1e0;
  border-radius: 10px;
  display: inline-block;
  height: 500px;
  margin: 1rem;
  position: relative;
  width: 500px;
}
</style>
