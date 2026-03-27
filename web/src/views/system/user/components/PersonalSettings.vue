<template>
  <el-row :gutter="20">
    <!-- 左侧信息 start -->
    <div class="left-box">
      <h3 class="title">
        <el-icon><User /></el-icon>
        个人设置
      </h3>
      <!-- 基本信息 start -->
      <div class="set">
        <h4>基础设置</h4>
        <el-form
          ref="basicFormRef"
          :rules="basicRules"
          status-icon
          :model="basic"
        >
          <el-row :gutter="20">
            <!-- 用户昵称 start -->
            <el-col :span="8">
              <el-form-item label="用户账号" style="width: 100%">
                <el-input v-model="basic.userName" disabled></el-input>
              </el-form-item>
            </el-col>
            <!-- 用户昵称 end -->
            <!-- 性别 -->
            <el-col :span="8">
              <el-form-item label="性别" prop="sex" style="width: 100%">
                <el-radio v-model="basic.sex" label="男"></el-radio>
                <el-radio v-model="basic.sex" label="女"></el-radio>
              </el-form-item>
            </el-col>
            <!-- 性别 end -->
            <!--头像-->
            <el-col :span="5">
              <el-form-item label="头像" style="margin: auto" prop="avatar">
                <el-upload
                  class="avatar-uploader"
                  :action="uploadURL"
                  :headers="headerObj"
                  name="fileResource"
                  :show-file-list="false"
                  :on-success="handlerAvatarSuccess"
                >
                  <img
                    v-if="basic.avatar"
                    :src="url + 'uploadFile/' + basic.avatar"
                    style="width: 50px; border-radius: 50%"
                  />
                  <img
                    v-else
                    src="@/assets/default.png"
                    style="
                      width: 50px;
                      height: 50px;
                      object-fit: cover;
                      border-radius: 50%;
                    "
                  />
                  <span
                    style="
                      margin-left: 10px;
                      font-size: 10px;
                      text-decoration: underline;
                      line-height: 20px;
                    "
                    >点击更换</span
                  >
                </el-upload>
              </el-form-item>
            </el-col>
            <!-- 提交按钮 -->
            <el-col :span="3">
              <el-form-item>
                <el-button
                  plain
                  color="#2fa7b9"
                  style="margin-left: 50px"
                  :loading="loading"
                  @click="onBasicSubmit(basicFormRef)"
                  >保存</el-button
                >
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <!-- 基本信息 end -->
      <el-divider border-style="dashed" />

      <!-- 安全设置区域 - 并排布局 -->
      <div class="security-settings-container">
        <!-- 绑定邮箱 start -->
        <div class="security-setting-item">
          <BindEmail />
        </div>
        <!-- 绑定邮箱 end -->

        <!-- 修改密码 start -->
        <div class="security-setting-item">
          <UpdatePwd />
        </div>
        <!-- 修改密码 end -->
      </div>
    </div>
    <!-- 左侧信息 end -->
  </el-row>
</template>

<script setup lang="ts">
import { ref, reactive, toRefs, onMounted } from "vue";
import { useUserStore } from "@/store/modules/user";
import { FormRules, FormInstance, ElMessage } from "element-plus";
import { updateUserInfoApi } from "@/api/system/user/user";
import BindEmail from "./BindEmail.vue";
import UpdatePwd from "./UpdatePwd.vue";
const basicFormRef = ref<FormInstance>();
const loading = ref(false);
const state = reactive({
  basic: {
    sex: "",
    avatar: "",
    userName: "",
  },
});
const basicRules = reactive<FormRules>({
  sex: [{ required: true, message: "请选择性别", trigger: "blur" }],
  avatar: [{ required: true, message: "请上传头像", trigger: "blur" }],
});
// 服务器路径
const url = import.meta.env.VITE_BASE_URL;
// 图片上传地址
const uploadURL = url + "upload/file";
// 上传图片使用的token
const userStore = useUserStore();
const headerObj = {
  AccessToken: userStore.token,
};
const handlerAvatarSuccess = (res) => {
  if (res.code == 200) {
    state.basic.avatar = res.fileName;
  }
};
const { basic } = toRefs(state);
// 提交函数
const onBasicSubmit = (formEl: FormInstance | undefined) => {
  if (!formEl) return;

  formEl.validate(async (valid) => {
    loading.value = true;
    if (valid) {
      const { data } = await updateUserInfoApi({ ...state.basic });
      if (data.code == 200) {
        userStore.setUserInfo({
          sex: state.basic.sex,
          avatar: state.basic.avatar,
        });
        ElMessage.success("信息修改成功");
      } else {
        ElMessage.error("信息修改失败");
      }
    } else {
      ElMessage.error("提交失败，请检查表单");
    }
    loading.value = false;
  });
};
onMounted(() => {
  state.basic.userName = userStore.userInfo.userName;
  state.basic.sex = userStore.userInfo.sex;
  state.basic.avatar = userStore.userInfo.avatar;
});
</script>

<style scoped>
.left-box {
  width: 100%;
  height: auto;
  background: white;
  padding: 20px;
  box-sizing: border-box;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.left-box .title {
  color: #333;
  margin-bottom: 20px;
  padding: 0 20px;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  font-size: 20px;
  font-weight: 600;
}

.left-box .title .el-icon {
  color: #409eff;
}

.left-box .set {
  text-align: left;
  padding: 0 20px;
  margin-bottom: 20px;
  line-height: 35px;
}

.left-box .set h4 {
  line-height: 45px;
  color: #333;
  font-weight: 600;
  margin-bottom: 15px;
  border-bottom: 1px solid #f0f0f0;
  padding-bottom: 10px;
}

/* 安全设置容器 */
.security-settings-container {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  width: 100%;
  margin-top: 20px;
}

/* 安全设置项 */
.security-setting-item {
  flex: 1;
  min-width: 300px;
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  overflow: hidden;
  background: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.security-setting-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

/* 响应式布局 */
@media (max-width: 1200px) {
  .security-setting-item {
    flex: 0 0 calc(50% - 10px);
  }
}

@media (max-width: 768px) {
  .security-setting-item {
    flex: 0 0 100%;
  }
  
  .left-box {
    padding: 15px;
  }
  
  .left-box .set {
    padding: 0 10px;
  }
}
</style>