<template>
  <div class="components-preview">
    <div class="question-preview">
      <QuestionParent v-bind:questionDataJson="sampleQuestionDataJson.q_4taku" />
      <QuestionParent v-bind:questionDataJson="sampleQuestionDataJson.q_anaume" />
      <QuestionParent v-bind:questionDataJson="sampleQuestionDataJson.q_collection" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import QuestionParent from "@/components/Question.vue";

export default {
  name: "ComponentsPreviewPage",
  components: {
    QuestionParent,
  },
  data () {
    return {
      sampleQuestionDataJson: require("@/testdata/question.json")
    }
  }
};
</script>

<style scoped>
.question-preview {
  display:flex;
  /*flex-flow: column;*/
  justify-content: center;
}
</style>
