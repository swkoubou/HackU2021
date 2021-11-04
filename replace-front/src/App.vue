<template>
  <div class="p-4">
    <router-link to="/" class="font-semibold text-gray-800">Home</router-link> |
    <router-link to="/about" class="font-semibold text-gray-800"
      >About</router-link
    > |
    <router-link to="/sample" class="font-semibold text-gray-800">Sample</router-link>
  </div>
  <router-view />
</template>

<style scoped>
a.router-link-exact-active {
  color: #42b983;
}
</style>
