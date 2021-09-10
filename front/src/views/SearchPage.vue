<template>
  <div>
    <!-- <h3>{{searchResultKeys}}</h3> -->
    <div v-if="searchResultKeys.length == 0">
        <h2 class="no-result">結果なし</h2>
    </div>
    <div v-for="searchResultKey in searchResultKeys" :key="searchResultKey">
      <PostList
        :title="searchResultKey"
        :postDataList="questionsWithKeyWord[searchResultKey]"
      />
    </div>
  </div>
</template>

<script>
import PostList from '@/components/QuestionAndCollectionPostList'

export default {
  name: 'SearchResultPage',
  props: {
    questionsWithKeyWord: {
      type: Object,
      required: true,
    },
  },
  components: {
    PostList,
  },
  computed: {
    searchResultKeys: function () {
      return Object.keys(this.questionsWithKeyWord)
    },
  },
}
</script>

<style scoped>
.no-result {
    /* display: flex;
    justify-content: center;
    align-content: center; */
    top: 50%;
    left: 50%;;
}
</style>