<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs

import { computed } from 'vue'
import { CreateUserRepository } from '../infrastructures/offline_user_repository_impl'

const userRepo = CreateUserRepository()
</script>

<template>
  <div class="sample">
    <h1>hello, {{ userRepo.Get('').displayName }}</h1>
  </div>
</template>
