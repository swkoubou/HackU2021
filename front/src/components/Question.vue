<template>
  <div class="question-card">
    <div v-if="question.questionType === 'anaume'">
      <QuestionAnaume :question="question" :isInCollection="false" />
    </div>
    <div v-else-if="question.questionType === '4taku'">
      <Question4taku :question="question" :isInCollection="false" />
    </div>
    <div v-else-if="question.questionType == null">
      <QuestionCollection :question="question" />
    </div>
    <div v-else>
      <h3>unknown</h3>
    </div>
  </div>
</template>

<script>
import QuestionAnaume from '@/components/QuestionAnaume.vue'
import Question4taku from '@/components/Question4taku.vue'
import QuestionCollection from '@/components/QuestionCollection.vue'

export default {
  name: 'Question',
  props: ['question'],
  components: {
    QuestionAnaume,
    Question4taku,
    QuestionCollection,
  },
  methods: {
    goScorePageAndCheckAnswers(answerData) {
      this.$router.push({
        name: 'ScorePage',
        params: { answersData: [answerData] },
      })
    },
  },
}
</script>

<style scoped>
.question-card {
  border: 2px solid;
  border-color: black;
  display: inline-block;
  height: 500px;
  margin: 1rem;
  position: relative;
  width: 500px;
}

@media (max-width: 500px) {
  .question-card {
    width: 90%;
  }
}
</style>
