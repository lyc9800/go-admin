<template>
  <div class="error-404">
    <h1>404</h1>
    <p>页面不存在</p>
    <button @click="goHome">返回首页</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const goHome = () => {
  router.push('/home')
}
</script>

<style scoped>
.error-404 {
  text-align: center;
  padding: 50px;
}
</style>