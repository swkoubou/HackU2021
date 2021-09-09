<template>
  <div class="question-preview">
    <div
      class="question-preview-header"
      v-if="question.questionType === 'anaume'"
    >
      <IconQuestionAnaume :width="100" :height="100" />
      <div class="question-preview-type-title"><h3>穴埋め問題</h3></div>
    </div>

    <div
      class="question-preview-header"
      v-else-if="question.questionType === '4taku'"
    >
      <IconQuestion4taku :width="100" :height="100" />
      <div class="question-preview-type-title"><h3>４択問題</h3></div>
    </div>

    <div
      class="question-preview-header"
      v-else-if="question.questionType == null"
    >
      <IconQuestionCollection :width="100" :height="100" />
      <h3>問題集</h3>
    </div>
    <div class="question-preview-footer">
      <div v-if="question.questionType != null" class="question-preview-title">
        「{{ question.questionName }}」
      </div>
      <div v-else class="question-preview-title">
        「{{ question.collectionName }}」
      </div>
      <div class="question-preview-info">
        <div class="question-preview-author">
          {{ question.author.userName }}
        </div>
        <div class="question-preview-updatetime">
          {{ updateTime }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import IconQuestion4taku from '@/components/icons/IconQuestion4taku.vue'
import IconQuestionAnaume from '@/components/icons/IconQuestionAnaume.vue'
import IconQuestionCollection from '@/components/icons/IconQuestionCollection.vue'
import { format, parse } from 'date-fns'

export default {
  name: 'QuestionPreview',
  props: ['question'],
  components: {
    IconQuestion4taku,
    IconQuestionAnaume,
    IconQuestionCollection,
  },
  created() {
    console.log(JSON.stringify(this.question))
  },
  computed: {
    updateTime() {
      const updateTime = parse(
        '1999-08-28 23:59:59.999999'.replace(/\.[0-9]*$/g, ''),
        'yyyy-MM-dd HH:mm:ss',
        new Date()
      )
      return format(updateTime, 'yyyy年MM月dd日')
    },
  },
}
</script>

<style scoped>
.question-preview {
  border-radius: 5px;
  width: 350px;
  background-color: aquamarine;
  display: flex;
  flex-flow: column;
  justify-content: space-between;
  border: solid 0.1px #707070;
}

.question-preview-header {
  display: flex;
  flex-flow: row;
  justify-content: space-between;
  padding: 20px;
}

.question-preview-footer {
  background-color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 5px 10px;
  border-top: solid 0.1px #707070;
  border-bottom-left-radius: 4px;
  border-bottom-right-radius: 4px;
}

.question-preview-type-title {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.question-preview-type-title h3 {
  font-weight: 400;
}

.question-preview-info {
  text-align: right;
}
</style>
