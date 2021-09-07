<template>
  <div class="components-preview">
    <h3>実際に解く問題</h3>
    <div class="question-compoents">
      <QuestionParent :questionData="sampleQuestionData.api.q_4taku" />
      <QuestionParent :questionData="sampleQuestionData.api.q_anaume" />
      <QuestionParent :questionData="sampleQuestionData.api.q_collection" />
    </div>

    <h3>ホームに表示される問題プレビュー</h3>
    <div class="question-preview-components">
      <QuestionPreview :question="sampleQuestionData.api.q_4taku" />
      <QuestionPreview :question="sampleQuestionData.api.q_anaume" />
      <QuestionPreview :question="sampleQuestionData.api.q_collection" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import QuestionParent from '@/components/Question.vue'
import QuestionPreview from '@/components/QuestionPreview.vue'

export default {
  name: 'ComponentsPreviewPage',
  components: {
    QuestionParent,
    QuestionPreview,
  },
  data() {
    return {
      sampleQuestionData: require('@/testdata/question.json'),
    }
  },
}
</script>

<style scoped>
.question-compoents {
  display: flex;
  justify-content: space-between;
}

.question-preview-components {
  display: flex;
  justify-content: space-between;
}
</style>
