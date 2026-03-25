<template>
  <div class="set">
    <h4>绑定邮箱</h4>
    <p><span>已绑定邮箱：</span><span>{{ userEmail }}</span></p>
    <p>邮箱为登录验证的重要方式，请谨慎操作</p>

    <!-- 步骤1：输入邮箱 + 发送验证码 -->
    <div v-if="step === 1" class="step-1">
      <el-input
        v-model="toBind.email"
        placeholder="请输入新的邮箱地址"
        class="email-input"
        clearable
      />
      <el-button
        type="primary"
        :disabled="!canSend"
        @click="getCode"
        class="send-code-btn"
      >
        {{ codeText }}
      </el-button>
    </div>

    <!-- 步骤2：输入验证码 + 操作按钮 -->
    <div v-if="step === 2" class="step-2">
      <el-input
        v-model="toBind.code"
        placeholder="请输入邮箱验证码"
        class="code-input"
        clearable
      />
      <div class="action-buttons">
        <el-button type="primary" @click="toBindSubmit">确定</el-button>
        <el-button @click="resetForm">返回</el-button>
      </div>
    </div>

    <!-- 初始状态：更换绑定邮箱按钮 -->
    <el-button v-if="step === 0" type="primary" plain @click="startChange">
      更换绑定邮箱
    </el-button>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, computed, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/store/modules/user'
import { sendEmailApi } from '@/api/system/user/user'

const userStore = useUserStore()
const userEmail = computed(() => {
  const { email } = userStore.userInfo
  return email ? email.replace(email.substring(3, email.length - 3), '****') : ''
})

// 表单数据
const toBind = reactive({
  email: '',
  code: ''
})

// 步骤控制：0=初始，1=输入邮箱，2=输入验证码
const step = ref(0)

// 倒计时相关
const TIME_COUNT = 60
const count = ref(0)
const canSend = ref(true)
const codeText = ref('获取验证码')
let timer: number | null = null

// 重置定时器
const resetTimer = () => {
  if (timer) {
    clearInterval(timer)
    timer = null
  }
  canSend.value = true
  codeText.value = '获取验证码'
  count.value = 0
}

// 开始更换流程
const startChange = () => {
  // 重置所有状态
  toBind.email = ''
  toBind.code = ''
  resetTimer()
  step.value = 1
}

// 获取验证码
const getCode = async () => {
  if (!toBind.email) {
    ElMessage.warning('请输入邮箱地址')
    return
  }
  // 简单邮箱格式校验
  const emailRegex = /^[^\s@]+@([^\s@]+\.)+[^\s@]+$/
  if (!emailRegex.test(toBind.email)) {
    ElMessage.warning('请输入正确的邮箱地址')
    return
  }

  if (timer) return

  // 启动倒计时
  count.value = TIME_COUNT
  canSend.value = false
  timer = window.setInterval(() => {
    if (count.value > 0) {
      count.value--
      codeText.value = `${count.value}秒后重新获取`
    } else {
      resetTimer()
    }
  }, 1000)

  try {
    const { data } = await sendEmailApi(toBind.email)
    if (data.code === 200) {
      ElMessage.success('验证码已发送，请查收')
      step.value = 2 // 进入验证码输入步骤
    } else {
      ElMessage.error(data.msg || '验证码发送失败')
      resetTimer()
    }
  } catch (error) {
    ElMessage.error('发送失败，请稍后重试')
    resetTimer()
  }
}

// 返回上一步（重置验证码、停止倒计时）
const resetForm = () => {
  toBind.code = ''
  resetTimer()
  step.value = 1
}

// 提交绑定（需对接实际API）
const toBindSubmit = async () => {
  if (!toBind.code) {
    ElMessage.warning('请输入验证码')
    return
  }
  // TODO: 调用绑定邮箱接口，例如 bindEmailApi(toBind.email, toBind.code)
  // 假设绑定成功
  try {
    // const { data } = await bindEmailApi(toBind.email, toBind.code)
    // if (data.code === 200) {
    //   ElMessage.success('绑定成功')
    //   // 更新用户信息 store
    //   await userStore.getUserInfo()
    //   step.value = 0
    //   toBind.email = ''
    //   toBind.code = ''
    // } else {
    //   ElMessage.error(data.msg || '绑定失败')
    // }
    // 示例模拟成功
    ElMessage.success('绑定成功（演示）')
    step.value = 0
    toBind.email = ''
    toBind.code = ''
  } catch (error) {
    ElMessage.error('绑定失败，请稍后重试')
  }
}

// 组件卸载时清除定时器
onUnmounted(() => {
  if (timer) clearInterval(timer)
})
</script>

<style scoped>
.set {
  padding: 20px;
  max-width: 500px;
}
.set h4 {
  margin-bottom: 15px;
}
.set p {
  margin-bottom: 10px;
  color: #666;
}

/* 步骤1布局：邮箱输入框 + 发送按钮在一行 */
.step-1 {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-top: 16px;
}
.email-input {
  flex: 1;
}
.send-code-btn {
  white-space: nowrap;
}

/* 步骤2布局：验证码输入框 + 操作按钮 */
.step-2 {
  margin-top: 16px;
}
.code-input {
  width: 100%;
  margin-bottom: 16px;
}
.action-buttons {
  display: flex;
  gap: 12px;
}
</style>