<template>
  <main class="home-page">
    <PostList title="最新の問題" :postDataList="newQuestionPosts" />
    <PostList title="最新の問題集" :postDataList="newCollectionPosts" />
    <PostList
      title="みんなが挑戦している問題"
      :postDataList="newQuestionPosts"
    />
    <PostList
      title="みんなが挑戦している問題集"
      :postDataList="newCollectionPosts"
    />
  </main>
</template>

<script>
import PostList from '@/components/QuestionAndCollectionPostList'

export default {
  name: 'Home',
  components: {
    PostList,
  },
  data() {
    return {
      newQuestionPosts: [],
      newCollectionPosts: [],
    }
  },
  async created() {
    {
      const response = await fetch('api/question/all')
      const allQuestions = await response.json()
      this.newQuestionPosts = allQuestions
    }
    {
      const response = await fetch('api/collection/all')
      const allCollection = await response.json()
      this.newCollectionPosts = allCollection
    }
  },
}
</script>

<style scoped>
.new-question-post {
  text-align: start;
  margin: 20px 10px;
}
.new-question-post-line {
  display: flex;
  gap: 0px 20px;
  overflow-x: scroll;
}
.new-question-post-title {
  padding: 10px 30px;
  font-size: 1em;
  font-weight: bold;
  border: solid #2c3e50 2px;
  border-radius: 50px;
  display: inline-block;
}
</style>
