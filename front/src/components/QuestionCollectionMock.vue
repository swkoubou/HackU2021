<template>
  <div>
    <div v-if="question.questions[questionNo].questionType === 'anaume'">
      <QuestionAnaumeMock :question="question.questions[questionNo]" />
    </div>
    <div v-else-if="question.questions[questionNo].questionType === '4taku'">
      <Question4takuMock :question="question.questions[questionNo]" />
    </div>

    <div class="question-paginator-buttons">
      <button class="question-paginator-button-back" @click="goBackQuestion">
        戻る
      </button>

      <div class="question-number">
        {{ this.questionNo + 1 }} / {{ question.questions.length }}
      </div>

      <button class="question-paginator-button-next" @click="goNextQuestoin">
        次へ
      </button>
    </div>

    <div v-if="this.questionNo >= this.question.questions.length - 1">
      <button>正誤判定をする</button>
    </div>
  </div>
</template>

<script>
import Question4takuMock from '@/components/Question4takuMock.vue'
import QuestionAnaumeMock from '@/components/QuestionAnaumeMock.vue'

export default {
  name: 'QuestionCollectionMock',
  props: ['question'],
  components: {
    Question4takuMock,
    QuestionAnaumeMock,
  },
  data() {
    return {
      questionNo: 0,
    }
  },
  methods: {
    // 前の問題へ
    goBackQuestion() {
      if (this.questionNo <= 0) {
        return
      } else {
        this.questionNo--
      }
    },
    // 次の問題へ
    goNextQuestoin() {
      if (this.questionNo >= this.question.questions.length - 1) {
        return
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
  /* position: absolute; */
  display: flex;
  align-items: center;
  justify-content: space-evenly;
  flex-direction: row;
}

/* .question-paginator-button-back {
  position: relative;
} */
</style>
