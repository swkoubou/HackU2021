<template>
  <div id="question-parent" class="question-card">
    <div v-if="questionData.questionType === '4taku'">
      <Question4taku :question="questionData" v-model="userAnswers[0]" />
    </div>
    <div v-else-if="questionData.questionType === 'anaume'">
      <QuestionAnaume :question="questionData" v-model="userAnswers[0]" />
    </div>
    <div v-else-if="questionData.questionType === 'collection'">
      <!-- QuestionCollectionは、QuestionParentの親の代わりになるので、全体をあげる -->
      <QuestionCollection :question="questionData" v-model="userAnswers" />
    </div>
    <div v-else>
      <h3>unknown : {{ questionData.questionType }}</h3>
    </div>
    <br />
    <button @click="checkTheAnswer">答え合わせ</button>
    <div>ユーザー回答Debug : {{ this.userAnswers }}</div>
    <div>答案Debug : {{ this.questionAnswers }}</div>
  </div>
</template>

<script>
import Question4taku from '@/components/Question4taku.vue'
import QuestionAnaume from '@/components/QuestionAnaume.vue'
import QuestionCollection from '@/components/QuestionCollection.vue'

export default {
  name: 'QuestionParent',
  props: ['questionData'],
  components: {
    Question4taku,
    QuestionAnaume,
    QuestionCollection,
  },
  data() {
    return {
      userAnswers: [],
      questionAnswers: [],
    }
  },
  created() {
    // ユーザーの回答を保存するための配列を作ります。
    // ついでに答案、答案データ格納用のmapも作ります。
    if (this.questionData.questionType === 'collection') {
      for (let i = 0; i < this.questionData.questions.length; i++) {
        this.userAnswers.push([])
        // this.questionAnswers[i] = this.questionData.questions[i].answers
        this.questionAnswers[i] = []
        for (
          let j = 0;
          j < this.questionData.questions[i].answers.length;
          j++
        ) {
          this.questionAnswers[i].push({
            questionAnswer: this.questionData.questions[i].answers[j],
            userAnswer: '',
          })
        }
      }
    } else {
      this.userAnswers.push([])
      // this.questionAnswers[0] = this.questionData.answers
      this.questionAnswers[0] = []
      for (let k = 0; k < this.questionData.answers.length; k++) {
        this.questionAnswers[0].push({
          questionAnswer: this.questionData.answers[k],
          userAnswer: '',
        })
      }
    }
  },
  methods: {
    // todo
    checkTheAnswer() {
      for (let i = 0; i < this.userAnswers.length; i++) {
        // 問題ごとに
        for (let j = 0; j < this.userAnswers[j].length; j++) {
          // 回答ごとに
        }
      }
    },
  },
}
</script>

<style scoped>
.question-card {
  background: #e2e1e0;
  border-radius: 10px;
  display: inline-block;
  height: 500px;
  margin: 1rem;
  position: relative;
  width: 500px;
}
</style>
