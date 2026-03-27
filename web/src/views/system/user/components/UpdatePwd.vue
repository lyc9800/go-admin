<template>
  <div class="update-password-container">
    <!-- 修改密码标题 -->
    <div class="password-header">
      <div class="info-header">
        <el-icon><Lock /></el-icon>
        <span class="title">修改密码</span>
      </div>
      <p class="hint-text">
        <el-icon><Warning /></el-icon>
        请定期修改密码以确保账户安全
      </p>
    </div>

    <!-- 修改密码表单 -->
    <div class="password-form">
      <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="80px"
        class="password-form-content"
      >
        <!-- 当前密码 -->
        <el-form-item label="当前密码" prop="oldPassword">
          <el-input
            v-model="passwordForm.oldPassword"
            placeholder="请输入当前密码"
            type="password"
            show-password
            size="medium"
            @keyup.enter="handleChangePassword"
          >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>

        <!-- 新密码 -->
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="passwordForm.newPassword"
            placeholder="请输入新密码（6-20位字符）"
            type="password"
            show-password
            size="medium"
            @keyup.enter="handleChangePassword"
          >
            <template #prefix>
              <el-icon><Key /></el-icon>
            </template>
          </el-input>
          <div class="password-strength" v-if="passwordForm.newPassword">
            强度：<span :class="strengthClass">{{ strengthText }}</span>
          </div>
        </el-form-item>

        <!-- 确认密码 -->
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="passwordForm.confirmPassword"
            placeholder="请再次输入新密码"
            type="password"
            show-password
            size="medium"
            @keyup.enter="handleChangePassword"
          >
            <template #prefix>
              <el-icon><Check /></el-icon>
            </template>
          </el-input>
        </el-form-item>

        <!-- 表单操作 -->
        <el-form-item>
          <div class="form-actions">
            <el-button
              type="primary"
              :loading="changing"
              @click="handleChangePassword"
              class="submit-btn"
            >
              {{ changing ? "修改中..." : "确认修改" }}
            </el-button>
            <el-button @click="handleResetForm" class="reset-btn">
              重置
            </el-button>
          </div>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from "vue";
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from "element-plus";
import { Lock, Warning, Key, Check } from "@element-plus/icons-vue";
import { useUserStore } from "@/store/modules/user";
import { updatePasswordApi } from "@/api/system/user/user";

const userStore = useUserStore();
const passwordFormRef = ref<FormInstance>();

// 状态管理
const changing = ref(false);

// 表单数据
const passwordForm = reactive({
  oldPassword: "",
  newPassword: "",
  confirmPassword: "",
});

// 密码强度计算
const passwordStrength = computed(() => {
  const password = passwordForm.newPassword;
  if (!password) return 0;

  let strength = 0;

  // 长度检查
  if (password.length >= 8) strength += 1;
  if (password.length >= 12) strength += 1;

  // 包含小写字母
  if (/[a-z]/.test(password)) strength += 1;

  // 包含大写字母
  if (/[A-Z]/.test(password)) strength += 1;

  // 包含数字
  if (/\d/.test(password)) strength += 1;

  // 包含特殊字符
  if (/[^a-zA-Z0-9]/.test(password)) strength += 1;

  return Math.min(strength, 5); // 最大强度为5
});

// 密码强度文本
const strengthText = computed(() => {
  const strength = passwordStrength.value;
  if (strength <= 1) return "弱";
  if (strength <= 3) return "中";
  return "强";
});

// 密码强度样式
const strengthClass = computed(() => {
  const strength = passwordStrength.value;
  if (strength <= 1) return "strength-weak";
  if (strength <= 3) return "strength-medium";
  return "strength-strong";
});

// 表单验证规则
const passwordRules = reactive<FormRules>({
  oldPassword: [
    { required: true, message: "请输入当前密码", trigger: "blur" },
    { min: 6, max: 20, message: "密码长度为6-20个字符", trigger: "blur" },
  ],
  newPassword: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 6, max: 20, message: "密码长度为6-20个字符", trigger: "blur" },
    {
      validator: (_, value, callback) => {
        if (value === passwordForm.oldPassword) {
          callback(new Error("新密码不能与当前密码相同"));
        } else {
          callback();
        }
      },
      trigger: "blur",
    },
  ],
  confirmPassword: [
    { required: true, message: "请确认新密码", trigger: "blur" },
    {
      validator: (_, value, callback) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error("两次输入的密码不一致"));
        } else {
          callback();
        }
      },
      trigger: "blur",
    },
  ],
});

// 处理修改密码
const handleChangePassword = async () => {
  if (!passwordFormRef.value) return;

  try {
    // 验证表单
    const valid = await passwordFormRef.value.validate();
    if (!valid) return;

    // 确认修改
    await ElMessageBox.confirm(
      "确认要修改密码吗？修改后需要重新登录。",
      "确认修改",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        closeOnClickModal: false,
      }
    );

    changing.value = true;

    // 调用修改密码接口
    const response = await updatePasswordApi({
      oldPassword: passwordForm.oldPassword,
      newPassword: passwordForm.newPassword,
    });

    if (response.data.code === 200) {
      ElMessage.success("密码修改成功，请重新登录");

      // 延迟跳转到登录页
      setTimeout(() => {
        // 清除用户信息
        userStore.setToken("");
        userStore.setUserInfo({});

        // 跳转到登录页
        window.location.href = "/login";
      }, 1500);
    } else {
      ElMessage.error(response.data.msg || "密码修改失败");
    }
  } catch (error: any) {
    // 用户取消了修改
    if (error === "cancel") {
      return;
    }

    // API 错误
    if (error.response?.data?.msg) {
      ElMessage.error(error.response.data.msg);
    } else if (error.code === "ERR_NETWORK") {
      ElMessage.error("网络连接失败，请检查网络");
    } else {
      ElMessage.error("密码修改失败，请稍后重试");
    }
  } finally {
    changing.value = false;
  }
};

// 重置表单
const handleResetForm = () => {
  if (passwordFormRef.value) {
    passwordFormRef.value.resetFields();
  }
};
</script>

<style scoped lang="scss">
.update-password-container {
  width: 100%;
  height: 100%;
  padding: 20px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
}

.password-header {
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f0f0f0;

  .info-header {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 8px;

    .title {
      font-size: 16px;
      font-weight: 600;
      color: #333;
    }

    .el-icon {
      color: #e6a23c;
    }
  }

  .hint-text {
    display: flex;
    align-items: center;
    gap: 6px;
    color: #e6a23c;
    font-size: 12px;
    margin: 0;

    .el-icon {
      font-size: 14px;
    }
  }
}

.password-form {
  flex: 1;
  display: flex;
  flex-direction: column;

  .password-form-content {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;

    :deep(.el-form-item) {
      margin-bottom: 16px;
    }

    :deep(.el-form-item__label) {
      font-weight: 500;
      color: #333;
      font-size: 13px;
      padding-right: 8px;
    }

    .el-input {
      width: 100%;
    }

    .password-strength {
      font-size: 11px;
      margin-top: 4px;
      color: #999;

      .strength-weak {
        color: #f56c6c;
        font-weight: 500;
      }

      .strength-medium {
        color: #e6a23c;
        font-weight: 500;
      }

      .strength-strong {
        color: #67c23a;
        font-weight: 500;
      }
    }
  }

  .form-actions {
    display: flex;
    gap: 12px;
    width: 100%;
    margin-top: 20px;

    .submit-btn {
      flex: 1;
      padding: 8px 16px;
      font-size: 13px;
      font-weight: 500;
    }

    .reset-btn {
      flex: 1;
      padding: 8px 16px;
      font-size: 13px;
    }
  }
}

// 响应式调整
@media (max-width: 768px) {
  .update-password-container {
    padding: 15px;
  }

  .form-actions {
    flex-direction: column;
    gap: 8px;
  }
}
</style>