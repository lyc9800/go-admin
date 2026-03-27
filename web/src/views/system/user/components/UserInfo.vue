<template>
  <el-popover :width="320" placement="bottom-end" trigger="hover" class="user-popover">
    <template #reference>
      <el-link :underline="false" class="user-info-link">
        <img v-if="userInfo?.avatar" :src="url + 'uploadFile/' + userInfo.avatar" class="user-avatar">
        <img v-else src="@/assets/default.png" class="user-avatar">
        <span class="user-name">{{ userInfo?.userName }}</span>
        <el-icon class="dropdown-icon"><ArrowDown /></el-icon>
      </el-link>
    </template>
    
    <template #default>
      <div class="user-popover-content">
        <!-- 头部信息卡片 -->
        <div class="user-header">
          <img v-if="userInfo?.avatar" :src="url + 'uploadFile/' + userInfo.avatar" class="user-avatar-large">
          <img v-else src="@/assets/default.png" class="user-avatar-large">
          <div class="user-basic-info">
            <div class="user-display-name">{{ userInfo?.userName || '未知用户' }}</div>
            <div class="user-role" v-if="userInfo?.role_name ">{{ userInfo.role_name  }}</div>
            <div class="user-role" v-else>普通用户</div>
          </div>
        </div>
        
        <!-- 详细信息列表 -->
        <div class="user-details">
          <div class="detail-item">
            <el-icon class="detail-icon"><Message /></el-icon>
            <div class="detail-content">
              <div class="detail-label">邮箱</div>
              <div class="detail-value">{{ userInfo?.email || '未设置' }}</div>
            </div>
          </div>
          
          <div class="detail-item">
            <el-icon class="detail-icon"><Calendar /></el-icon>
            <div class="detail-content">
              <div class="detail-label">注册时间</div>
              <div class="detail-value">{{ formatTime(userInfo?.CreatedAt, 'yyyy-MM-dd') }}</div>
            </div>
          </div>
          
          <div class="detail-item">
            <el-icon class="detail-icon"><Timer /></el-icon>
            <div class="detail-content">
              <div class="detail-label">注册天数</div>
              <div class="detail-value">{{ calculateDays(userInfo?.CreatedAt) }} 天</div>
            </div>
          </div>
        </div>
        
        <!-- 操作按钮 -->
        <div class="user-actions">
          <el-button type="primary" plain size="small" @click="goToProfile" class="action-btn">
            <el-icon><User /></el-icon>
            个人中心
          </el-button>
          <el-button type="danger" plain size="small" @click="handleExit" class="action-btn">
            <el-icon><SwitchButton /></el-icon>
            退出登录
          </el-button>
        </div>
      </div>
    </template>
  </el-popover>
</template>

<script setup lang="ts">
import { SwitchButton, Message, Calendar, Timer, User, ArrowDown } from '@element-plus/icons-vue'
import { useUserStore } from '@/store/modules/user'
import { formatTime, calculateDays } from '@/utils/date';
import { useRouter } from 'vue-router';
import { ElMessageBox } from 'element-plus'

// 获取登陆用户信息
const userStore = useUserStore()
const userInfo = userStore.userInfo
const router = useRouter()

// 服务器路径
const url = import.meta.env.VITE_BASE_URL
console.log('userInfo 完整数据:', userInfo)
console.log('role_name:', userInfo?.role_name)
// 退出系统
const exit = () => {
  window.localStorage.removeItem("userStore")
  window.localStorage.removeItem("menuState")
  router.push('/login')
}

// 处理退出登录（带确认框）
const handleExit = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要退出登录吗？',
      '确认退出',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        closeOnClickModal: false
      }
    )
    exit()
  } catch {
    // 用户取消了退出
  }
}

// 跳转到个人中心
const goToProfile = () => {
  router.push('/profile')
}
</script>

<style scoped>
/* 用户信息链接样式 */
.user-info-link {
  display: flex;
  align-items: center;
  gap: 12px;
  color: #606266;
  padding: 8px 10px;
  border-radius: 6px;
  transition: all 0.3s ease;
  cursor: pointer;
}

.user-info-link:hover {
  background: #f5f7fa;
  color: #409eff;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e4e7ed;
  transition: all 0.3s ease;
}

.user-info-link:hover .user-avatar {
  border-color: #409eff;
  transform: scale(1.05);
}

.user-name {
  font-weight: 500;
  font-size: 14px;
  margin-left: 4px;
}

.dropdown-icon {
  font-size: 12px;
  color: #c0c4cc;
  transition: transform 0.3s ease;
  margin-left: 4px;
}

.user-info-link:hover .dropdown-icon {
  color: #409eff;
  transform: rotate(180deg);
}

/* 弹出框内容样式 */
.user-popover-content {
  padding: 0;
}

/* 头部信息卡片 */
.user-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  color: white;
}

.user-avatar-large {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  border: 3px solid rgba(255, 255, 255, 0.3);
  object-fit: cover;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.user-basic-info {
  flex: 1;
}

.user-display-name {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 6px;
}

.user-role {
  font-size: 12px;
  opacity: 0.9;
  background: rgba(255, 255, 255, 0.2);
  padding: 3px 10px;
  border-radius: 12px;
  display: inline-block;
}

/* 详细信息列表 */
.user-details {
  padding: 16px 20px;
  background: white;
}

.detail-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid #f0f0f0;
}

.detail-item:last-child {
  border-bottom: none;
}

.detail-icon {
  width: 20px;
  height: 20px;
  color: #667eea;
  margin-top: 2px;
  flex-shrink: 0;
}

.detail-content {
  flex: 1;
  min-width: 0;
}

.detail-label {
  font-size: 12px;
  color: #909399;
  margin-bottom: 4px;
}

.detail-value {
  font-size: 14px;
  color: #303133;
  font-weight: 500;
  word-break: break-all;
}

/* 操作按钮 */
.user-actions {
  padding: 12px 20px 16px;
  display: flex;
  gap: 12px;
  background: white;
  border-top: 1px solid #f0f0f0;
}

.action-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  border-radius: 6px;
  font-size: 13px;
  transition: all 0.3s ease;
}

.action-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
</style>