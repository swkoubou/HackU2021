<template>
  <div class="question-4taku">
    <h3 class="question-name">「{{ question.questionName }}」</h3>
    <h3>
      {{ question.questionBody }}
    </h3>

    <div
      class="question-choice-list"
      v-for="(choice, index) in question.values"
      :key="`choice-${index}`"
    >
      <p class="selecting-mark" v-if="selectingQuestionIndex === index">
        選択中
      </p>
      <h3
        class="question-choice"
        :id="`choice-${index}`"
        :class="{
          isSelecting: selectingQuestionIndex === index ? true : false,
        }"
        @click="setSelecting(index)"
      >
        {{ choice }}
      </h3>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Question4taku',
  props: {
    question: {
      type: Object,
      required: true,
    },
    isInCollection: {
      type: Boolean,
      required: true,
    },
  },
  data() {
    return {
      selectingQuestionIndex: null,
    }
  },
  created(){
    this.selectingQuestionIndex = null;
  },
  methods: {
    setSelecting: function (index) {
      if (this.selectingQuestionIndex != index) {
        this.selectingQuestionIndex = index
      } else {
        this.selectingQuestionIndex = null
      }

      const answersData = {
        questionAnswers: this.question.answers,
        userAnswers: [this.question.values[this.selectingQuestionIndex]],
        type: this.question.questionType,
      }

      if (!this.isInCollection) {
        this.$parent.goScorePageAndCheckAnswers(answersData)
      } else {
        this.$parent.storeUserAnswers(answersData)
        this.selectingQuestionIndex = null
        this.$parent.goNextQuestion()
      }
    },
  },
}
</script>

<style scoped>
.selecting-mark {
  margin: 0;
  height: 40px;
  width: 60px;

  display: flex;
  position: absolute;
  justify-content: center;
  align-items: center;
  left: 20px;

  color: white;
  background-color: #ffc000;
  border: 2px solid;
  border-color: #bc8c00;
}

.isSelecting {
  background-color: #8faadc !important;
}

.question-choice {
  width: 90%;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;

  background-color: #4472c4;
  color: white;
  border: 2px solid;
  border-color: #00237e;
}

.question-choice-list {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
