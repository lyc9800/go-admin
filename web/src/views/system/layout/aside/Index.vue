<template>
    <el-menu
        color="white"
        text-color="#2d3436"
        router
        :default-active="route.path"
        :unique-opened="false"
        :default-openeds="[route.path]"
        class="el-menu-vertical-demo modern-sidebar"
        :collapse="isCollapse"
        :collapse-transition="true"
    >
        <!-- logo start-->
        <div class="imagBox" v-if="!isCollapse">
            <img src="@/assets/logo03.svg" alt="后台管理系统">
        </div>
        <!-- logo end-->
        <!-- 遍历菜单 start-->
        <template v-for="(v,index) in menuData" :key="index">
            <!-- 如果菜单有子菜单，则循环子菜单 start-->
            <el-sub-menu v-if="v.sub_menus.length>0" :index="index">
                <template #title>
                    <el-icon class="menu-icon">
                        <component :is="v.web_icon"></component>
                    </el-icon>
                    <span class="menu-text">{{ v.name }}</span>
                </template>
                <el-menu-item v-for="child in v.sub_menus" :key="child.path"
                    :index="child.path">
                    <el-icon>
                        <component :is="child.web_icon"></component>
                    </el-icon>
                    {{ child.name }}
                </el-menu-item>
            </el-sub-menu>
            <!-- 如果菜单有子菜单，则循环子菜单 end-->

            <!-- 没有子菜单 start-->
            <el-menu-item v-if="v.sub_menus.length==0" :key="v.path"
                :index="v.path">
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
/* 全局重置：清除 el-menu 默认背景色 */
.el-menu-vertical-demo {
    background-color: #fff !important; /* 强制白色背景 */
    border-right: none !important; /* 移除默认右边框 */
}

.imagBox {
    width: 100%;
    height: 70px; 
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 8px 0 5px 0; /* 上边距减少，下边距增加，让logo向上移动 */
    box-sizing: border-box;
    background-color: #fff; /* 与菜单背景一致 */
    margin-bottom: 0; /* 移除底部 margin，避免间隙 */
    border-bottom: 1px solid #f0f0f0; /* 用细线分隔，更现代 */
}

.imagBox img {
    width: 90%; 
    max-width: 180px; 
    height: auto;
    object-fit: contain;
    display: block;
    transition: transform 0.2s ease; 
}

.imagBox img:hover {
    transform: scale(1.05); 
}

/* 菜单容器：弹性布局，让 Logo 和菜单项垂直分布 */
.el-menu {
    height: 100%;
    border: 0px;
    width: 240px; 
    background: #fff; /* 强制白色背景 */
    overflow: hidden;
    display: flex;
    flex-direction: column;
    padding-top: 0; /* 移除顶部内边距 */
}

/* 选中菜单样式：更现代的渐变 + 细边框 */
:deep(.el-menu-item.is-active) {
    color: #333;  
    background: linear-gradient(135deg, #fcd34d 0%, #f59e0b 100%);
    font-weight: 500;
    border-left: 3px solid #4a6fa5; /* 科技蓝边框，更现代 */
}

/* 菜单项间距优化 */
.el-menu-item, .el-sub-menu__title {
    padding-left: 24px !important;
    height: 44px !important; 
    font-size: 14px;
    line-height: 44px !important; /* 垂直居中 */
}

/* 子菜单标题背景色：与菜单一致，避免灰色 */
:deep(.el-sub-menu__title) {
    background-color: #fff !important; 
    color: #2d3436; /* 深色文字，更清晰 */
}

/* 子菜单展开时的背景色（可选） */
:deep(.el-sub-menu.is-opened .el-sub-menu__title) {
    background-color: #f8fafc !important; 
}

/* 图标和文字的间距 */
.menu-icon {
    margin-right: 8px;
    font-size: 16px;
}

/* 菜单文字垂直居中 */
.menu-text {
    vertical-align: middle;
}
</style>