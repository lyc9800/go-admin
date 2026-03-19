<template>
    <div class="icon-picker">
        <div class="icon-header">
            <span class="icon-total">共 {{ total }} 个图标</span>
            <el-input 
                class="icon-search" 
                size="small" 
                v-model="keyword" 
                clearable
                placeholder="搜索图标"
                :prefix-icon="Search"
            />
        </div>

        <el-scrollbar height="300px" class="icon-scrollbar">
            <div class="icon-grid">
                <div 
                    class="icon-card" 
                    v-for="(item, index) in list" 
                    :key="index"
                    @click="onChangeIcon(item?.name || '')"
                    :title="item?.name"
                >
                    <div class="icon-wrapper">
                        <el-icon size="22" v-if="item">
                            <component :is="item" />
                        </el-icon>
                    </div>
                    <div class="icon-label">{{ item?.name || '' }}</div>
                </div>
            </div>
        </el-scrollbar>

        <div class="icon-footer">
            <el-pagination
                small
                :total="total"
                :page-size="pageSize"
                :current-page="currentPage"
                layout="prev, pager, next"
                background
                @current-change="onCurrentChange"
            />
        </div>
    </div>
</template>

<script lang='ts' setup>
import { computed, ref, watch } from 'vue'
import * as Icons from '@element-plus/icons-vue'
import { Search } from '@element-plus/icons-vue'

const emit = defineEmits(['onChangeIcon'])

const iconList = ref([
    Icons.Edit, Icons.Delete, Icons.Plus, Icons.Search, Icons.Refresh,
    Icons.Download, Icons.Upload, Icons.CopyDocument, Icons.Share, Icons.Scissors,
    Icons.User, Icons.UserFilled, Icons.Lock, Icons.Unlock, Icons.Key,
    Icons.Setting, Icons.HomeFilled, Icons.Menu, Icons.Fold, Icons.Expand,
    Icons.Bell, Icons.Message, Icons.Comment, Icons.ChatDotRound, Icons.ChatLineRound,
    Icons.Document, Icons.Files, Icons.Folder, Icons.DataLine, Icons.PieChart
])

const pageSize = ref(24)
const currentPage = ref(1)
const keyword = ref('')

const filteredList = computed(() => {
    if (!keyword.value) return iconList.value
    const kw = keyword.value.toLowerCase()
    return iconList.value.filter(item => 
        item?.name?.toLowerCase().includes(kw)
    )
})

const list = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    return filteredList.value.slice(start, start + pageSize.value)
})

const total = computed(() => filteredList.value.length)

watch(keyword, () => {
    currentPage.value = 1
})

const onCurrentChange = (page: number) => {
    currentPage.value = page
}

const onChangeIcon = (iconName: string) => {
    if (iconName) {
        emit('onChangeIcon', iconName)
    }
}
</script>

<style scoped>
/* 最简单粗暴的修复 */
.container ul {
    display: grid !important;
    grid-template-columns: repeat(6, 1fr) !important;
    gap: 8px !important;
    padding: 10px !important;
}

.container li {
    width: 100% !important;
    height: 70px !important;
    margin: 0 !important;
    padding: 8px !important;
    border: 1px solid #e4e7ed !important;
    border-radius: 6px !important;
    list-style: none !important;
    display: flex !important;
    flex-direction: column !important;
    align-items: center !important;
    justify-content: center !important;
}

.container li:hover {
    border-color: #e99d53 !important;
}

.container li svg {
    width: 24px !important;
    height: 24px !important;
    margin-bottom: 4px !important;
}

.container li span {
    font-size: 11px !important;
    color: #666 !important;
}

.search-box {
    text-align: center;
    margin: 10px 0 20px;
}

.search {
    width: 280px;
}

.el-pagination {
    margin-top: 20px;
    display: flex;
    justify-content: center;
}
</style>