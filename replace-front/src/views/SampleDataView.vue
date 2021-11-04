<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs

import { computed } from 'vue'
import { CreateUserRepository } from '../infrastructures/offline_user_repository_impl'

const userRepo = CreateUserRepository()
const displayName = computed(() => userRepo.Get('').displayName)
</script>

<template>
  <div class="sample">
    <h1>hello, {{ displayName }}</h1>
  </div>
</template>
