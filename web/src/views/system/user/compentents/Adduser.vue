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
                <el-form-item label="手机号" prop="phone">
                    <el-input v-model="formUser.phone" placeholder="请输入手机号码"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="12">
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="formUser.email" placeholder="请输入邮箱地址"></el-input>
                </el-form-item>
            </el-col>
            <el-col :span="24">
                <el-form-item label="备注">
                    <el-input :rows="2" type="textarea"  v-model="formUser.remark" placeholder="请输入备注"></el-input>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button @click="addUser(ruleFormRef)" :loading="subLoading" type="success" color="#e99d53" class="add-user">保存</el-button>
    </div>
</template>

<script lang='ts' setup>
import { ElMessage,FormInstance,FormRules } from 'element-plus'
import {ref,reactive} from 'vue'
import { addUserApi } from '@/api/system/user/user'

// 按钮状态
const subLoading = ref(false)
// 表单数据对象
const formUser = reactive({
    username: '',
    password: '123456',
    phone: '',
    remark: '',
    email: '',
})
// 表单验证规则
const ruleFormRef  = ref<FormInstance>()
const rules=reactive<FormRules>({
    username: [{ required: true, message: '用户名不能为空', trigger: 'blur' },],
    password: [{ required: true, message: '密码不能为空', trigger: 'blur' },],
    phone: [{ required: true, message: '手机号不能为空', trigger: 'blur' },],
    email: [{ required: true, message: '邮箱不能为空', trigger: 'blur' },],
})
// 添加用户
const addUser= async(formEl:FormInstance|undefined)=>{
    console.log('1. 函数开始执行，formEl:', formEl)
    if(!formEl) {
        console.log('2. formEl 为空，直接返回')
        return
    }
    await formEl.validate(async(valid,fields) => {
            subLoading.value=true
            if(valid){
                // 验证通过
                const {data} =await addUserApi(formUser)
                if(data.code===200){
                    ElMessage.success(data.msg)
                    emit('success')
                }else{
                    ElMessage.error(data.msg)
                }
            }else{
                // 验证失败
                ElMessage.error('请检查表单数据')
            }
    })
    subLoading.value=false
}
// 关闭窗口
const emit =defineEmits(['closeAdduserForm','success'])
const close=()=>{
    emit('closeAdduserForm')
}
</script>

<style scoped>
.dialog-button-wrap{
    text-align: center;
    margin-top: 20px;
}
.add-user{
    color: white;
}
</style>