<template>
    <div id="ebg"></div>
    <div class="common-layout">
<el-container>
    <!--左侧菜单 start-->
        <el-aside width="auto">
            <AsideMenu/>
        </el-aside>
    <!--左侧菜单 end-->
<!--顶部导航栏 start-->
<el-header>这里是顶部内容</el-header>
<!--顶部导航栏 end-->
<!--主要内容 start-->
<el-main>
    <router-view>这里放主要内容</router-view>
</el-main>
<!--主要内容 end-->
</el-container>
    </div>
</template>
<script setup lang="ts">
// 引入左侧菜单组件
import AsideMenu from './aside/Index.vue'
</script>
<style  scoped>
/* 顶部导航栏样式 */
.el-header{
    --el-header-padding:0px;
    --el-header-height:auto;
    height: 80px;
    background: #f3f3f3;
}
.el-main{
    background: #f3f3f3;
}
.el-container{
    height: 100%;
}
.common-layout{
    width: 99vw;
    height: 98vh;
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
    background-color: white;
    overflow: hidden;
}

/* 背景图片样式 */
#ebg{
    width: 100%;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    padding: 10px;
    background-image: url("../../assets/system-bg.jpg");
    background-size: cover;
    background-position: center center;
    background-repeat: no-repeat;
}
</style>
