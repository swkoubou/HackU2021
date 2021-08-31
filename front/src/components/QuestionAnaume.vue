<template>
  <div id='question-anaume'>
    <h3>{{ question.name }}</h3>
    <div v-for='(wordData, index) in parsedQuestion' :key='`word-${index}`'>
      <div v-if="wordData.word === '[]'">
        <input v-model='answers[wordData.index]' type='text' />
      </div>
      <div v-else>
        {{ wordData.word }}
      </div>
    </div>

    <h3>結果</h3>
    <div v-for='(ans, index) in question.answer' :key='`answer-${index}`'>
      <div>
        答え: {{ ans }} | あなたの入力: {{ answers[index] }} |
        {{ ans === answers[index] ? '正解' : '不正解' }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'QuestionAnaume',
  props: ['question'],
  data() {
    return {
      answers: []
    }
  },
  computed: {
    parsedQuestion() {
      // 入力: "今日は[]のち[]です" | 出力(一部省略): [{今日は}, {0}, {のち}, {1}, {です}]
      // 目的: v-modelと答えを一致させる
      // 参考: https://qiita.com/iwato/items/183e6dd676bf547ea341
      // 分割対象を残した状態で、分割する。
      let questionSplit = this.question.question.split(/(\[\])/g)
      let questionWordWithIndex = []
      let holeIndex = 0
      for (let i = 0; i < questionSplit.length; i++) {
        let word = questionSplit[i]
        let wordData = {
          word: word
        }
        // 穴埋めの、穴の部分であれば、indexをつけてあげる。
        if (word === '[]') {
          wordData['index'] = holeIndex
          holeIndex++
        }
        questionWordWithIndex.push(wordData)
      }
      return questionWordWithIndex
    }
  }
}
</script>

<style scoped></style>
