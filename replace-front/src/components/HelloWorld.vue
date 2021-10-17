<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '../store'
defineProps<{ msg: string }>()

const store = useStore()

const count = computed((): Readonly<number> => store.state.count)
const increment = () => store.dispatch('increment')
</script>

<template>
  <h1>{{ msg }}</h1>

  <p>
    Recommended IDE setup:
    <a href="https://code.visualstudio.com/" target="_blank" class="a-tag"
      >VSCode</a
    >
    +
    <a
      href="https://github.com/johnsoncodehk/volar"
      target="_blank"
      class="a-tag"
      >Volar</a
    >
  </p>

  <p>
    See
    <code class="code">README.md</code>
    for more information.
  </p>

  <p>
    <a
      href="https://vitejs.dev/guide/features.html"
      target="_blank"
      class="a-tag"
    >
      Vite Docs
    </a>
    |
    <a href="https://v3.vuejs.org/" target="_blank" class="a-tag">Vue 3 Docs</a>
  </p>

  <button type="button" @click="increment" class="button">
    count is: {{ count }}
  </button>
  <p>
    Edit
    <code class="code">components/HelloWorld.vue</code>
    to test hot module replacement.
  </p>
</template>
