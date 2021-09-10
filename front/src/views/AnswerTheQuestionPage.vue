<template>
  <div>
    <Question :question="this.sampleQuestionData" />
  </div>
</template>

<script>
import Question from '@/components/Question.vue'

export default {
  name: 'AnswerTheQuestionPage',
  components: {
    Question,
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
