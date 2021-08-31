<template>
    <div id="question-anaume">
      <h3>{{question.name}}</h3>
      <div v-for="wordData in parsedQuestion" :key="wordData.id">
        <div v-if="wordData.word === '[]'">
          <input type="text" v-model="answers[wordData.index]">
        </div>
        <div v-else>
          {{ wordData.word }}
        </div>
      </div>

      <h3>結果</h3>
      <div v-for="(ans, index) in question.answer" :key="ans.id">
        <div>答え: {{ans}} | あなたの入力: {{answers[index]}} | {{(ans === answers[index]) ? "正解" : "不正解"}}</div>
      </div>

    </div>
</template>

<script>

export default {
  name: "QuestionAnaume",
  props: ["question"],
  data (){
    return {
      answers: []
    };
  },
  computed: {
    parsedQuestion() {
      // 入力: "今日は[]のち[]です" | 出力(一部省略): [{今日は}, {0}, {のち}, {1}, {です}]
      // 目的: v-modelと答えを一致させる
      // 参考: https://qiita.com/iwato/items/183e6dd676bf547ea341
      // 分割対象を残した状態で、分割する。
      let question_split = this.question.question.split(/(\[\])/g);
      let question_word_with_index = [];
      let holeIndex = 0;
      for (let i=0; i<question_split.length; i++) {
        let word = question_split[i];
        let d = {
          "word": word,
        };
        // 穴埋めの、穴の部分であれば、indexをつけてあげる。
        if (word === "[]") {
          d["index"] = holeIndex
          holeIndex++
        }
        question_word_with_index.push(d);
      }
      return question_word_with_index;
    },
  }
}

</script>


<style scoped>
</style>
