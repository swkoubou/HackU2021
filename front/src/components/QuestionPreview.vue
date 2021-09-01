<template>
  <div class="question-preview">
    <div
      class="question-preview-header"
      v-if="question.questionType === 'anaume'"
    >
      <h3>icon</h3>
      <h3>穴埋め問題</h3>
    </div>
    <div
      class="question-preview-header"
      v-else-if="question.questionType === '4taku'"
    >
      <h3>icon</h3>
      <h3>四択問題</h3>
    </div>

    <div
      class="question-preview-header"
      v-else-if="question.questionType === 'collection'"
    >
      <h3>icon</h3>
      <h3>問題集</h3>
    </div>

    <div class="question-preview-header" v-else>unknown</div>

    <div class="question-preview-footer">
      <div class="question-preview-title">「{{ question.name }}」</div>
      <div class="question-preview-info">
        <div class="question-preview-author">
          {{ question.author.displayName }}
        </div>
        <div class="question-preview-updatetime">
          {{ question.updateTime }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import '@/components/icons/IconQuestion4taku.vue'

export default {
  name: 'QuestionPreview',
  props: ['question'],
  components: {},
}
</script>

<style scoped>
.question-preview {
  width: 350px;
  height: 180px;
  background-color: aquamarine;
  display: flex;
  flex-flow: column;
  justify-content: space-between;
}

.question-preview-header {
  display: flex;
  flex-flow: row;
  justify-content: space-between;
  padding: 20px;
}

.question-preview-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-right: 10px;
  padding-left: 10px;
}
</style>
