<template>
  <header class="header">
    <div
      class="header-context-menu-background"
      @click="toggleContextMenu"
      v-if="isContextMenuOpen"
    >
      <div class="header-context-menu">
        <button @click="notImplement">マイページ</button>
        <button @click="notImplement">設定</button>
        <button @click="toggleContextMenu">閉じる</button>
      </div>
    </div>
    <button class="header-button" @click="changeHomePage">
      <FontAwesomeIcon icon="home" />
    </button>
    <div class="search-box search-box-pc-window">
      <input
        type="search"
        placeholder="問題や問題集を検索する"
        v-model="searchValue"
      />
      <button class="search-button" @click="searchOnEnter">
        <FontAwesomeIcon icon="search" />
      </button>
    </div>

    <div
      class="search-box search-box-smartphone-window"
      v-show="isSearchBoxOpenFromSmartPhoneWindow"
    >
      <input
        type="search"
        placeholder="問題や問題集を検索する"
        v-model="searchValue"
      />
      <button class="search-button" @click="searchOnEnter">
        <FontAwesomeIcon icon="search" />
      </button>
    </div>
    <button
      class="header-button search-toggle-button"
      @click="toggleSearchBoxFromSmartPhone"
    >
      <FontAwesomeIcon icon="search" />
    </button>
    <button class="header-button" @click="notImplement">
      <FontAwesomeIcon icon="pencil-alt" />
    </button>
    <button class="header-button" @click="changeNotifiCationPage">
      <FontAwesomeIcon icon="bell" />
    </button>
    <button class="header-button" @click="toggleContextMenu">
      <FontAwesomeIcon icon="user-circle" />
    </button>
  </header>
</template>

<script>
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faBell,
  faHome,
  faPencilAlt,
  faSearch,
  faUserCircle,
} from '@fortawesome/free-solid-svg-icons'

library.add(faHome, faSearch, faPencilAlt, faBell, faUserCircle)
export default {
  name: 'Header',
  components: {
    FontAwesomeIcon,
  },
  data() {
    return {
      isContextMenuOpen: false,
      isSearchBoxOpenFromSmartPhoneWindow: false,
      searchValue: '',
      testQuestionData: require('@/testdata/question.json'),
    }
  },
  computed: {
    searchSource: function () {
      return this.testQuestionData.searchSource
    },
  },

  methods: {
    notImplement() {
      alert('まだないよ')
    },
    changeHomePage() {
      this.$router.push({ path: '/' })
    },
    changeNotifiCationPage() {
      this.$router.push({ path: '/notification' })
    },

    // -- 検索機能ローカルテスト用関数群:開始 --

    // タグで検索
    searchQuestionsWithTags(tags) {
      // todo : 問題毎に一致するタグがあるかforで探すとても非効率的な方法なので、検索方法、要検討

      // 検索結果を格納する入れ物を用意します
      // 何で検索したのか: []問題達
      let result = {}
      for (
        let searchSourceIndex = 0;
        searchSourceIndex < this.searchSource.length;
        searchSourceIndex++
      ) {
        // 取り出した問題に、名前をつけます
        const question = this.searchSource[searchSourceIndex]
        // 一致するタグがあるか探します。
        for (let tagsIndex = 0; tagsIndex < tags.length; tagsIndex++) {
          // 取り出したタグに名前をつけます
          const tagRaw = tags[tagsIndex]
          const tag = tagRaw.replace('#', '')
          // もし、タグを含んでいれば、結果に追加します
          if (question.questionTag.includes(tag)) {
            // タグが結果に存在するか確認し、なかったら空の配列を設置します
            if (!Object.prototype.hasOwnProperty.call(result, tagRaw)) {
              result[tagRaw] = []
            }
            // 結果に追加します
            result[tagRaw].push(question)
          }
        }
      }
      return result
    },

    // 通常文字で検索
    // [ 優先順位 ]
    // 1. 通常文字に一致するタイトルを持つ問題を探す
    // 2. 通常文字に一致するquestionBodyを持つ問題を探す
    searchQuestionsWithNormalWords(normalWords) {
      // todo : 問題毎に一致する文字列があるかforで探すとても非効率的な方法なので、検索方法、要検討
      // 部分一致参考 :
      // https://qiita.com/aqril_1132/items/9f69575bfbcf24bdf7b5#%E9%83%A8%E5%88%86%E4%B8%80%E8%87%B4

      // 検索結果を格納する入れ物を用意します
      // 何で検索したのか: []問題達
      let result = {}
      for (
        let searchSourceIndex = 0;
        searchSourceIndex < this.searchSource.length;
        searchSourceIndex++
      ) {
        // 取り出した問題に名前をつけます。
        const question = this.searchSource[searchSourceIndex]
        // 通常文字列を取り出し、名前をつけます。
        for (
          let normalWordsIndex = 0;
          normalWordsIndex < normalWords.length;
          normalWordsIndex++
        ) {
          const normalWord = normalWords[normalWordsIndex]
          // もし、タイトルと通常文字列が部分一致したら、結果に追加します。
          // また、もし、questionBodyと通常文字列が部分一致したら、結果に追加します。

          // 同じものが、複数存在することを避けるため、else ifで分岐します。
          if (question.questionName.indexOf(normalWord) > -1) {
            // 通常文字が結果になかったら空の配列を用意します。
            if (!Object.prototype.hasOwnProperty.call(result, normalWord)) {
              result[normalWord] = []
            }
            // 結果に追加します。
            result[normalWord].push(question)
          } else if (question.questionBody.indexOf(normalWord) > -1) {
            // 通常文字が結果になかったら空の配列を用意します。
            if (!Object.prototype.hasOwnProperty.call(result, normalWord)) {
              result[normalWord] = []
            }
            // 結果に追加します。
            result[normalWord].push(question)
          }
        }
      }
      return result
    },

    // 検索欄でエンターが押された時に呼び出される関数
    searchOnEnter() {
      // 検索欄から得られた、タグ達を入れる物です。
      let tags = []
      // 検索欄から得られた、通常文字達を入れる物です。
      let normalWords = []

      // 全角の空白を全て半角の空白に変更します。
      const searchValueHankaku = this.searchValue.replace('　', ' ')

      // 関数を分けたいので、タグと、通常文字を分けます。
      // 空白で区切ります。

      const searchValueSplit = searchValueHankaku.split(' ')
      for (let i = 0; i < searchValueSplit.length; i++) {
        // 区切られた中身を取り出し、使いやすいように名前をつけます。
        const word = searchValueSplit[i]
        // もし、一番最初の文字が「#」ならタグです。
        // それ以外は通常文字です。
        if (word[0] == '#') {
          tags.push(word)
        } else {
          normalWords.push(word)
        }
      }

      // 実際に検索を行います。
      const questionsFromTagSearch = this.searchQuestionsWithTags(tags)
      const questionsFromNormalWordSearch =
        this.searchQuestionsWithNormalWords(normalWords)

      // 結果を取り出すため、鍵一覧の配列を準備します。
      const questionsFromTagSearchKeys = Object.keys(questionsFromTagSearch)
      const questionsFromNormalWordSearchKeys = Object.keys(
        questionsFromNormalWordSearch
      )

      // 検索結果の入れ物を用意します。
      // 基本的に重なったキーワードがあった場合、上書きされます。
      let resultQuestions = {}

      for (
        let questionsFromTagSearchKeysIndex = 0;
        questionsFromTagSearchKeysIndex < questionsFromTagSearchKeys.length;
        questionsFromTagSearchKeysIndex++
      ) {
        // 鍵を取り出し、名前をつけます。
        const questionsFromTagSearchKey =
          questionsFromTagSearchKeys[questionsFromTagSearchKeysIndex]
        // 各々の検索結果を全体の検索結果に統合します。
        resultQuestions[questionsFromTagSearchKey] =
          questionsFromTagSearch[questionsFromTagSearchKey]
      }

      for (
        let questionsFromNormalWordSearchKeysIndex = 0;
        questionsFromNormalWordSearchKeysIndex <
        questionsFromNormalWordSearchKeys.length;
        questionsFromNormalWordSearchKeysIndex++
      ) {
        // 取り出して、名前をつけます。
        const questionsFromNormalWordSearchKey =
          questionsFromNormalWordSearchKeys[
            questionsFromNormalWordSearchKeysIndex
          ]
        resultQuestions[questionsFromNormalWordSearchKey] =
          questionsFromNormalWordSearch[questionsFromNormalWordSearchKey]
      }

      // ユーザーが検索欄に入力した内容を消します。
      this.searchValue = ''

      const SearchKeysForQueryRaw = [
        ...questionsFromTagSearchKeys,
        ...questionsFromNormalWordSearchKeys,
      ]
      let SearchKeysForQuerySafe = {}
      for (
        let SearchKeysForQueryRawIndex = 0;
        SearchKeysForQueryRawIndex < SearchKeysForQueryRaw.length;
        SearchKeysForQueryRawIndex++
      ) {
        const query = SearchKeysForQueryRaw[SearchKeysForQueryRawIndex]
        const safeQuery = encodeURIComponent(query)
        SearchKeysForQuerySafe[`q${SearchKeysForQueryRawIndex + 1}`] = safeQuery
      }

      this.$router.push({
        path: 'search',
        name: 'SearchPage',
        params: { questionsWithKeyWord: resultQuestions },
        query: SearchKeysForQuerySafe,
      })
    },
    // -- 検索機能ローカルテスト用関数群:終了 --

    toggleContextMenu() {
      this.isContextMenuOpen = !this.isContextMenuOpen
    },
    toggleSearchBoxFromSmartPhone() {
      this.isSearchBoxOpenFromSmartPhoneWindow =
        !this.isSearchBoxOpenFromSmartPhoneWindow
    },
  },
}
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  background-color: #ffdbdb;
  padding: 20px;
  margin: 0px;
}

.header-button {
  height: 100%;
  font-size: 1.5em;
  margin: 0px;
  flex: 0 0 auto;
  margin: 0px 5px;
  background-color: transparent;
  border: none;
}

.search-box {
  background-color: white;
  border-radius: 10px;
  border: solid 2px transparent;
  outline: solid 2px transparent;
  outline-offset: 1px;
  display: flex;
}

.search-box:focus-within {
  border: solid 2px #0060df;
  outline: solid 2px #b2ddf9;
}

.search-box input[type='search'] {
  flex: 1 1 auto;
  border: none;
  border-top-left-radius: 11px;
  border-bottom-left-radius: 11px;
}

.search-box input[type='search']:focus {
  outline: none;
}

.search-button {
  background-color: #0060df;
  border: solid 1px #0060df;
  padding: 0px 8px;
  border-top-right-radius: 11px;
  border-bottom-right-radius: 11px;
}

.search-box-pc-window {
  flex: 1 1 auto;
  margin: 0px 5px;
}

.search-box-smartphone-window {
  position: absolute;
  top: 80px;
  height: 1.2em;
  font-size: 1.2em;
  border: solid 2px transparent;
  outline: solid 2px #2c3e50;
  outline-offset: -3px;
  width: calc(100% - 40px);
}

.search-box-smartphone-window:focus-within {
  outline-offset: 1px;
}

.search-button-smartphone-window {
  position: absolute;
  top: 80px;
  font-size: 1.2em;
  width: calc(100% - 40px);
}

.search-toggle-button {
  display: none;
}

@media (max-width: 600px) {
  .search-box-pc-window {
    display: none;
  }

  .search-toggle-button {
    display: block;
  }
}

.header-context-menu-background {
  position: absolute;
  top: 0px;
  left: 0px;
  width: 100%;
  height: 100%;
}

.header-context-menu {
  background-color: white;
  position: absolute;
  right: 40px;
  top: 35px;
  min-height: 50px;
  display: flex;
  flex-direction: column;
}
.header-context-menu button {
  border-radius: 0px;
  border: solid 1px #9a9a9a;
  font-size: 1em;
  padding: 0.5em 1.5em;
}
</style>
