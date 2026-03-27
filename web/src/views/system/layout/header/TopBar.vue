<template>
  <div class="main">
    <div class="tool-left">
      <CollapseIcon/>
      <Hamburger/>
    </div>
    <div class="linkBox">
      <!-- 使用用户信息组件 -->
      <UserInfo/>
    </div>
  </div>
  <!-- tabs 多标签页 start -->
  <TabsView/>
  <!-- tabs 多标签页 end-->
</template>

<script setup lang="ts">
import CollapseIcon from './CollapseIcon.vue';
import Hamburger from './Hamburger.vue';
import TabsView from '@/views/system/layout/tags/Index.vue';
import UserInfo from '@/views/system/user/components/UserInfo.vue';
</script>

<style scoped>
.main {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
  box-shadow: rgb(0 0 0 / 10%) 0px 0px 10px;
  background: white;
}

.linkBox {
  height: 100%;
  display: flex;
  align-items: center;
  float: right;
  gap: 12px;
  padding-right: 10px;
}

.tool-left {
  display: flex;
  align-items: center;
  margin-left: 10px;
}

:deep(.main-tabs-view) {
  background: white !important;
}
</style>