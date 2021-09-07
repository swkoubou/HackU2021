<template>
  <div>
    <QuestionParent :questionData="this.sampleQuestionData" />
    <!-- <div>{{ this.sampleQuestionData }}</div> -->
  </div>
</template>

<script>
import QuestionParent from '@/components/Question.vue'

export default {
  name: 'AnswerTheQuestionPage',
  components: {
    QuestionParent,
  },
  data() {
    return {
      sampleQuestionDatas: require('@/testdata/question.json'),
    }
  },
  created() {
    let samples = [
      this.sampleQuestionDatas.api.q_4taku,
      this.sampleQuestionDatas.api.q_anaume,
      this.sampleQuestionDatas.api.q_collection,
    ]
    this.sampleQuestionData =
      samples[Math.floor(Math.random() * samples.length)]
  },
}
</script>
