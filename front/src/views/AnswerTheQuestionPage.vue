<template>
  <div>
    <Question :question="this.question" />
  </div>
</template>

<script>
import Question from '@/components/Question.vue'

export default {
  name: 'AnswerTheQuestionPage',
  components: {
    Question,
  },
  props: ['question'],
}
</script>
