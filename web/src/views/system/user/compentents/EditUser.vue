<template>
    <el-form :model="formUser" label-width="80px">
        <el-row>
            <el-col :span="12">
                <el-form-item label="用户名">
                    <el-input v-model="formUser.username" placeholder="请输入用户名"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="密码">
                    <el-input v-model="formUser.password" placeholder="请输入密码"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="手机号">
                    <el-input v-model="formUser.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="邮箱">
                    <el-input v-model="formUser.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="24">
                <el-form-item label="备注">
                    <el-input type="textarea" :rows="2" v-model="formUser.remark" placeholder="请输入备注"></el-input>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button @click="editUser" :loading="subLoading" color="#e99d53" class="edit-user-btn">保存</el-button>
    </div>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus'
import {ref, reactive} from 'vue'
import {editUserApi} from '@/api/system/user/user'
// 按钮状态
const subLoading = ref(false)
// 表单数据对象
const formUser = reactive({
    id: 0,
    username: '',
    password: '',
    phone: '',
    email: '',
    remark: ''
})
// 获取父组件UserList组件的属性
const props=defineProps(['userInfo'])
const userinfo=props.userInfo
// 给表单填充数据
for(const key in formUser){
    formUser[key]=userinfo[key]
}
// 更新用户信息
const editUser=async()=>{
    subLoading.value=true
    const {data} = await editUserApi(formUser)
    if(data.code === 200){
        ElMessage.success(data.msg)
    }else{
        ElMessage.error(data.msg)

    }
    subLoading.value=false
}
</script>

<style  scoped>
.dialog-button-wrap{
    text-align: center;
    margin-top: 20px;
}
.edit-user-btn{
    color: white;
}
</style>
