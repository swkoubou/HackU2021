<template>
  <div>
    <QuestionMock :question="this.sampleQuestionData" />
  </div>
</template>

<script>
import QuestionMock from '@/components/QuestionMock.vue'

export default {
  name: 'AnswerTheQuestionPage',
  components: {
    QuestionMock,
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
