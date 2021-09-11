<template>
  <button class="toggle-button" @click="click" :class="{ selected: selected }">
    <slot></slot>
  </button>
</template>

<script>
export default {
  name: 'QuestionPreview',
  props: {
    selected: {
      type: Boolean,
      default: () => false,
    },
  },
  methods: {
    click: function () {
      this.selected = !this.selected
      this.$emit('click', this.selected)
    },
  },
  model: {
    prop: 'selected',
    event: 'click',
  },
}
</script>

<style scoped>
.toggle-button {
  border: solid 2px black;
  background: white;
  color: black;
  padding: 5px 10px;
}

.selected {
  border: solid 2px #1eb776;
  color: #1eb776;
}
</style>
