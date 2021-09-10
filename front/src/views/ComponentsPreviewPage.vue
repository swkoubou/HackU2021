<template>
  <div class="components-preview">
    <h3>ホームに表示される問題プレビュー</h3>
    <div class="question-preview-components">
      <QuestionPreview :question="sampleQuestionData.api.q_4taku" />
      <QuestionPreview :question="sampleQuestionData.api.q_anaume" />
      <QuestionPreview :question="sampleQuestionData.api.q_collection" />
    </div>

    <h3>新デザイン 実際に解く問題</h3>
    <div class="question-components">
      <Question :question="sampleQuestionData.api.q_4taku" />
      <Question :question="sampleQuestionData.api.q_anaume" />
      <Question :question="sampleQuestionData.api.q_collection" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import QuestionPreview from '@/components/QuestionPreview.vue'
import Question from '@/components/Question.vue'

export default {
  name: 'ComponentsPreviewPage',
  components: {
    QuestionPreview,
    Question,
  },
  data() {
    return {
      sampleQuestionData: require('@/testdata/question.json'),
    }
  },
}
</script>

<style scoped>
.question-components {
  display: flex;
  justify-content: space-between;
}

.question-preview-components {
  display: flex;
  justify-content: space-between;
}
</style>
