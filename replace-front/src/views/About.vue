<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '../store'
defineProps<{ msg: string }>()

const store = useStore()

const count = computed((): Readonly<number> => store.state.count)
</script>

<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div>count is:{{ count }}</div>
  </div>
</template>
