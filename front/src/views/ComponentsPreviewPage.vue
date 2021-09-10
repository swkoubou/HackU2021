<template>
  <div class="components-preview">
    <div id="nav">
      <router-link to="/">Home</router-link>
      |
      <router-link to="/about">About</router-link>
      |
      <router-link to="/notification">NotifiCation</router-link>
      |
      <router-link to="/components">Components</router-link>
      |
      <router-link to="/answerthequestion">問題回答ページ</router-link>
      |
      <router-link to="/score">スコア</router-link>
      |
      <router-link to="/loginsuccesspreviewpage">ログイン済み確認</router-link>
      |
      <router-link to="/login">Login</router-link>
    </div>

    <h3>ホームに表示される問題プレビュー</h3>
    <div class="question-preview-components">
      <QuestionPreview :question="sampleQuestionData.api.q_4taku" />
      <QuestionPreview :question="sampleQuestionData.api.q_anaume" />
      <QuestionPreview :question="sampleQuestionData.api.q_collection" />
    </div>

    <h3>新デザイン 実際に解く問題</h3>
    <div class="question-mock-components">
      <QuestionMock :question="sampleQuestionData.api.q_4taku" />
      <QuestionMock :question="sampleQuestionData.api.q_anaume" />
      <QuestionMock :question="sampleQuestionData.api.q_collection" />
    </div>

    <h3>実際に解く問題</h3>
    <div class="question-compoents">
      <QuestionParent :questionData="sampleQuestionData.api.q_4taku" />
      <QuestionParent :questionData="sampleQuestionData.api.q_anaume" />
      <QuestionParent :questionData="sampleQuestionData.api.q_collection" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import QuestionParent from '@/components/Question.vue'
import QuestionPreview from '@/components/QuestionPreview.vue'
import QuestionMock from '@/components/QuestionMock.vue'

export default {
  name: 'ComponentsPreviewPage',
  components: {
    QuestionParent,
    QuestionPreview,
    QuestionMock,
  },
  data() {
    return {
      sampleQuestionData: require('@/testdata/question.json'),
    }
  },
}
</script>

<style scoped>
.question-mock-components {
  display: flex;
  justify-content: space-between;
}

.question-compoents {
  display: flex;
  justify-content: space-between;
}

.question-preview-components {
  display: flex;
  justify-content: space-between;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
