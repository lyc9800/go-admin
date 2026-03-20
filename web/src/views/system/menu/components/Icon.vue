<template>
  <div class="icon-picker">
    <!-- 图标网格：自适应填充，分页自动推到底部 -->
    <div class="icon-grid">
      <div 
        class="icon-card" 
        v-for="(item, index) in list" 
        :key="index"
        @click="onChangeIcon(item?.name || '')"
      >
        <div class="icon-wrapper">
          <el-icon v-if="item" :size="18">
            <component :is="item" />
          </el-icon>
          <span v-else>?</span>
        </div>
      </div>
    </div>

    <!-- 分页：自动推到底部 -->
    <div class="icon-footer">
      <el-pagination
        :total="total"
        :page-size="pageSize"
        :current-page="currentPage"
        layout="prev, pager, next"
        small
        @current-change="onCurrentChange"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import * as Icons from '@element-plus/icons-vue'

const emit = defineEmits(['onChangeIcon'])

const iconList = ref([
  Icons.Edit, Icons.Delete, Icons.Plus, Icons.Search, Icons.Refresh,
  Icons.Download, Icons.Upload, Icons.CopyDocument, Icons.Share, Icons.Scissor,
  Icons.User, Icons.UserFilled, Icons.Lock, Icons.Unlock, Icons.Key,
  Icons.Setting, Icons.HomeFilled, Icons.Menu, Icons.Fold, Icons.Expand,
  Icons.Bell, Icons.Message, Icons.Comment, Icons.ChatDotRound, Icons.ChatLineRound,
  Icons.Document, Icons.Files, Icons.Folder, Icons.DataLine, Icons.PieChart,
  Icons.Picture, Icons.Camera, Icons.Star, Icons.Location,
  Icons.House, Icons.Trophy, Icons.Coin, Icons.Goods, Icons.Wallet,
  Icons.Ship, Icons.Bicycle, 
  Icons.Timer, Icons.AlarmClock, Icons.Calendar, Icons.Clock, Icons.Watch,
  Icons.Phone, Icons.VideoCamera, Icons.Microphone, Icons.Headset,
  Icons.Printer, Icons.Cpu
])

const pageSize = ref(20)
const currentPage = ref(1)

const list = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  return iconList.value.slice(start, start + pageSize.value)
})

const total = computed(() => iconList.value.length)

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
/* 整体容器：弹窗样式 */
.icon-picker {
  width: 100%;
  max-width: 500px;
  margin: 0 auto;
  padding: 16px;
  background: #fff;
  border-radius: 8px;
  box-shadow: none;
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  overflow: hidden;
  height: 340px; /* 固定高度，让分页自动推到底部 */
}

/* 图标网格：自适应填充，分页自动推到底部 */
.icon-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 0;
  width: 100%;
  box-sizing: border-box;
  flex: 1; /* 占满剩余空间 */
  overflow: hidden;
  /* 移除边框和阴影，让图标区域和分页无缝衔接 */
  border: none;
  border-radius: 0;
  background: #fff;
}

/* 图标卡片：紧凑大小，无多余边框 */
.icon-card {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 65px;
  border-right: 1px solid #f5f5f5;
  border-bottom: 1px solid #f5f5f5;
  cursor: pointer;
  transition: all 0.2s;
  box-sizing: border-box;
  background: white;
  position: relative;
}

/* 最后一列移除右侧边框 */
.icon-card:nth-child(5n) {
  border-right: none;
}

/* 最后一行移除底部边框 */
.icon-card:nth-last-child(-n+5) {
  border-bottom: none;
}

/* 悬停效果 */
.icon-card:hover {
  background: #fffaf5;
  z-index: 1;
  box-shadow: 0 2px 8px rgba(233, 157, 83, 0.2);
}

/* 图标样式 */
.icon-card .el-icon {
  width: 20px;
  height: 20px;
  color: #606266;
  transition: all 0.2s;
}

.icon-card:hover .el-icon {
  transform: scale(1.1);
  color: #e99d53;
}

/* 分页区域：固定在底部中间，无阴影 */
.icon-footer {
  display: flex;
  justify-content: center;
  padding: 6px 0;
  width: 100%;
  margin-top: auto; /* 自动推到最底部 */
  flex-shrink: 0; /* 防止被压缩 */
  /* 移除可能的阴影 */
  box-shadow: none;
}

/* 分页样式 */
.el-pagination {
  margin: 0;
  --el-pagination-button-bg-color: #fff;
  --el-pagination-bg-color: #fff;
  --el-pagination-hover-color: #e99d53;
  /* 移除分页的默认阴影 */
  box-shadow: none;
}
</style>