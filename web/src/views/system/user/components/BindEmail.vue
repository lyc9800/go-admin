<template>
  <div class="bind-email-container">
    <!-- 当前邮箱信息 -->
    <div class="current-email-section">
      <div class="info-header">
        <el-icon><Message /></el-icon>
        <span class="title">绑定邮箱</span>
      </div>
      <div class="info-content">
        <div class="current-email-display">
          <span class="label">当前绑定邮箱：</span>
          <span class="value">{{ userInfo.email || "未绑定" }}</span>
        </div>
        <p class="hint-text">
          <el-icon><Warning /></el-icon>
          邮箱是登录验证的重要方式
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
          {{ userInfo.email ? "更换绑定邮箱" : "绑定邮箱" }}
        </el-button>
      </template>

      <!-- 更换流程表单 -->
      <template v-else>
        <div class="form-container">
          <!-- 步骤指示器 -->
          <div class="steps">
            <div
              class="step-item"
              :class="{ active: currentStep === 1, completed: currentStep > 1 }"
            >
              <div class="step-number">1</div>
              <div class="step-label">输入邮箱</div>
            </div>
            <div class="step-line"></div>
            <div
              class="step-item"
              :class="{ active: currentStep === 2, completed: currentStep > 2 }"
            >
              <div class="step-number">2</div>
              <div class="step-label">验证</div>
            </div>
            <div class="step-line"></div>
            <div
              class="step-item"
              :class="{ active: currentStep === 3 }"
            >
              <div class="step-number">3</div>
              <div class="step-label">完成</div>
            </div>
          </div>

          <!-- 步骤 1：输入新邮箱 -->
          <div v-show="currentStep === 1" class="step-content step-1">
            <div class="form-group">
              <label class="form-label">邮箱地址</label>
              <div class="input-with-action">
                <el-input
                  v-model="form.email"
                  :placeholder="userInfo.email ? '请输入新的邮箱地址' : '请输入要绑定的邮箱地址'"
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
                        {{ sendingCode ? "发送中..." : "发送验证码" }}
                      </el-button>
                    </div>
                  </template>
                </el-input>
                <el-button @click="cancelChange" class="cancel-btn">
                  取消
                </el-button>
              </div>
              <div v-if="emailError" class="error-text">{{ emailError }}</div>
              <div v-else class="hint">
                {{ userInfo.email ? "请输入要绑定的新邮箱地址" : "请输入要绑定的邮箱地址" }}
              </div>
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
                  @keyup.enter="verifyEmailCode"
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
                <el-button
                  type="primary"
                  :disabled="!form.code || form.code.length !== 6"
                  @click="verifyEmailCode"
                  :loading="verifying"
                  class="verify-btn"
                >
                  {{ userInfo.email ? "确定更换" : "确定绑定" }}
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
              <h3>{{ userInfo.email ? "邮箱更换成功！" : "邮箱绑定成功！" }}</h3>
              <p class="success-text">
                {{ userInfo.email ? "您的绑定邮箱已成功更换为：" : "您的邮箱已成功绑定：" }}
              </p>
              <p class="new-email">{{ form.email }}</p>
              <p class="tips">请使用此邮箱进行登录和验证操作</p>
              <div class="actions">
                <el-button
                  type="primary"
                  @click="completeChange"
                  :loading="completing"
                >
                  {{ completing ? "刷新中..." : "完成" }}
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onUnmounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import {
  Message,
  Warning,
  Refresh,
  Lock,
  CircleCheck,
  CircleClose,
} from "@element-plus/icons-vue";
import { useUserStore } from "@/store/modules/user";
import { sendEmailApi, verifyCode, changeUserEmail } from "@/api/system/user/user";

const userStore = useUserStore();
const { userInfo } = userStore;

// 状态管理
const isChanging = ref(false);
const currentStep = ref(1);
const sendingCode = ref(false);
const verifying = ref(false);
const completing = ref(false);
const countdown = ref(0);
let countdownTimer: NodeJS.Timeout | null = null;

// 表单数据
const form = reactive({
  email: "",
  code: "",
});

// 错误信息
const emailError = ref("");
const codeError = ref("");

// 计算属性
const countdownText = computed(() => {
  if (countdown.value > 0) {
    return `${countdown.value}秒后重发`;
  }
  return "重新发送";
});

// 邮箱验证
const validateEmail = () => {
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

  if (!form.email) {
    emailError.value = "请输入邮箱地址";
    return false;
  }

  if (form.email === userInfo.email) {
    emailError.value = "新邮箱不能与当前绑定邮箱相同";
    return false;
  }

  if (!emailRegex.test(form.email)) {
    emailError.value = "请输入有效的邮箱地址";
    return false;
  }

  emailError.value = "";
  return true;
};

// 清空邮箱
const clearEmail = () => {
  form.email = "";
  emailError.value = "";
};

// 开始更换流程
const startChangeEmail = () => {
  isChanging.value = true;
  currentStep.value = 1;
  resetForm();
};

// 取消更换
const cancelChange = async () => {
  try {
    await ElMessageBox.confirm(
      "确定要取消邮箱操作吗？已填写的邮箱地址将会丢失。",
      "提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        closeOnClickModal: false,
      }
    );
    isChanging.value = false;
    resetForm();
  } catch {
    // 用户点击了取消
  }
};

// 发送验证码
const sendVerificationCode = async () => {
  if (!validateEmail()) return;

  sendingCode.value = true;
  try {
    const response = await sendEmailApi(form.email);

    if (response.data.code === 200) {
      ElMessage.success(`验证码已发送至 ${form.email}，请查收邮件`);
      startCountdown();
      currentStep.value = 2;
    } else if (response.data.code === 60403) {
      ElMessage.error("登录状态已失效，请重新登录");
      window.location.href = "/login";
    } else {
      ElMessage.error(response.data.msg || "验证码发送失败");
    }
  } catch (error: any) {
    if (error.response?.data?.msg) {
      ElMessage.error(error.response.data.msg);
    } else if (error.code === "ERR_NETWORK") {
      ElMessage.error("网络连接失败，请检查网络");
    } else {
      ElMessage.error("发送失败，请稍后重试");
    }
  } finally {
    sendingCode.value = false;
  }
};

// 重新发送验证码
const resendCode = async () => {
  if (countdown.value > 0) return;
  await sendVerificationCode();
};

// 启动倒计时
const startCountdown = () => {
  countdown.value = 60;
  if (countdownTimer) {
    clearInterval(countdownTimer);
  }
  countdownTimer = setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--;
    } else {
      clearInterval(countdownTimer!);
      countdownTimer = null;
    }
  }, 1000);
};

// 验证验证码
const verifyEmailCode = async () => {
  if (!form.code || form.code.length !== 6) {
    codeError.value = "请输入 6 位验证码";
    return;
  }

  verifying.value = true;
  try {
    // 步骤1：验证验证码
    const verifyResponse = await verifyCode(form.email, form.code);

    if (verifyResponse.data.code === 200) {
      // 验证码验证成功
      ElMessage.success("验证码验证成功");

      // 步骤2：调用更换邮箱接口
      const changeResponse = await changeUserEmail(form.email, form.code);

      if (changeResponse.data.code === 200) {
        // 邮箱更换成功
        currentStep.value = 3;
        ElMessage.success(
          userInfo.email ? "邮箱更换成功" : "邮箱绑定成功"
        );
        userStore.setUserPartInfo({ email: form.email });
      } else {
        // 更换邮箱失败
        codeError.value = changeResponse.data.msg || "操作失败";
        ElMessage.error(changeResponse.data.msg || "操作失败，请重试");
      }
    } else {
      // 验证码验证失败
      codeError.value = verifyResponse.data.msg || "验证码错误";
      ElMessage.error(verifyResponse.data.msg || "验证码验证失败");
    }
  } catch (error: any) {
    if (error.response?.data?.msg) {
      codeError.value = error.response.data.msg;
      ElMessage.error(error.response.data.msg);
    } else if (error.code === "ERR_NETWORK") {
      codeError.value = "网络连接失败，请检查网络";
      ElMessage.error("网络连接失败，请检查网络");
    } else {
      codeError.value = "验证失败，请稍后重试";
      ElMessage.error("验证失败，请稍后重试");
    }
  } finally {
    verifying.value = false;
  }
};

// 完成更换
const completeChange = async () => {
  completing.value = true;
  try {
    isChanging.value = false;
    ElMessage.success("操作已完成");
    resetForm();
  } catch (error) {
    console.error("完成流程出错:", error);
  } finally {
    completing.value = false;
  }
};

// 重置表单
const resetForm = () => {
  form.email = "";
  form.code = "";
  emailError.value = "";
  codeError.value = "";
  currentStep.value = 1;
  sendingCode.value = false;
  verifying.value = false;
  if (countdownTimer) {
    clearInterval(countdownTimer);
    countdownTimer = null;
  }
  countdown.value = 0;
};

// 清理定时器
onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer);
  }
})
</script>

<style scoped lang="scss">
.bind-email-container {
  width: 100%;
  height: 100%;
  padding: 20px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
}

.current-email-section {
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f0f0f0;

  .info-header {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 12px;

    .title {
      font-size: 16px;
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
      margin-bottom: 8px;

      .label {
        color: #666;
        font-size: 13px;
      }

      .value {
        color: #333;
        font-size: 14px;
        font-weight: 500;
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
}

.change-form-section {
  flex: 1;
  display: flex;
  flex-direction: column;

  .change-btn {
    padding: 8px 16px;
    font-size: 13px;
    width: 100%;

    .el-icon {
      margin-right: 4px;
    }
  }

  .form-container {
    flex: 1;
    display: flex;
    flex-direction: column;

    .steps {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-bottom: 20px;

      .step-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        position: relative;
        min-width: 60px;

        .step-number {
          width: 24px;
          height: 24px;
          border-radius: 50%;
          background: #f5f7fa;
          border: 2px solid #dcdfe6;
          display: flex;
          align-items: center;
          justify-content: center;
          color: #c0c4cc;
          font-size: 12px;
          font-weight: 600;
          margin-bottom: 6px;
          transition: all 0.3s;
        }

        .step-label {
          color: #909399;
          font-size: 11px;
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
            color: white;
          }
        }
      }

      .step-line {
        flex: 1;
        height: 2px;
        background: #f0f0f0;
        margin: 0 8px;
        position: relative;
        top: -10px;
        min-width: 20px;
      }
    }

    .step-content {
      flex: 1;
      display: flex;
      flex-direction: column;
      justify-content: center;

      .form-group {
        margin-bottom: 15px;

        .form-label {
          display: block;
          margin-bottom: 6px;
          color: #333;
          font-weight: 500;
          font-size: 13px;
        }

        .hint {
          color: #999;
          font-size: 11px;
          margin: 4px 0 8px 0;
        }

        .input-with-action {
          display: flex;
          gap: 8px;
          align-items: center;
          width: 100%;
          box-sizing: border-box;

          .el-input {
            flex: 1;
            min-width: 0;
          }

          .cancel-btn {
            flex: 0 0 50px;
            height: 36px;
            font-size: 12px;
            white-space: nowrap;
          }

          .verify-btn {
            flex: 0 0 80px;
            height: 36px;
            font-size: 12px;
            font-weight: 500;
            white-space: nowrap;
          }
        }

        .error-text {
          color: #f56c6c;
          font-size: 11px;
          margin-top: 4px;
        }
      }
    }

    .success-step {
      text-align: center;
      padding: 20px 0;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      .success-icon {
        margin-bottom: 12px;
      }

      h3 {
        color: #333;
        margin-bottom: 8px;
        font-weight: 600;
        font-size: 16px;
      }

      .success-text {
        color: #666;
        margin-bottom: 6px;
        font-size: 13px;
      }

      .new-email {
        color: #409eff;
        font-size: 16px;
        font-weight: 600;
        margin-bottom: 12px;
      }

      .tips {
        color: #999;
        font-size: 12px;
        margin-bottom: 16px;
      }

      .actions {
        display: flex;
        justify-content: center;
      }
    }
  }
}
</style>