<template>
  <div id="question-4taku">
    <h3 class="title">{{ question.name }}</h3>
    <h3>Q. {{ question.question }}</h3>
    <span></span>
    <div v-for="(value, index) in question.value" :key="index">
      <label class="radio">
        <input
          v-model="userAnswers[0]"
          :value="value"
          name="answer"
          type="radio"
        />
        {{ value }}
      </label>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Question4taku',
  model: {
    prop: 'propUserAnswers',
    event: 'change',
  },
  props: {
    question: {},
    propUserAnswers: [],
  },
  computed: {
    userAnswers: {
      get() {
        return this.propUserAnswers
      },
      set(value) {
        console.log(value)
        this.$emit('change', value)
      },
    },
  },
}
</script>

<style scoped></style>
