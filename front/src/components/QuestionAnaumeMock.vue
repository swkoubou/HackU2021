<template>
  <div>
    <h3>「{{ question.questionName }}」</h3>

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
    <div class="anaume-reset">
      <button class="anaume-reset-button" @click="resetSelecting()">
        <FontAwesomeIcon icon="undo" class="anaume-reset-button-icon" /><span
          class="anaume-reset-button-text"
          >選択し直す</span
        >
      </button>

    </div>
  </div>
</template>

<script>
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faUndo } from '@fortawesome/free-solid-svg-icons'

library.add(faUndo)
export default {
  name: 'QuestionAnaumeMock',
  props: {
    question: {
      type: Object,
      required: true,
    },
    isInCollection: {
      type: Boolean,
      required: true,
    },
  },
  components: {
    FontAwesomeIcon,
  },
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

    let questionNoForHuman = 0

    for (let i = 0; i < this.questionSplit.length; i++) {
      if (this.questionSplit[i] == '[]') {
        questionNoForHuman++
        this.squareBracketsIndexs.push(i)
        this.questionSplit[i] = '[ ' + questionNoForHuman + ' ]'
      }
    }
  },
  methods: {
    shuffleArray: function (target) {
      let array = target.concat()
      let currentIndex = array.length,
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
    resetSelecting: function () {
      for (let i = 0; i < this.selectingIndexs.length; i++) {
        this.questionSplit[this.squareBracketsIndexs[i]] = '[ ' + (i + 1) + ' ]'
      }
      this.selectingIndexs = []
    },
    setSelecting: function (index) {
      // すでに選択されているか
      if (this.selectingIndexs.includes(index)) {
        // 最後かどうか
        if (this.selectingIndexs[this.selectingIndexs.length - 1] == index) {
          // 選択を取り消します
          this.questionSplit[
            this.squareBracketsIndexs[this.selectingIndexs.indexOf(index)]
          ] = '[ ' + (this.selectingIndexs.indexOf(index) + 1) + ' ]'
          this.selectingIndexs.splice(this.selectingIndexs.indexOf(index), 1)
        }
      } else {
        this.selectingIndexs.push(index)
      }

      for (let i = 0; i < this.selectingIndexs.length; i++) {
        this.questionSplit[this.squareBracketsIndexs[i]] =
          '[ ' + this.choices[this.selectingIndexs[i]] + ' ]'
      }

      if (this.selectingIndexs.length == this.choices.length) {
        let choicesText = []
        for (let i = 0; i < this.selectingIndexs.length; i++) {
          choicesText.push(this.choices[this.selectingIndexs[i]])
        }

        const answersData = {
          squareBracketsIndexs: this.squareBracketsIndexs,
          questionSplit: this.questionSplit,
          questionAnswers: this.question.answers,
          userAnswers: choicesText,
          type: this.question.questionType,
        }

        if (!this.isInCollection) {
          this.$parent.goScorePageAndCheckAnswers(answersData)
        } else {
          this.$parent.storeUserAnswers(answersData)
          this.$parent.goNextQuestion()
        }
      }
    },
  },
}
</script>

<style scoped>
.anaume-reset-button-text {
  font-size: 15px;
}

.anaume-reset-button-icon {
  padding-right: 5px;
}

.anaume-reset-button {
  font-size: 15px;
  background-color: transparent;
  border: 2px solid;
  border-color: #00237e;
  border-radius: 30px;
}

.anaume-reset {
  display: flex;
  justify-content: flex-end;
  padding-right: 5%;
}

.question-paper {
  display: inline-block;
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
  border: 2px solid;
  border-color: #00237e;
  background-color: #4472c4;
}

.question-choice-list {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
