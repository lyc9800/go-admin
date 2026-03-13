<template>
    <el-menu color="white" text-color="#67879b" router
    :default-active="route.path"
    :unique-opened="false" :default-openeds="[route.path]"
    class="el-menu-vertical-demo"
    :collapse="isCollapse"
    :collapse-transition="true"
    >
    <!-- logo start-->
    <div class="imagBox" v-if="!isCollapse">
        <img src="@/assets/logo03.png" alt="后台管理系统">
    </div>
    <!-- logo end-->
    <!-- 遍历菜单 start-->
    <template v-for="(v,index) in menuData" :key="index">
        <!-- 如果菜单有子菜单，则循环子菜单 start-->
         <el-sub-menu v-if="v.sub_menus.length>0" :index="index">
            <template #title>
                <el-icon>
                    <component :is="v.web_icon"></component>
                </el-icon>
                <span>{{ v.name }}</span>
            </template>
            <el-menu-item v-for="child in v.sub_menus" :key="child.path" :index="child.path">
                <el-icon>
                    <component :is="child.web_icon"></component>
                </el-icon>
                {{ child.name }}
            </el-menu-item>
         </el-sub-menu>
        <!-- 如果菜单有子菜单，则循环子菜单 end-->

        <!-- 没有子菜单 start-->
         <el-menu-item v-if="v.sub_menus.length==0" :key="v.path" :index="v.path">
            <el-icon>
                <component :is="v.web_icon"></component>
            </el-icon>
            <span>{{ v.name }}</span>
         </el-menu-item>
        <!-- 没有子菜单 end-->
    </template>
    <!-- 遍历菜单 end-->
    </el-menu>
</template>

<script setup lang="ts">
import {useRoute} from 'vue-router'
import {ref,computed} from 'vue' 
import { useMenuStore } from '@/store/modules/menu';
import { useSettingStore } from '@/store/modules/setting';
const route = useRoute()
// 获取菜单数据
const {routers} = useMenuStore()
const menuData = ref()
menuData.value=routers
const settingStore = useSettingStore()
const isCollapse = computed(()=>settingStore.isCollapse)
</script>

<style scoped>
.imagBox {
    width: 100%;
    height: 70px; /* 稍微增加高度，让 Logo 更有呼吸感 */
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 10px 0;
    box-sizing: border-box;
    background-color: white; /* 确保背景色一致 */
    /* border-bottom: 1px solid #f0f0f0;  */
}

.imagBox img {
    width: 90%; 
    max-width: 200px; 
    height: 80px; 
    object-fit: contain;
    display: block;
    transition: transform 0.2s ease; 

}

.imagBox img:hover {
    transform: scale(1.05); /* 可选：鼠标悬停轻微放大，增强交互 */
}

.el-menu {
    height: 100%;
    border: 0px;
    width: 240px; /* 菜单总宽度，可按需调整 */
    /* background-color: white; */
    background: #f5f7fa;
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

/* 选中菜单时候的颜色 */
/* :deep(.el-menu-item.is-active) {
    color: white;
    background: linear-gradient(to right, #1890ff, #36cfc9);
    font-weight: bold;
} */
 :deep(.el-menu-item.is-active) {
    color: #333;  /* 文字颜色改为深色，在米黄背景上更清晰 */
    background: linear-gradient(to right, #e6f7ff, #b3e5fc);
    /* 从亮米黄色渐变到小麦色 */
    font-weight: bold;
    border-left: 3px solid #d4b483;  /* 添加左侧边框增加层次感 */
}

/* 菜单项间距优化 */
.el-menu-item, .el-sub-menu__title {
    padding-left: 20px !important;
    font-size: 14px;
}

</style>