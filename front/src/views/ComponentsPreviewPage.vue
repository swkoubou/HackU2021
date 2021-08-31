<template>
  <div class="components-preview">
    <div class="question-preview">
      <QuestionParent :questionData="sampleQuestionData.q_4taku" />
      <QuestionParent :questionData="sampleQuestionData.q_anaume" />
      <QuestionParent :questionData="sampleQuestionData.q_collection" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import QuestionParent from '@/components/Question.vue'

export default {
  name: 'ComponentsPreviewPage',
  components: {
    QuestionParent,
  },
  data() {
    return {
      sampleQuestionData: require('@/testdata/question.json'),
    }
  },
}
</script>

<style scoped>
.question-preview {
  display: flex;
  /*flex-flow: column;*/
  justify-content: center;
}
</style>
