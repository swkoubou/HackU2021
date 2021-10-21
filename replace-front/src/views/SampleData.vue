<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs

import {computed} from 'vue';
import { OfflineUserController } from '../repo/user'
const userController = computed(() => new OfflineUserController());

</script>

<template>
    <div class="sample">
        <h1>hello, {{ userController.getUserWithId("a").displayName }}</h1>
    </div>
</template>

