<template>
    <el-form ref="ruleFormRef" :rules="rules" :model="formUser" label-width="80px">
        <el-row>
            <el-col :span="12">
                <el-form-item label="用户名" prop="username">
                    <el-input v-model="formUser.username" placeholder="请输入用户名"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="密码" prop="password">
                    <el-input v-model="formUser.password" placeholder="请输入密码"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="手机号"  prop="phone">
                    <el-input v-model="formUser.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="邮箱"  prop="email">
                    <el-input v-model="formUser.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="24">
                <el-form-item label="备注">
                    <el-input type="textarea" :rows="2" v-model="formUser.remarks" placeholder="请输入备注"></el-input>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button @click="editUser(ruleFormRef)" :loading="subLoading" color="#e99d53" class="edit-user-btn">保存</el-button>
    </div>
</template>

<script setup lang="ts">
import { ElMessage,FormInstance,FormRules } from 'element-plus'
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
    remarks: ''
})
// 获取父组件UserList组件的属性
const props=defineProps(['userInfo'])
const userinfo=props.userInfo
// 给表单填充数据
for(const key in formUser){
    formUser[key]=userinfo[key]
}
// 表单验证
const ruleFormRef=ref<FormInstance>()
const rules=reactive<FormRules>({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' }
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
    ],
    phone: [
        { required: true, message: '请输入手机号', trigger: 'blur' },
        { min: 11, max: 11, message: '长度为11个字符', trigger: 'blur' }
    ],
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
    ]
})

// 更新用户信息
const editUser=async( formEl: FormInstance | undefined )=>{
    if(!formEl) return
    await formEl.validate(async(valid,fields) => {
        subLoading.value=true        
        if(valid){
            const {data} = await editUserApi(formUser)
            if(data.code === 200){
                ElMessage.success(data.msg)
                emit('success')
            }else{
                ElMessage.error(data.msg)

            }
        }else{
            ElMessage.error("提交失败，请检查表单")
        }
    })
    subLoading.value=false
}
// 关闭对话框
const emit=defineEmits(['closeEditUserForm','success'])
const close=()=>{
    emit('closeEditUserForm')
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
