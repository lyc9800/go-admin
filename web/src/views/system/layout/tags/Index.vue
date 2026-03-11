<template>
    <div class="main-tabs-view">
        <div class="tabs-view">
            <el-tabs v-model="activeTabsValue" type="card" @tab-click="tabClick" @tab-remove="removeTab">
                <el-tab-pane v-for="item in visitedViews"
                :key="item.path"
                :path="item.path"
                :label="item.title"
                :name="item.path"
                :closable="!(item.meta&&item.meta.affix)"  
                >
                <template #label>
                    <el-icon class="tabs-icon" v-if="item.icon">
                        <component :is="item.icon"></component>
                    </el-icon>
                    {{item.title}}
                </template>
                </el-tab-pane>
            </el-tabs>
        </div>
        <div class="right-tbn">
            <MoreButton/>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref,computed,watch,onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {useTagsViewsStore}  from '@/store/modules/tagsViews'
import type  { TabsPaneContext } from 'element-plus'
import MoreButton from './components/MoreButton.vue'
const tagsViewsStore = useTagsViewsStore()
const route = useRoute()
const router = useRouter()
const visitedViews = computed(()=>tagsViewsStore.visitedViews)
const activeTabsValue = computed({
    get:()=>{
        return tagsViewsStore.activeTabsValue
    },
    set:val=>{
        tagsViewsStore.setTabsMenuValue(val)
    }
})
// 显示上一个或者下一个tabs标签
function toLastView(acticeTabPath:any){
    // 当前tabs标签索引
    let index=visitedViews.value.findIndex(item=>item.path===acticeTabPath)
    // 获取下一个tab索引或者上一个索引
    const nextTab = visitedViews.value[index+1] || visitedViews.value[index-1]
    if(!nextTab) return
    router.push(nextTab.path)
    // 新增tabs标签的函数
    tagsViewsStore.setTabsMenuValue(nextTab.path)
}
// 点击tabs标签事件
const tabClick=(tabItem:TabsPaneContext)=>{
    let path=tabItem.props.name as string
    router.push(path)
    tagsViewsStore.setTabsMenuValue(path) 
}
// 添加tabs标签函数
const addTabs=()=>{
    const {name}=route
    if(name==='Login')return
    if(name){
        tagsViewsStore.addView(route)

    }
    return false
}
// 判断是否删除的是否为当前标签
const isActive=(path:any)=>{
    return path===route.path
}
// 删除tabs标签函数
const removeTab=async (activeTabPath:string)=>{
    if(isActive(activeTabPath)){
        toLastView(activeTabPath)
    }
    await tagsViewsStore.delView(activeTabPath)
}
onMounted(()=>{
    addTabs()
})
watch(route,()=>{
    addTabs()
})
</script>
<style scoped>
.main-tabs-view {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-left: 0px;
    padding-right: 0px;
    background: white;
    height: 40px;
    border-bottom: 1px solid #f0f0f0;
}

.tabs-view {
    flex: 1;
    overflow: hidden;
    box-sizing: border-box;
    background: white;
    height: 100%;
    display: flex;
    align-items: center;
    position: relative; /* 添加定位上下文 */
}

/* 整个tabs组件 */
:deep(.el-tabs) {
    border-top: 1px solid #e99d53;
    background-color: white;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
}

/* 标签头部 */
:deep(.el-tabs__header) {
    background-color: white;
    border-bottom: none;
    margin: 0;
    height: 100%;
    padding: 0;
    position: relative;
    z-index: 1;
}

/* 滚动容器 */
:deep(.el-tabs__nav-wrap.is-top) {
    background-color: white;
    height: 100%;
    padding: 0;
    margin: 0;
}

:deep(.el-tabs__nav-scroll) {
    background-color: white;
    height: 100%;
    padding: 0;
    overflow: hidden;
}

:deep(.el-tabs__nav) {
    border: none;
    background-color: white;
    height: 100%;
    display: flex;
    align-items: center;
}

/* 单个标签项 */
:deep(.el-tabs .el-tabs__header .el-tabs__item) {
    border: none;
    color: #cccccc;
    background-color: white;
    padding: 0 16px;
    height: 100%;
    line-height: 40px;
    display: flex;
    align-items: center;
}

/* 激活状态标签 */
:deep(.el-tabs .el-tabs__header .el-tabs__item.is-active) {
    border-bottom: 2px solid #e99d53;
    color: #e99d53;
    background-color: white;
}

/* 确保"更多"按钮区域背景为白色 */
.right-tbn {
    background-color: white;
    height: 100%;
    display: flex;
    align-items: center;
    padding: 0 10px;
    z-index: 2;
    position: relative;
}

/* 隐藏默认的tabs内容区域，避免灰色背景 */
:deep(.el-tabs__content) {
    background-color: white !important;
    padding: 0 !important;
    margin: 0 !important;
    height: 0 !important;
    overflow: hidden !important;
}
</style>