<template>
  <div>
    <div v-if="question.questions[questionNo].questionType === 'anaume'">
      <QuestionAnaume
        :question="question.questions[questionNo]"
        :isInCollection="true"
      />
    </div>
    <div v-else-if="question.questions[questionNo].questionType === '4taku'">
      <Question4taku
        :question="question.questions[questionNo]"
        :isInCollection="true"
      />
    </div>

    <div class="question-paginator-buttons">
      <div class="question-number">
        {{ this.questionNo + 1 }} / {{ question.questions.length }}
      </div>
    </div>
  </div>
</template>

<script>
import Question4taku from '@/components/Question4taku.vue'
import QuestionAnaume from '@/components/QuestionAnaume.vue'

export default {
  name: 'QuestionCollection',
  props: ['question'],
  components: {
    Question4taku,
    QuestionAnaume,
  },
  data() {
    return {
      questionNo: 0,
      userAnswersPerQuestion: [],
    }
  },
  methods: {
    storeUserAnswers(answerData) {
      this.userAnswersPerQuestion[this.questionNo] = answerData
    },

    // 次の問題へ
    goNextQuestion() {
      if (this.questionNo >= this.question.questions.length - 1) {
        this.$router.push({
          name: 'ScorePage',
          params: { answersData: this.userAnswersPerQuestion },
        })
      } else {
        this.questionNo++
      }
    },
  },
}
</script>

<style scoped>
.question-paginator-button-back {
  width: 50px;
  height: 28px;
}

.question-paginator-button-next {
  width: 50px;
  height: 28px;
}

.question-number {
  width: 50px;
  height: 28px;

  display: flex;
  justify-content: center;
  align-items: center;
}

.question-paginator-buttons {
  display: flex;
  align-items: center;
  justify-content: space-evenly;
  flex-direction: row;
}
</style>
