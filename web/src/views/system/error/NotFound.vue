<template>
  <div class="not-found">
    <h2>404 - 组件未找到</h2>
    <p>路径: {{ $route.path }}</p>
  </div>
</template>

<script setup lang="ts">
// 简单提示组件
</script>

<style scoped>
.not-found {
  padding: 20px;
  text-align: center;
  color: #999;
}
</style>