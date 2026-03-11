<template>
  <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules">
    <el-form-item label="" prop="username">
      <el-input
        placeholder="请输入用户名"
        autocomplete="on"
        style="position: relative"
        v-model="ruleForm.username"
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
        :type="passwordType"
      >
        <template #prefix>
          <el-icon><GoodsFilled /></el-icon>
        </template>
        <template #suffix>
          <div class="show-pwd" @click="showPwd">
            <svg-icon
              :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"/>
          </div>
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
import SvgIcon from '@/components/SvgIcon/Index.vue'
import { ref, reactive } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { ElNotification } from 'element-plus'
import { useRouter } from 'vue-router'
import { UserFilled, GoodsFilled } from '@element-plus/icons-vue'

// 路由对象
const router = useRouter()

// 表单实例对象
const ruleFormRef = ref<FormInstance>()

// 密码类型
const passwordType = ref('password')

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

// 显示/隐藏密码
const showPwd = () => {
  passwordType.value = passwordType.value === 'password' ? '' : 'password'
}

</script>

<style scoped>

</style>