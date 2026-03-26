<template>
  <div class="email-change-container">
    <!-- 当前邮箱信息 -->
    <div class="current-email-section">
      <div class="info-header">
        <el-icon><Message /></el-icon>
        <span class="title">绑定邮箱</span>
      </div>
      <div class="info-content">
        <div class="current-email-display">
          <span class="label">当前绑定邮箱：</span>
          <span class="value">{{ userInfo.email }}</span>
        </div>
        <p class="hint-text">
          <el-icon><Warning /></el-icon>
          邮箱是登录验证的重要方式，请谨慎操作
        </p>
      </div>
    </div>

    <!-- 更换邮箱表单区域 -->
    <div class="change-form-section">
      <!-- 初始状态：显示更换按钮 -->
      <template v-if="!isChanging">
        <el-button 
          type="primary" 
          plain 
          @click="startChangeEmail"
          class="change-btn"
        >
          <el-icon><Refresh /></el-icon>
          更换绑定邮箱
        </el-button>
      </template>

      <!-- 更换流程表单 -->
      <template v-else>
        <div class="form-container">
          <!-- 步骤指示器 -->
          <div class="steps">
            <div class="step-item" :class="{ active: currentStep === 1, completed: currentStep > 1 }">
              <div class="step-number">1</div>
              <div class="step-label">输入新邮箱</div>
            </div>
            <div class="step-line"></div>
            <div class="step-item" :class="{ active: currentStep === 2, completed: currentStep > 2 }">
              <div class="step-number">2</div>
              <div class="step-label">验证邮箱</div>
            </div>
            <div class="step-line"></div>
            <div class="step-item" :class="{ active: currentStep === 3 }">
              <div class="step-number">3</div>
              <div class="step-label">完成更换</div>
            </div>
          </div>

          <!-- 步骤 1：输入新邮箱 -->
          <div v-show="currentStep === 1" class="step-content step-1">
            <div class="form-group">
              <label class="form-label">新邮箱地址</label>
              <div class="input-with-action">
                <el-input
                  v-model="form.email"
                  placeholder="请输入新的邮箱地址"
                  size="large"
                  @blur="validateEmail"
                  @keyup.enter="sendVerificationCode"
                  class="email-input-with-btn"
                >
                  <template #prefix>
                    <el-icon><Message /></el-icon>
                  </template>
                  <template #suffix>
                    <div class="suffix-group">
                      <el-icon 
                        v-if="form.email" 
                        class="clear-btn" 
                        @click="clearEmail"
                      >
                        <CircleClose />
                      </el-icon>
                      <el-button
                        size="small"
                        :disabled="!form.email || !!emailError || sendingCode"
                        @click="sendVerificationCode"
                        :loading="sendingCode"
                        class="send-code-btn"
                      >
                        {{ sendingCode ? '发送中...' : '发送验证码' }}
                      </el-button>
                    </div>
                  </template>
                </el-input>
                <el-button 
                  @click="cancelChange"
                  class="cancel-btn"
                >
                  取消
                </el-button>
              </div>
              <div v-if="emailError" class="error-text">{{ emailError }}</div>
              <div v-else class="hint">请输入您要绑定的新邮箱地址</div>
            </div>
          </div>

          <!-- 步骤 2：输入验证码 -->
          <div v-show="currentStep === 2" class="step-content step-2">
            <div class="form-group">
              <label class="form-label">验证码</label>
              <p class="hint">验证码已发送至邮箱：{{ form.email }}</p>
              <div class="input-with-action">
                <el-input
                  v-model="form.code"
                  placeholder="请输入 6 位验证码"
                  size="large"
                  maxlength="6"
                  class="code-input-with-btn"
                  @keyup.enter="verifyCode"
                >
                  <template #prefix>
                    <el-icon><Lock /></el-icon>
                  </template>
                  <template #suffix>
                    <el-button
                      size="small"
                      plain
                      :disabled="countdown > 0"
                      @click="resendCode"
                      class="resend-code-btn"
                    >
                      {{ countdownText }}
                    </el-button>
                  </template>
                </el-input>
                <!-- 移除了上一步按钮，只保留确定更换按钮 -->
                <el-button 
                  type="primary" 
                  :disabled="!form.code || form.code.length !== 6"
                  @click="verifyEmailCode"
                  :loading="verifying"
                  class="verify-btn"
                >
                  确定更换
                </el-button>
              </div>
              <div v-if="codeError" class="error-text">{{ codeError }}</div>
            </div>
          </div>

          <!-- 步骤 3：更换成功 -->
          <div v-show="currentStep === 3" class="step-content step-3 success-step">
            <div class="success-content">
              <div class="success-icon">
                <el-icon color="#67C23A" :size="60"><CircleCheck /></el-icon>
              </div>
              <h3>邮箱更换成功！</h3>
              <p class="success-text">您的绑定邮箱已成功更换为：</p>
              <p class="new-email">{{ form.email }}</p>
              <p class="tips">请使用新邮箱进行登录和验证操作</p>
              <div class="actions">
                <el-button 
                  type="primary" 
                  @click="completeChange"
                  :loading="completing"
                >
                  {{ completing ? '刷新中...' : '完成' }}
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>

    <!-- 安全提示 -->
    <div class="security-tips">
      <el-alert
        title="安全提示"
        type="warning"
        :closable="false"
        show-icon
      >
        <ul>
          <li>请确保您能正常登录新邮箱，以便接收验证邮件</li>
          <li>更换邮箱后，原邮箱将无法用于登录和验证</li>
          <li>验证码将在 5 分钟后失效，请及时使用</li>
        </ul>
      </el-alert>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onUnmounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Message,
  Warning,
  Refresh,
  Lock,
  CircleCheck,
  CircleClose
} from '@element-plus/icons-vue'
import { useUserStore } from '@/store/modules/user'
import { sendEmailApi, verifyCode, changeUserEmail } from '@/api/system/user/user'

const userStore = useUserStore()
const { userInfo } = userStore

// 状态管理
const isChanging = ref(false)
const currentStep = ref(1)
const sendingCode = ref(false)
const verifying = ref(false)
const completing = ref(false) // 新增：完成按钮加载状态
const countdown = ref(0)
let countdownTimer: NodeJS.Timeout | null = null

// 表单数据
const form = reactive({
  email: '',
  code: ''
})

// 错误信息
const emailError = ref('')
const codeError = ref('')

// 计算属性
const countdownText = computed(() => {
  if (countdown.value > 0) {
    return `${countdown.value}秒后重发`
  }
  return '重新发送'
})

// 邮箱验证
const validateEmail = () => {
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
  
  if (!form.email) {
    emailError.value = '请输入邮箱地址'
    return false
  }
  
  if (form.email === userInfo.email) {
    emailError.value = '新邮箱不能与当前绑定邮箱相同'
    return false
  }
  
  if (!emailRegex.test(form.email)) {
    emailError.value = '请输入有效的邮箱地址'
    return false
  }
  
  emailError.value = ''
  return true
}

// 清空邮箱
const clearEmail = () => {
  form.email = ''
  emailError.value = ''
}

// 开始更换流程
const startChangeEmail = () => {
  isChanging.value = true
  currentStep.value = 1
  resetForm()
}

// 取消更换
const cancelChange = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要取消更换绑定邮箱吗？已填写的邮箱地址将会丢失。',
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        closeOnClickModal: false
      }
    )
    isChanging.value = false
    resetForm()
  } catch {
    // 用户点击了取消
  }
}

// 发送验证码
const sendVerificationCode = async () => {
  if (!validateEmail()) return
  
  sendingCode.value = true
  try {
    const response = await sendEmailApi(form.email)
    
    if (response.data.code === 200) {
      ElMessage.success(`验证码已发送至 ${form.email}，请查收邮件`)
      startCountdown()
      currentStep.value = 2
    } else if (response.data.code === 60403) {
      ElMessage.error('登录状态已失效，请重新登录')
      // 跳转到登录页
      window.location.href = '/login'
    } else {
      ElMessage.error(response.data.msg || '验证码发送失败')
    }
  } catch (error: any) {
    if (error.response?.data?.msg) {
      ElMessage.error(error.response.data.msg)
    } else if (error.code === 'ERR_NETWORK') {
      ElMessage.error('网络连接失败，请检查网络')
    } else {
      ElMessage.error('发送失败，请稍后重试')
    }
  } finally {
    sendingCode.value = false
  }
}

// 重新发送验证码
const resendCode = async () => {
  if (countdown.value > 0) return
  await sendVerificationCode()
}

// 启动倒计时
const startCountdown = () => {
  countdown.value = 60
  if (countdownTimer) {
    clearInterval(countdownTimer)
  }
  countdownTimer = setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--
    } else {
      clearInterval(countdownTimer!)
      countdownTimer = null
    }
  }, 1000)
}

// 修改验证验证码函数
const verifyEmailCode = async () => {
  if (!form.code || form.code.length !== 6) {
    codeError.value = '请输入 6 位验证码'
    return
  }
  
  verifying.value = true
  try {
    // 步骤1：验证验证码
    const verifyResponse = await verifyCode(form.email, form.code)
    
    if (verifyResponse.data.code === 200) {
      // 验证码验证成功
      ElMessage.success('验证码验证成功')
      
      // 步骤2：调用更换邮箱接口
      const changeResponse = await changeUserEmail(form.email, form.code)
      
      if (changeResponse.data.code === 200) {
        // 邮箱更换成功
        currentStep.value = 3
        ElMessage.success('邮箱更换成功')
        userStore.setUserPartInfo({ email: form.email })
        
        // 更新本地用户信息
        await refreshUserInfo()
        
      } else {
        // 更换邮箱失败
        codeError.value = changeResponse.data.msg || '邮箱更换失败'
        ElMessage.error(changeResponse.data.msg || '邮箱更换失败，请重试')
      }
      
    } else {
      // 验证码验证失败
      codeError.value = verifyResponse.data.msg || '验证码错误'
      ElMessage.error(verifyResponse.data.msg || '验证码验证失败')
    }
    
  } catch (error: any) {
    if (error.response?.data?.msg) {
      codeError.value = error.response.data.msg
      ElMessage.error(error.response.data.msg)
    } else if (error.code === 'ERR_NETWORK') {
      codeError.value = '网络连接失败，请检查网络'
      ElMessage.error('网络连接失败，请检查网络')
    } else {
      codeError.value = '验证失败，请稍后重试'
      ElMessage.error('验证失败，请稍后重试')
    }
  } finally {
    verifying.value = false
  }
}

// 刷新用户信息
const refreshUserInfo = async () => {
  try {
    // 调用用户store中的方法来刷新用户信息
    // 假设您的userStore有一个getUserInfo方法
    await userStore.refreshUserInfo()
    
    // 或者如果您有单独的API接口来获取用户信息
    // const userInfoResponse = await getUserInfoApi()
    // userStore.setUserInfo(userInfoResponse.data)
    
  } catch (error) {
    console.error('刷新用户信息失败:', error)
    // 这里不阻止流程，即使刷新失败也继续
  }
}

// 完成更换
const completeChange = async () => {
  completing.value = true
  try {
    // 再次刷新用户信息，确保显示最新数据
    await refreshUserInfo()
    
    isChanging.value = false
    ElMessage.success('邮箱更换已完成')
    resetForm()
  } catch (error) {
    console.error('完成流程出错:', error)
  } finally {
    completing.value = false
  }
}

// 重置表单
const resetForm = () => {
  form.email = ''
  form.code = ''
  emailError.value = ''
  codeError.value = ''
  currentStep.value = 1
  sendingCode.value = false
  verifying.value = false
  if (countdownTimer) {
    clearInterval(countdownTimer)
    countdownTimer = null
  }
  countdown.value = 0
}

// 清理定时器
onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer)
  }
})

</script>



<style scoped lang="scss">
.email-change-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 24px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.current-email-section {
  margin-bottom: 32px;
  padding-bottom: 24px;
  border-bottom: 1px solid #f0f0f0;
  
  .info-header {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 16px;
    
    .title {
      font-size: 18px;
      font-weight: 600;
      color: #333;
    }
    
    .el-icon {
      color: #409eff;
    }
  }
  
  .info-content {
    .current-email-display {
      display: flex;
      align-items: center;
      gap: 8px;
      margin-bottom: 12px;
      
      .label {
        color: #666;
        font-size: 14px;
      }
      
      .value {
        color: #333;
        font-size: 16px;
        font-weight: 500;
      }
    }
    
    .hint-text {
      display: flex;
      align-items: center;
      gap: 8px;
      color: #e6a23c;
      font-size: 13px;
      margin: 0;
      
      .el-icon {
        font-size: 16px;
      }
    }
  }
}

.change-form-section {
  margin-bottom: 32px;
  
  .change-btn {
    padding: 10px 20px;
    font-size: 14px;
    
    .el-icon {
      margin-right: 6px;
    }
  }
  
  .form-container {
    .steps {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-bottom: 40px;
      
      .step-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        position: relative;
        
        .step-number {
          width: 32px;
          height: 32px;
          border-radius: 50%;
          background: #f5f7fa;
          border: 2px solid #dcdfe6;
          display: flex;
          align-items: center;
          justify-content: center;
          color: #c0c4cc;
          font-size: 14px;
          font-weight: 600;
          margin-bottom: 8px;
          transition: all 0.3s;
        }
        
        .step-label {
          color: #909399;
          font-size: 12px;
          white-space: nowrap;
        }
        
        &.active {
          .step-number {
            background: #409eff;
            border-color: #409eff;
            color: white;
          }
          
          .step-label {
            color: #409eff;
            font-weight: 500;
          }
        }
        
        &.completed {
          .step-number {
            background: #67c23a;
            border-color: #67c23a;
            color: white
          }
        }
      }
      
      .step-line {
        width: 60px;
        height: 2px;
        background: #f0f0f0;
        margin: 0 16px;
        position: relative;
        top: -15px;
      }
    }
    
    .step-content {
      .form-group {
        margin-bottom: 24px;
        
        .form-label {
          display: block;
          margin-bottom: 8px;
          color: #333;
          font-weight: 500;
          font-size: 14px;
        }
        
        .hint {
          color: #999;
          font-size: 12px;
          margin-top: 6px;
        }
        
        // 包裹输入框和操作按钮的容器
        .input-with-action {
          display: flex;
          gap: 12px;
          align-items: center;
          width: 100%;
          box-sizing: border-box;
          
          // 两个输入框统一宽度
          .el-input {
            flex: 1;
            min-width: 0;
          }
          
          // 步骤 1 取消按钮
          .cancel-btn {
            flex: 0 0 70px;
            height: 40px;
            white-space: nowrap;
          }
          
          // 步骤 2 确定更换按钮
          .verify-btn {
            flex: 0 0 100px;
            height: 40px;
            font-size: 14px;
            font-weight: 500;
            white-space: nowrap;
          }
        }
        
        // 邮箱输入框带按钮样式（按钮在 suffix，灰色）
        .email-input-with-btn {
          :deep(.el-input__suffix) {
            right: 4px;
            display: flex;
            align-items: center;
            
            .suffix-group {
              display: flex;
              align-items: center;
              gap: 4px;
            }
            
            .clear-btn {
              width: 16px;
              height: 16px;
              color: #c0c4cc;
              cursor: pointer;
              display: flex;
              align-items: center;
              justify-content: center;
              border-radius: 50%;
              flex-shrink: 0;
              
              &:hover {
                color: #f56c6c;
                background-color: #f5f7fa;
              }
            }
            
            .send-code-btn {
              height: 30px;
              padding: 0 10px;
              font-size: 12px;
              border-radius: 4px;
              color: #909399;
              border-color: #dcdfe6;
              background-color: #f5f7fa;
              flex-shrink: 0;
              
              &:disabled {
                opacity: 0.6;
                cursor: not-allowed;
              }
              
              &:not(:disabled):hover {
                color: #409eff;
                border-color: #409eff;
                background-color: #ecf5ff;
              }
            }
          }
          
          :deep(.el-input__inner) {
            padding-right: 95px;
          }
        }
        
        // 验证码输入框带按钮样式（按钮在 suffix，灰色样式）
        .code-input-with-btn {
          :deep(.el-input__prefix) {
            left: 8px;
            
            .el-icon {
              color: #909399;
            }
          }
          
          :deep(.el-input__suffix) {
            right: 4px;
            display: flex;
            align-items: center;
            
            .resend-code-btn {
              height: 30px;
              padding: 0 8px;
              font-size: 12px;
              border-radius: 4px;
              color: #909399;
              border-color: #dcdfe6;
              background-color: #f5f7fa;
              flex-shrink: 0;
              
              &:disabled {
                opacity: 0.6;
                cursor: not-allowed;
              }
              
              &:not(:disabled):hover {
                color: #409eff;
                border-color: #409eff;
                background-color: #ecf5ff;
              }
            }
          }
          
          :deep(.el-input__inner) {
            padding-right: 75px;
          }
        }
        
        .code-input-group {
          display: flex;
          gap: 12px;
          
          .code-input {
            flex: 1;
          }
          
          .resend-btn {
            white-space: nowrap;
            min-width: 120px;
          }
        }
      }
      
      .actions {
        display: flex;
        justify-content: flex-end;
        gap: 12px;
        margin-top: 32px;
      }
      
      .error-text {
        color: #f56c6c;
        font-size: 12px;
        margin-top: 6px;
      }
    }
    
    .success-step {
      text-align: center;
      padding: 40px 0;
      
      .success-icon {
        margin-bottom: 20px;
        
        .el-icon {
          animation: scaleIn 0.5s ease-out;
        }
      }
      
      h3 {
        color: #333;
        margin-bottom: 12px;
        font-weight: 600;
      }
      
      .success-text {
        color: #666;
        margin-bottom: 8px;
      }
      
      .new-email {
        color: #409eff;
        font-size: 18px;
        font-weight: 600;
        margin-bottom: 16px;
      }
      
      .tips {
        color: #999;
        font-size: 13px;
        margin-bottom: 24px;
      }
    }
  }
}

.security-tips {
  margin-top: 24px;
  
  .el-alert {
    background-color: #fdf6ec;
    border: 1px solid #f5dab1;
    
    :deep(.el-alert__title) {
      color: #e6a23c;
      font-weight: 600;
    }
    
    ul {
      margin: 8px 0 0 0;
      padding-left: 20px;
      color: #e6a23c;
      
      li {
        margin-bottom: 4px;
        font-size: 13px;
      }
    }
  }
}

@keyframes scaleIn {
  0% {
    transform: scale(0);
    opacity: 0;
  }
  50% {
    transform: scale(1.2);
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}
</style>