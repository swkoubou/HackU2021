import { InjectionKey } from 'vue'
import { createStore, Store, useStore as baseUseStore } from 'vuex'

// stateの型定義
type State = {
  count: number
}

// storeをprovide/injectするためのキー
export const key: InjectionKey<Store<State>> = Symbol()

// store本体
export const store = createStore<State>({
  state: {
    count: 0,
  },
  getters: {
    // 型が定義されないので使わない
    // stateを直接参照
    // computed経由でReadOnly<T>にして使う事
    //ex) const count = computed((): Readonly<number> => store.state.count)
  },
  actions: {
    increment({ commit }) {
      commit('increment')
    },
  },
  mutations: {
    increment(state) {
      state.count++
    },
  },
})

// useStoreを使う時にキーの指定を省略するためのラッパー関数
export const useStore = () => {
  return baseUseStore(key)
}
