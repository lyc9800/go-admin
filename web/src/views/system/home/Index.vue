<template>
  <div class="home-container">
    <!-- 欢迎区域 -->
    <el-row :gutter="16" class="welcome-section">
      <el-col :span="16">
        <h2>👋 欢迎回来，Admin</h2>
        <p class="subtitle">今天是 {{ currentDate }}，系统运行正常</p>
      </el-col>
      <el-col :span="8" class="text-right">
        <el-button type="primary" size="small" @click="refreshData">
          <el-icon><Refresh /></el-icon> 刷新数据
        </el-button>
      </el-col>
    </el-row>

    <!-- 快捷操作 -->
    <el-row :gutter="16" class="quick-actions">
      <el-col :span="24">
        <el-card shadow="hover" header="⚡ 快捷操作">
          <el-space wrap size="small">
            <el-button type="primary" plain size="small" @click="goToPage('pipeline')">
              <el-icon><VideoPlay /></el-icon> 流水线
            </el-button>
            <el-button type="success" plain size="small" @click="goToPage('service')">
              <el-icon><Service /></el-icon> 服务
            </el-button>
            <el-button type="warning" plain size="small" @click="goToPage('alert')">
              <el-icon><Bell /></el-icon> 告警
            </el-button>
            <el-button type="info" plain size="small" @click="goToPage('monitor')">
              <el-icon><Monitor /></el-icon> 监控
            </el-button>
            <el-button type="danger" plain size="small" @click="goToPage('config')">
              <el-icon><Setting /></el-icon> 配置
            </el-button>
          </el-space>
        </el-card>
      </el-col>
    </el-row>

    <!-- 状态卡片 -->
    <el-row :gutter="16" class="stat-cards">
      <el-col :xs="12" :sm="6" v-for="item in statCards" :key="item.title">
        <el-card shadow="hover" class="stat-card">
          <div class="card-icon" :style="{ background: item.color }">
            <el-icon :size="22"><component :is="item.icon" /></el-icon>
          </div>
          <div class="card-content">
            <div class="card-value">{{ item.value }}</div>
            <div class="card-title">{{ item.title }}</div>
            <div class="card-trend" :class="item.trend > 0 ? 'up' : 'down'">
              {{ item.trend > 0 ? '↑' : '↓' }} {{ Math.abs(item.trend) }}%
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 图表区域 -->
    <el-row :gutter="16" class="chart-section">
      <el-col :span="12">
        <el-card shadow="hover" header="📊 CPU 使用率" style="height: 200px;">
          <div ref="cpuChartRef" class="chart-box"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover" header="📊 内存使用率" style="height: 200px;">
          <div ref="memChartRef" class="chart-box"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import * as echarts from 'echarts'
import {
  Refresh,
  Monitor,
  Service,
  Bell,
  VideoPlay,
  Setting
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const router = useRouter()

// 当前日期
const currentDate = new Date().toLocaleDateString('zh-CN', {
  month: 'long',
  day: 'numeric',
  weekday: 'short'
})

// 状态卡片数据
const statCards = ref([
  { title: '服务总数', value: 128, trend: 5, icon: 'Grid', color: '#409eff' },
  { title: '正常运行', value: 120, trend: 2, icon: 'CircleCheck', color: '#67c23a' },
  { title: '告警数量', value: 8, trend: -3, icon: 'Warning', color: '#e6a23c' },
  { title: '成功率', value: '96.5%', trend: 1.2, icon: 'TrendCharts', color: '#f56c6c' }
])

// 图表引用
const cpuChartRef = ref<HTMLElement>()
const memChartRef = ref<HTMLElement>()
let cpuChart: echarts.ECharts | null = null
let memChart: echarts.ECharts | null = null

// 初始化图表
const initCharts = () => {
  if (cpuChartRef.value) {
    cpuChart = echarts.init(cpuChartRef.value)
    cpuChart.setOption({
      grid: { top: 10, right: 10, bottom: 10, left: 10 },
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: ['0点', '4点', '8点', '12点', '16点', '20点', '24点'], axisLabel: { fontSize: 10 } },
      yAxis: { type: 'value', max: 100, axisLabel: { fontSize: 10 } },
      series: [{ data: [45, 52, 38, 65, 58, 72, 48], type: 'line', smooth: true, symbol: 'none', lineStyle: { color: '#409eff' } }]
    })
  }
  if (memChartRef.value) {
    memChart = echarts.init(memChartRef.value)
    memChart.setOption({
      grid: { top: 10, right: 10, bottom: 10, left: 10 },
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: ['0点', '4点', '8点', '12点', '16点', '20点', '24点'], axisLabel: { fontSize: 10 } },
      yAxis: { type: 'value', max: 100, axisLabel: { fontSize: 10 } },
      series: [{ data: [62, 58, 55, 70, 68, 75, 60], type: 'line', smooth: true, symbol: 'none', lineStyle: { color: '#67c23a' } }]
    })
  }
}

// 窗口大小变化时重绘图表
const handleResize = () => {
  cpuChart?.resize()
  memChart?.resize()
}

// 刷新数据
const refreshData = () => {
  ElMessage.success('数据已刷新')
}

// 页面跳转
const goToPage = (page: string) => {
  ElMessage.info(`跳转到 ${page} 页面`)
}

onMounted(() => {
  initCharts()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  cpuChart?.dispose()
  memChart?.dispose()
})
</script>

<style scoped>
.home-container {
  padding: 12px;
  background: #f5f7fa;
  min-height: calc(100vh - 84px);
  box-sizing: border-box;
  overflow: hidden;
}

.welcome-section {
  margin-bottom: 12px;
}

.welcome-section h2 {
  margin: 0 0 4px 0;
  font-size: 18px;
  color: #303133;
}

.subtitle {
  color: #909399;
  margin: 0;
  font-size: 12px;
}

.text-right {
  text-align: right;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

/* 快捷操作 */
.quick-actions {
  margin-bottom: 12px;
}

.quick-actions .el-card__body {
  padding: 10px 12px;
}

.quick-actions .el-space {
  gap: 4px !important;
}

/* 状态卡片 */
.stat-cards {
  margin-bottom: 12px;
}

.stat-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 12px 8px;
  text-align: center;
  height: 100%;
  min-height: 110px;
}

.card-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  margin-right: 0;
  margin-bottom: 8px;
}

.card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.card-value {
  font-size: 22px;
  font-weight: bold;
  color: #303133;
  line-height: 1.2;
}

.card-title {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

.card-trend {
  font-size: 10px;
  margin-top: 2px;
}

.card-trend.up {
  color: #67c23a;
}

.card-trend.down {
  color: #f56c6c;
}

/* 图表区域 */
.chart-section {
  margin-bottom: 0;
}

.chart-section .el-card {
  height: 200px;
}

.chart-section .el-card__body {
  padding: 10px 12px;
  height: calc(100% - 45px);
}

.chart-box {
  height: 100%;
  width: 100%;
}
</style>