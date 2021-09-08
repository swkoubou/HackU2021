<template>
  <div class="question-card">
    <div v-if="question.questionType === 'anaume'">
      <QuestionAnaumeMock :question="question" />
    </div>
    <div v-else-if="question.questionType === '4taku'">
      <Question4takuMock :question="question" />
    </div>
    <div v-else-if="question.questionType == null">
      <QuestionCollectionMock :question="question" />
    </div>
    <div v-else>
      <h3>unknown</h3>
    </div>
  </div>
</template>

<script>
import QuestionAnaumeMock from '@/components/QuestionAnaumeMock.vue'
import Question4takuMock from '@/components/Question4takuMock.vue'
import QuestionCollectionMock from '@/components/QuestionCollectionMock.vue'

export default {
  name: 'QuestionMock',
  props: ['question'],
  components: {
    QuestionAnaumeMock,
    Question4takuMock,
    QuestionCollectionMock,
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
</style>
