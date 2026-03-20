<template>
    <el-form :model="formMenu" label-width="100px">
        <el-row>
            <el-col :span="12">
                <el-form-item label="菜单名称">
                    <el-input v-model="formMenu.name" placeholder="请输入菜单名称"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="12">
                <el-form-item label="菜单排序">
                    <el-input-number v-model="formMenu.sort" :min="0" placeholder="请输入菜单排序"></el-input-number>
                </el-form-item>
            </el-col>

            <el-col :span="12">
                <el-form-item label="菜单路径">
                    <el-input v-model="formMenu.path"  placeholder="请输入菜单路径"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="12" v-if="formMenu.level==1">
                <el-form-item label="组件路径">
                    <el-input v-model="formMenu.component_name"  placeholder="请输入组件路径"></el-input>
                </el-form-item>
            </el-col>

            <el-col :span="12">
                <el-form-item label="菜单图标" prop="web_icon">
                    <el-input v-model="formMenu.web_icon" placeholder="请选择菜单图标" disabled>
                        <template #append>
                            <el-popover :width="350"
                            trigger="click"
                            popper-class="pure-popper"
                            :popper-options="{placement:'right'}"
                            :visible="selectIconVisible"
                            >
                            <template #reference>
                                <div style="display: flex; justify-content: center; align-items: center;" @click="selectIconVisible=!selectIconVisible">
                                    <el-icon style="cursor: pointer;">
                                        <component style="font-size: 24px;" :is="formMenu.web_icon"></component>
                                    </el-icon>
                                </div>
                            </template>
                                <el-scrollbar height="347px">
                                    <Icon @onChangeIcon="onChangeIcon"/>
                                </el-scrollbar>
                            </el-popover>
                        </template>
                    </el-input>
                </el-form-item>
            </el-col>

        </el-row>
    </el-form>

    <div class="dialog-button-wrap">
        <el-button @click="close">取消</el-button>
        <el-button @click="editMenu" class="edit-menu-btn" color="#e99d53"
        :loading="subLoading"
        >确定</el-button>
    </div>

</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
const props =defineProps(['menuInfo'])
const menuInfo=props.menuInfo

const subLoading = ref(false)
const formMenu = reactive({
    id: 0,
    name: '',
    sort: 0,
    parent_id: 0,
    path: '',
    component_name: '',
    web_icon: '',
    level: 0,
})
// 填充数据
for (const key in formMenu) {
    formMenu[key] = menuInfo[key]
}


</script>

<style  scoped>
.dialog-button-wrap{
    text-align: center;
    margin-top: 20px;
}
.edit-menu-btn{
    color: white;
}
</style>
