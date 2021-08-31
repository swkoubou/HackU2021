<template>
    <div id="question-4taku">
      <h3 class="title" > {{ question.name }} </h3>
      <h3>Q. {{ question.question }}</h3>
      <span></span>
      <div v-for="(value, index) in question.value" :key="index">
        <label class="radio">
          <input type="radio" name="answer" :value="value" v-model="userChoice">
          {{ value }}
        </label>
      </div>
      <h3>あなたの回答: {{ (userChoice === "") ? "未回答" : userChoice }} ({{ (userChoice === question.answer ) ? "正解" : "不正解"}})</h3>
    </div>
</template>

<script>
export default {
    name: "Question4taku",
    props: ["question"],
    data() {
      return {
        userChoice: ""
      }
    },
    setup() {
    },
}
</script>

<style scoped>

</style>
