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
    <h3>Toggle Button Component</h3>
    <p>buttonIsSelected {{ buttonIsSelected }}</p>
    <ToggleButtonComponent v-model="buttonIsSelected">
      テスト
    </ToggleButtonComponent>
    <ToggleButtonComponent
      style="background-color: red"
      v-model="buttonIsSelected"
    >
    </ToggleButtonComponent>
    <button @click="changeToggleButton">
      change buttonSelected from parent
    </button>

    <h3>ホームに表示される問題プレビュー</h3>
    <div class="question-preview-components">
      <QuestionPreview :question="sampleQuestionData.api.q_4taku" />
      <QuestionPreview :question="sampleQuestionData.api.q_anaume" />
      <QuestionPreview :question="sampleQuestionData.api.q_collection" />
    </div>

    <h3>新デザイン 実際に解く問題</h3>
    <div class="question-components">
      <Question :question="sampleQuestionData.api.q_4taku" />
      <Question :question="sampleQuestionData.api.q_anaume" />
      <Question :question="sampleQuestionData.api.q_collection" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import QuestionPreview from '@/components/QuestionPreview.vue'
import Question from '@/components/Question.vue'
import ToggleButtonComponent from '@/components/ToggleButtonComponent.vue'

export default {
  name: 'ComponentsPreviewPage',
  components: {
    QuestionPreview,
    Question,
    ToggleButtonComponent,
  },
  data() {
    return {
      sampleQuestionData: require('@/testdata/question.json'),
      buttonIsSelected: false,
    }
  },
  methods: {
    changeToggleButton: function () {
      this.buttonIsSelected = !this.buttonIsSelected
    },
  },
}
</script>

<style scoped>
.question-components {
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
