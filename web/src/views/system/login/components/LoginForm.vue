<template>
  <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules">
    <el-form-item label="" prop="username">
      <el-input
        placeholder="请输入用户名"
        autocomplete="on"
        style="position: relative"
        v-model="ruleForm.username"
        class="fixed-height-input"
      >
        <template #prefix>
          <el-icon><UserFilled/></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item label="" prop="password">
      <el-input
        placeholder="请输入密码"
        autocomplete="on"
        v-model="ruleForm.password"
        class="fixed-height-input"
        type="password"
        show-password 
      >
        <template #prefix>
          <el-icon><GoodsFilled/></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item style="width: 100%;">
      <el-button
        :loading="loading"
        class="login-btn"
        color="#e99d53"
        @click="submitForm(ruleFormRef)"
      >
        登录
      </el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import  type { FormInstance, FormRules } from 'element-plus'
import { useRouter } from 'vue-router'
import { UserFilled, GoodsFilled } from '@element-plus/icons-vue'
import {ElNotification} from 'element-plus'
import { loginApi } from '@/api/system/login/login'
import { useUserStore } from '@/store/modules/user'

const userStore = useUserStore()
// 路由对象
const router = useRouter()

// 表单实例对象
const ruleFormRef = ref<FormInstance>()

// 密码类型

// 提交表单后状态
const loading = ref(false)

// 表单规则
const rules = reactive<FormRules>({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
})

// 表单数据对象
const ruleForm = reactive({
  username: 'admin',
  password: '123456'
})
// 提交表单函数
const submitForm=(formEl: FormInstance | undefined)=>{
  loading.value = true
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      // 登录成功
      const {data} = await loginApi({...ruleForm})
      if(data.code === 200){
        // 设置token
        userStore.setToken(data.token)
        // 设置用户信息
        userStore.setUserInfo(data.userInfo)
        // 跳转到首页
        setTimeout(()=>{
          // window.location.href='/home'
          router.push({path:'/home'})
        },1500)
        ElNotification({
          title:'登陆成功',
          message:'欢迎登陆后台管理系统',
          type:'success',
          duration: 3000
        })
      }else{
        ElNotification({
          title:'温馨提示',
          message:data.msg,
          type:'error',
          duration: 3000
        })
      }
    } else {
      // 登录失败
      ElNotification({
        title:'温馨提示',
        message:'还有必填项没有填写',
        type:'error',
        duration: 3000
      })
      return false  
    }
  })
  loading.value = false
}
</script>

<style scoped>
.login-btn {
  margin-top: 1px;
  width: 100%;
  height: 32px;
  color: white;
  font-size: 16px;

}
/* ========== 修复密码输入框尺寸闪烁问题 ========== */

/* 1. 统一所有输入框高度 */
:deep(.el-input) {
  --input-height: 40px; /* 定义统一高度变量 */
  height: var(--input-height);
  min-height: var(input-height);
  display: block;
  position: relative;
}

/* 2. 固定输入框包装器 - 核心修复 */
:deep(.el-input__wrapper) {
  height: var(--input-height) !important;
  min-height: var(--input-height) !important;
  display: flex;
  align-items: center;
  padding: 0 12px;
  box-sizing: border-box;
  /* 关键：为眼睛图标预留固定空间 */
  padding-right: 40px !important;
  /* 禁用所有可能导致重排的过渡 */
  transition: none !important;
  animation: none !important;
}

/* 3. 内部输入框 */
:deep(.el-input__inner) {
  height: 100%;
  min-height: 100%;
  line-height: normal;
  padding: 0;
  flex: 1;
  /* 防止文本变化影响布局 */
  min-width: 0;
}

/* 4. 眼睛图标容器 - 绝对定位，永不重排 */
:deep(.el-input__suffix) {
  position: absolute !important;
  right: 8px;
  top: 0;
  height: var(--input-height) !important;
  width: 24px !important;
  min-width: 24px !important;
  display: flex;
  align-items: center;
  justify-content: center;
  /* 关键：始终可见，不随状态变化 */
  opacity: 1 !important;
  visibility: visible !important;
  /* 禁用所有过渡 */
  transition: none !important;
  transform: none !important;
  /* 确保在清空按钮上层 */
  z-index: 3;
}

/* 5. 眼睛图标本身 */
:deep(.el-input__suffix .el-icon) {
  width: 20px;
  height: 20px;
  font-size: 20px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: none !important;
  cursor: pointer;
}

/* 6. 清空按钮容器 - 在眼睛图标后面 */
:deep(.el-input__suffix .el-input__clear) {
  position: absolute;
  right: 32px; /* 在眼睛图标左边 */
  z-index: 2;
}

/* 7. 前缀图标 */
:deep(.el-input__prefix) {
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1;
}

/* 8. 用户名输入框（没有眼睛图标）调整 */
:deep(.el-form-item:first-child .el-input__wrapper) {
  padding-right: 12px !important; /* 正常内边距 */
}

/* 9. 登录按钮 */
.login-btn {
  margin-top: 1px;
  width: 100%;
  height: 40px; /* 与输入框保持一致 */
  color: white;
  font-size: 16px;
  font-weight: 500;
}

</style>