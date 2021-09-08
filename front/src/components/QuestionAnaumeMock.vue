<template>
  <div>
    <h3>
      {{ question.questionName }}
    </h3>

    <h3>
      <div
        v-for="(questionParts, index) in questionSplit"
        :key="`questionparts-${index}`"
        class="question-paper"
      >
        <span>{{ questionParts }}</span>
      </div>
    </h3>

    <div
      class="question-choice-list"
      v-for="(choice, index) in this.choices"
      :key="`choice-${index}`"
    >
      <p class="selecting-mark" v-if="selectingIndexs.indexOf(index) != -1">
        {{ selectingIndexs.indexOf(index) + 1 }}番目
      </p>
      <h3
        class="question-choice"
        :id="`choice-${index}`"
        :class="{ isSelecting: selectingIndexs.indexOf(index) != -1 }"
        @click="setSelecting(index)"
      >
        {{ choice }}
      </h3>
    </div>
    <button>正誤判定をする</button>
  </div>
</template>

<script>
export default {
  name: 'QuestionAnaumeMock',
  props: ['question'],
  data() {
    return {
      selectingIndexs: [],
      choices: [],
      questionSplit: [],
      squareBracketsIndexs: [],
    }
  },
  created() {
    // 答えの順番をシャッフルしたものを選択肢とします。
    this.choices = this.shuffleArray(this.question.answers)
    // 置換のために、四角と文章を分ける
    this.questionSplit = this.question.values[0].split(/(\[\])/g)
    for (let i = 0; i < this.questionSplit.length; i++) {
      if (this.questionSplit[i] == '[]') {
        this.squareBracketsIndexs.push(i)
        this.questionSplit[i] = '[ ]'
      }
    }
  },
  methods: {
    shuffleArray: function (array) {
      var currentIndex = array.length,
        randomIndex

      // While there remain elements to shuffle...
      while (currentIndex != 0) {
        // Pick a remaining element...
        randomIndex = Math.floor(Math.random() * currentIndex)
        currentIndex--

        // And swap it with the current element.
        ;[array[currentIndex], array[randomIndex]] = [
          array[randomIndex],
          array[currentIndex],
        ]
      }
      return array
    },
    setSelecting: function (index) {
      // すでに選択されているか
      if (this.selectingIndexs.includes(index)) {
        // 最後かどうか
        if (this.selectingIndexs[this.selectingIndexs.length - 1] == index) {
          // 選択を取り消します
          // console.log(`${this.squareBracketsIndexs[this.selectingIndexs[index]]} : ${this.questionSplit[this.squareBracketsIndexs[this.selectingIndexs[index]]]}`)
          // this.questionSplit[this.squareBracketsIndexs[this.selectingIndexs[index]]] == "[]"
          this.questionSplit[
            this.squareBracketsIndexs[this.selectingIndexs.indexOf(index)]
          ] = '[ ]'
          this.selectingIndexs.splice(this.selectingIndexs.indexOf(index), 1)
        }
      } else {
        this.selectingIndexs.push(index)
      }

      for (let i = 0; i < this.selectingIndexs.length; i++) {
        this.questionSplit[this.squareBracketsIndexs[i]] =
          '[ ' + this.choices[this.selectingIndexs[i]] + ' ]'
      }
    },
  },
}
</script>

<style scoped>
.question-paper {
  display: contents;
}

.selecting-mark {
  margin: 0;
  height: 40px;
  width: 60px;

  display: flex;
  position: absolute;
  justify-content: center;
  align-items: center;
  left: 20px;

  color: white;
  background-color: #ffc000;
  border: 2px solid;
  border-color: #bc8c00;
}

.isSelecting {
  background-color: #8faadc !important;
}

.question-choice {
  width: 90%;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;

  color: white;
  /* border-top: 2px solid; */
  /* border-bottom: 2px solid; */
  border: 2px solid;
  border-color: #00237e;
  background-color: #4472c4;
}

.question-choice-list {
  /* list-style-type: none; */
  /* text-align:center; */
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
