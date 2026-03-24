<template>
    <el-form ref="ruleFormRef" :rules="rules" :model="formRole" label-width="100px">
        <el-row>
            <el-col :span="8">
                <el-form-item label="角色名称" prop="name">
                    <el-input v-model="formRole.name" placeholder="请输入角色名称"/>
                </el-form-item>
            </el-col>

            <el-col :span="8">
                <el-form-item label="角色排序" prop="sort">
                    <el-input-number v-model="formRole.sort" :min="0" placeholder="请输入角色排序"/>
                </el-form-item>
            </el-col>

            <el-col :span="8">
                <el-form-item label="是否为管理员">
                    <el-switch v-model="formRole.is_admin" :active-value="1" :inactive-value="0"
                    style="--el-switch-on-color:#e99d53"/>
                </el-form-item>
            </el-col>

            <el-col :span="24">
                <el-form-item label="分配菜单">
                    <el-tree
                    ref="treeRef"
                    :data="menus"
                    :props="treeProps"
                    multiple
                    :render-after-expand="false"
                    node-key="id"
                    show-checkbox
                    :default-checked-keys="formRole.menuId"
                    check-on-click-node
                    style="display: block;width: 100%;"/>
                </el-form-item>
            </el-col>


            <el-col :span="24">
                <el-form-item label="角色描述">
                    <el-input type="textarea"  :rows="2" v-model="formRole.remarks" placeholder="请输入角色描述"/>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button @click="addRole(ruleFormRef)" class="add-role-btn" color="#e99d53" :loading="subLoading">确定</el-button>
    </div>
</template>

<script setup lang="ts">
import  {ref,reactive,onMounted} from "vue"
import {ElMessage,FormInstance,FormRules} from "element-plus"
import type {ElTree} from "element-plus"
import {addRoleApi} from "@/api/system/role/role"
import { getMenuListApi } from "@/api/system/menu/menu" 
// 表单实例对象
const ruleFormRef = ref<FormInstance>()
const rules = reactive<FormRules>({
    name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
    sort: [{ required: true, message: '请输入角色排序', trigger: 'blur' }],
})
// 按钮状态
const subLoading = ref(false)
// 表单数据
const formRole = reactive({
    name: '',
    sort: 0,
    is_admin: 0,
    remarks: '',
    menuId: []
})
const addRole=async(formEl:FormInstance|undefined)=>{ 
    if(!formEl) return
    await formEl.validate(async(valid) => {
        subLoading.value=true
        if(valid){
            formRole.menuId = treeRef.value!.getCheckedKeys(false)
            console.log('提交的数据:', JSON.parse(JSON.stringify(formRole)))  // ✅ 打印
            const {data} = await addRoleApi(formRole)
            if(data.code===200){
                ElMessage.success(data.msg)
                emit('success')
            }else{
                ElMessage.error(data.msg)
            }
        }else{
            ElMessage.error('请填写必填项')
        }
    })
    subLoading.value=false
}
const emit=defineEmits(['closeAddRoleForm', 'success'])
const close=()=>{
    emit('closeAddRoleForm')
}
// 树形菜单对象
const treeRef = ref<InstanceType<typeof ElTree>>()
const menus = ref()
const treeProps = { 
    children: 'sub_menus',
    label: 'name'
}
const menuList = async () => {
    const { data } = await getMenuListApi()
    menus.value = data.result
}
onMounted(() => {
    menuList()
})
</script>

<style  scoped>
.dialog-button-wrap{
    text-align: center;
    margin-top: 20px;
}
.add-role-btn{
    color:white
}
</style>
