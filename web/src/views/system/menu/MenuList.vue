<template>
    <el-card>
        <!-- 头部 start -->
         <template #header>
            <div class="card-header">
                <h3>
                    <el-icon style="margin-right: 10px;"><Stamp/></el-icon>菜单管理
                </h3>
                <!-- 搜索 start -->
                 <div class="card-search">
                    <el-row :gutter="10">
                        <!-- <el-col :span="11">
                            <el-input prefix-icon="Search" v-model="searchValue" @keyup.enter="search" placeholder="关键字搜索(回车)"/>
                        </el-col> -->
                        <el-col :span="20">
                            <div class="my-button">
                                <el-button  @click="addMenu" color="#e99d53" plain style="width: 100%;">添加顶级菜单</el-button>                      
                            </div>
                        </el-col>

                        <el-col :span="4" style="display: inline-flex; justify-content: center; align-items: center; cursor: pointer;">
                            <el-icon style="font-size:20px; color:#b3b6bc" @click="refresh"><Refresh/></el-icon>
                        </el-col>

                    </el-row>
                 </div>
                <!-- 搜索 end -->
            </div>
         </template>
        <!-- 头部 end -->
        <!-- 表格区域 start -->
         <div class="table-box">
            <el-table row-key="id" :tree-props="{children: 'sub_menus'}" element-loading-text="数据加载中……" 
            v-loading="loading"
            :data="tableData"
            style="width: 100%; text-align: center;" :cell-style="{textAlign:'center'}"
            :header-cell-style="{fontSize:'15px',background:'#e99d53',color:'white',textAlign:'center'}">
            <el-table-column label="序号" type="index" width="100"/>
            <el-table-column label="菜单名称">
                <template #default="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>
            <el-table-column label="菜单图标">
                <template #default="scope">
                    <span>
                        <el-icon>
                            <component  style="font-size: 24px;" :is="scope.row.web_icon"></component>
                        </el-icon>
                    </span>
                </template>
            </el-table-column>
            <el-table-column label="菜单排序">
                <template #default="scope">
                    <span>{{ scope.row.sort }}</span>
                </template>
            </el-table-column>
            <el-table-column label="菜单类型">
                <template #default="scope">
                    <el-tag  type="success" v-if="scope.row.level===0">目录</el-tag>
                    <el-tag  type="info" v-else-if="scope.row.level===1">菜单</el-tag>
                    <el-tag v-else-if="scope.row.level===2">按钮</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="路由路径" prop="path"/>
            <el-table-column label="组件路径" prop="component_name"/>
            <el-table-column label="操作" width="260">
                <template #default="scope">
                    <el-button v-if="scope.row.level!=2" color="#e99d53" style="color: white;" size="small" 
                    @click="addSubMenu(scope.row.level,scope.row.name)"
                    type="primary">新增子级</el-button>
                    <el-button size="small" @click="editMenu(scope.row.id)">
                        <el-icon><Edit/></el-icon>
                        编辑
                    </el-button>
                    <el-popconfirm width="200px" confirm-button-text="确定" cancel-button-text="取消"
                    icon="Delete" icon-color="#626AEF"
                    :title="'确定删除菜单名为'+scope.row.name+'的菜单吗？'"
                    @confirm="delMenu(scope.row.id)">
                    <template #reference>
                        <el-button size="small" type="danger">
                            <el-icon><Delete/></el-icon>
                            删除
                        </el-button>
                    </template>
                    </el-popconfirm>
                </template>
            </el-table-column>
            </el-table>
         </div>
        <!-- 表格区域 end -->
    </el-card>
    <!--- 新增菜单弹出框 start -->
    <el-dialog align-center  v-model="addMenuFormVisible" width="45%" destroy-on-close>
        <template #header>
            <div class="my-header">
                <el-icon size="26px"><EditPen/></el-icon>
                <h1>{{addTitle}}</h1>
            </div>
        </template>
        <!-- 新增菜单 start -->
         <AddMenu @closeAddMenuForm="closeAddMenuForm" @success="success"/>
        <!-- 新增菜单 end -->
    </el-dialog>
    <!--- 新增菜单弹出框 end -->
</template>

<script setup lang="ts">
import { ref,reactive,toRefs,onMounted } from 'vue'
import { getMenuListApi } from "@/api/system/menu/menu";
import AddMenu from './components/AddMenu.vue';
// import { ElMessage } from 'element-plus'
const state = reactive({
    // 搜索关键字
    searchValue:'',
    loading:false,
    tableData:[]
})
// 获取菜单数据
const loadData = async (state:any) => {
    state.loading = true
    state.tableData = []
    const params = {
        'keyword': state.searchValue
    }   
    const { data } = await getMenuListApi(params)
    
    // state.tableData = data.result
    state.tableData=data.result
    state.loading = false
}
onMounted(() => {
    loadData(state)
})
const {tableData,searchValue}=toRefs(state)
// 定义弹出事件窗

const addMenu = () => {
    addMenuFormVisible.value = true
}
const refresh = () => {
    loadData(state)
}
const addTitle = ref('新增菜单')
// 新增弹窗框状态
const addMenuFormVisible = ref(false)
const closeAddMenuForm = () => {
    addMenuFormVisible.value = false
}
const success=()=>{
    loadData(state)
    closeAddMenuForm()
}
// const search = () => {
//     if(state.searchValue!=null&&state.searchValue!=''){
//         loadData(state)
//         ElMessage({
//             type:'success',
//             message:`关键字"${state.searchValue}"搜索内容如下`
//         })
//     }
// }

</script>

<style  scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.card-header h3{
    display: inline-flex;
    justify-content: center;
    align-items: center;
    color:#e99d53
}
.my-button{
    display: flex;
    justify-content: space-between;
}
.my-header{
    display: flex;
    justify-content: flex-start;
    color: #e99d53;
}
</style>
