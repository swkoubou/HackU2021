<template>
  <div class="components-preview">
    <QuestionParent v-bind:questionDataJson="sampleQuestionDataJson.q_4taku" />
    <QuestionParent v-bind:questionDataJson="sampleQuestionDataJson.q_anaume" />
    <QuestionParent v-bind:questionDataJson="sampleQuestionDataJson.q_collection" />
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
