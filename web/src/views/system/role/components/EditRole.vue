<template>
    <el-form ref="ruleFormRef" :rules="rules"  :model="formRole" label-width="100px">
        <el-row>
            <el-col :span="8">
                <el-form-item label="角色名称" prop="role_name">
                    <el-input v-model="formRole.role_name" placeholder="请输入角色名称"/>
                </el-form-item>
            </el-col>

            <el-col :span="8">
                <el-form-item label="角色排序" prop="sort">
                    <el-input-number v-model="formRole.sort" placeholder="请输入角色排序"/>
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
                <el-form-item label="角色备注">
                    <el-input v-model="formRole.remarks" type="textarea" :rows="2" placeholder="请输入角色备注"/>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button @click="editRole(ruleFormRef)" :loading="subLoading" color="#e99d53" class="edit-role-btn">确定</el-button>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive,onMounted } from 'vue'
import { editRoleApi } from '@/api/system/role/role'
import { ElMessage,FormInstance,FormRules } from 'element-plus'
import type {ElTree} from "element-plus"
import { getMenuListApi } from "@/api/system/menu/menu" 
// 获取父组件传递的参数
const props = defineProps(['roleInfo'])
const emit = defineEmits(['closeEditRoleForm','success'])
const roleInfo = props.roleInfo
// 按钮状态
const subLoading = ref(false)
const formRole = reactive({
    id:'',
    role_name: '',
    sort: 0,
    is_admin: 0,
    remarks: '',
    menuId: []
}) 
// 给表单填充数据
for (const key in formRole) {
    formRole[key] = roleInfo[key]
}
// 编辑角色
const editRole = async(formEl:FormInstance|undefined) => {
    if(!formEl) return
    await formEl.validate(async(valid) => {
        subLoading.value = true
        if(valid) {
            formRole.menuId = treeRef.value!.getCheckedKeys(false)
            const {data} =await editRoleApi(formRole)
            if (data.code === 200) {
                ElMessage.success(data.msg)
                emit('success')
            } else {
                ElMessage.error(data.msg)
            }
        }else { 
            ElMessage.error('请填写正确的数据')
        }
           
})

    
    subLoading.value = false
}
const close = () => {
    emit('closeEditRoleForm')
}
const ruleFormRef = ref<FormInstance>()
const rules = reactive<FormRules>({
    name: [
        { required: true, message: '请输入角色名称', trigger: 'blur' },
        { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' }
    ],
    sort: [
        { required: true, message: '请输入角色排序', trigger: 'blur' },
    ]
})
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
.edit-role-btn{
    color: white;
}
</style>
