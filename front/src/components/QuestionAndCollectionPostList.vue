<template>
  <div class="post-list">
    <h2 class="post-title">{{ title }}</h2>
    <div class="post-line" v-if="postDataList.length == 0">
      <h3>{{ title }}はありません</h3>
    </div>
    <div class="post-line" v-else>
      <div
        class="line-body"
        v-for="data in postDataList"
        :key="data.questionID"
      >
        <QuestionPreview :question="data" :style="setRandomColorStyle" />
      </div>
    </div>
  </div>
</template>

<script>
import QuestionPreview from '@/components/QuestionPreview.vue'

export default {
  name: 'QuestionAndCollectionPostList',
  components: {
    QuestionPreview,
  },
  data() {
    return {
      // とりあえずランダムに色をつける
      questionPreviewColors: [
        '#A2F1FF',
        '#FFB4C5',
        '#9AF7B6',
        '#FFE796',
        '#D4D9F2',
      ],
    }
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    postDataList: {
      type: Array,
      default: () => [],
      required: false,
    },
  },
  computed: {
    // とりあえずランダムに色をつける
    setRandomColorStyle() {
      const min = Math.ceil(0)
      const max = Math.floor(this.questionPreviewColors.length)
      const index = Math.floor(Math.random() * (max - min) + min)
      return {
        'background-color': this.questionPreviewColors[index],
      }
    },
  },
}
</script>

<style scoped>
.post-list {
  text-align: start;
  margin: 50px 10px;
}
.post-line {
  display: flex;
  gap: 0px 20px;
  overflow-x: scroll;
  height: 215px;
}
.post-title {
  padding: 10px 30px;
  font-size: 1em;
  font-weight: bold;
  border: solid #2c3e50 2px;
  border-radius: 50px;
  display: inline-block;
}
</style>
