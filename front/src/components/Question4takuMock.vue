<template>
  <div class="question-4taku">
    <h3 class="question-name">
      {{ question.questionName }}
    </h3>
    <h4>
      {{ question.questionBody }}
    </h4>

    <div
      class="question-choice-list"
      v-for="(choice, index) in question.values"
      :key="`choice-${index}`"
    >
      <p class="selecting-mark" v-if="selectingQuestionIndex === index">
        選択中
      </p>
      <h3
        class="question-choice"
        :id="`choice-${index}`"
        :class="{
          isSelecting: selectingQuestionIndex === index ? true : false,
        }"
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
  name: 'Question4takuMock',
  props: ['question'],
  data() {
    return {
      selectingQuestionIndex: null,
    }
  },
  methods: {
    setSelecting: function (index) {
      // let choiceData = document.getElementById(`choice-${v}`).getBoundingClientRect();
      // console.log(`top: ${choiceData.top}, right: ${choiceData.right}, left: ${choiceData.left}, bottom: ${choiceData.bottom}`)
      //   let choiceData = document.getElementById(`choice-${index}`)
      //   console.log(
      //     `top: ${choiceData.offsetTop}, left: ${choiceData.offsetLeft}`
      //   )
      if (this.selectingQuestionIndex != index) {
        this.selectingQuestionIndex = index
      } else {
        this.selectingQuestionIndex = null
      }
    },
  },
}
</script>

<style scoped>
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

  background-color: #4472c4;
  color: white;
  border: 2px solid;
  border-color: #00237e;
}

.question-choice-list {
  /* list-style-type: none; */
  /* text-align:center; */
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
