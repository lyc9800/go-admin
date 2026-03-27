<template>
    <el-dropdown trigger="hover">
        <el-button size="small" color="#e99d53" class="more-btn">
            <span class="more">更多</span>
            <el-icon><arrow-down/></el-icon>
        </el-button>
        <template #dropdown>
            <el-dropdown-menu>
                <el-dropdown-item @click="closeOtherTab">
                    <el-icon size="14"><Close/></el-icon>关闭其他
                </el-dropdown-item>
                <el-dropdown-item @click="closeAllTab">
                    <el-icon size="14"><FolderDelete/></el-icon>关闭所有
                </el-dropdown-item>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useSettingStore } from '@/store/modules/setting';
import { useTagsViewsStore } from '@/store/modules/tagsViews';
import { useRoute, useRouter } from 'vue-router';

const settingStore = useSettingStore()
const route = useRoute()
const router = useRouter()
const tagsViewsStore = useTagsViewsStore()
const visitedViews = computed(() => tagsViewsStore.visitedViews)

// 关闭其他标签
const closeOtherTab = async () => {
   await tagsViewsStore.delOtherViews(route.path)
}

// 关闭所有标签
const closeAllTab = async () => {
    await tagsViewsStore.delAllViews()
    tagsViewsStore.goHome()
}
</script>

<style scoped>
.more-btn {
    margin-right: 20px;
}

.more {
    color: white;
    margin-right: 4px;  /* ✅ 新增：文字右侧间距 */
}

.more-icon {
    margin-left: 2px;  /* ✅ 新增：图标左侧间距 */
    color: white;      /* ✅ 新增：图标颜色与文字一致 */
}
</style>