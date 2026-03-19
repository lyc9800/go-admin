<template>
    <el-form :model="formMenu" label-width="100px">
        <el-row>
            <el-col :span="12">
                <el-form-item label="菜单名称">
                    <el-input v-model="formMenu.name" placeholder="请输入菜单名称"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="12">
                <el-form-item label="路由路径">
                    <el-input v-model="formMenu.path" placeholder="请输入路由路径"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="12" v-if="formMenu.level===1">
                <el-form-item label="组件路径">
                    <el-input v-model="formMenu.component_name" placeholder="请输入组件路径"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="8">
                <el-form-item label="菜单图标">
                    <el-input v-model="formMenu.web_icon" placeholder="请输入菜单图标"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="8">
                <el-form-item label="菜单排序">
                    <el-input-number v-model="formMenu.sort"  :min="0" placeholder="请输入菜单排序"/>
                </el-form-item>
            </el-col>

            <el-col :span="8">
                <el-form-item label="菜单等级">
                    <el-select v-model="formMenu.level"   placeholder="请输入菜单等级">
                        <el-option label="目录" :value="0"></el-option>
                        <el-option label="菜单" :value="1"></el-option>
                        <el-option label="按钮" :value="2"></el-option>
                    </el-select>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button  @click="addMenu" class="add-menu-btn" color="#e99d53" :loading="subLoading">确定</el-button>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, toRaw, watch } from 'vue'
import {ElMessage} from 'element-plus'
import {addMenuApi} from '@/api/system/menu/menu'

// 定义按钮状态
const subLoading = ref(false)
const formMenu = reactive({
    name: '',
    path: '',
    component_name: '',
    web_icon: 'Search',
    sort: 0,
    level: 0,
})
// 新增菜单
const addMenu =async () => { 
    subLoading.value = true
    const {data} = await addMenuApi(formMenu)
    if(data.code === 200){ 
        ElMessage.success(data.msg)
        emit('success')
    }else{ 
        ElMessage.error(data.msg)
    }
    subLoading.value = false
}
// 定义事件
const emit = defineEmits(['closeAddMenuForm','success'])
const close = () => { 
    emit('closeAddMenuForm')
}
</script>

<style  scoped>
.dialog-button-wrap{
    text-align: center;
    margin-top: 20px;
}
.add-menu-btn{
    color: white;
}
</style>
