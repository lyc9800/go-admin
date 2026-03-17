<template>
    <el-card >
        <!-- 头部 start -->
         <template #header>
            <div class="card-header">
                <h3>
                    <el-icon style="margin-right: 10px;"><stamp/></el-icon>角色管理
                </h3>
                <!-- 搜索框 start -->
                 <div class="card-search">
                    <el-row :gutter="10">
                        <el-col :span="10">
                            <el-input prefix-icon="Search" v-model="searchValue" @keyup.enter="search" placeholder="关键字搜索(回车)"></el-input>
                        </el-col>
                        <el-col :span="11">
                            <div class="my-button">
                                <el-button plain style="width: 100%;" color="#e99d53" @click="addRole">
                                    <el-icon><plus/></el-icon>
                                    添加角色
                                </el-button>
                                <el-button plain style="width: 100%;" color="#076613" @click="exportExcelAction">
                                    <el-icon><download/></el-icon>
                                    导出Excel
                                </el-button>
                            </div>
                        </el-col>

                        <el-col :span="3" style="display: inline-flex;justify-content: center;align-items: center;cursor: pointer;">
                            <el-icon style="font-size: 20px;" color="#b3b6bc" @click="refresh"><refresh/></el-icon>
                        </el-col>
                    </el-row>
                 </div>
                <!-- 搜索框 end -->
            </div>
         </template>
        <!-- 头部 end -->
         <!-- 表格 start -->
          <div class="table-box">
            <el-table element-loading-text="数据加载中……" v-loading="loading" :data="tableData"
            style="width: 100%; text-align: center;"
            :cell-style="{textAlign:'center'}"
            :header-cell-style="{fontSize:'15px',background:'#e99d53',color:'white',textAlign:'center'}">
            <el-table-column label="序号" width="100" type="index" :index="Nindex"/>
            <el-table-column label="角色名称">
                <template #default="scope">
                    <span>{{ scope.row.name }}</span>
                </template>
            </el-table-column>

            <el-table-column label="角色排序">
                <template #default="scope">
                    <span>{{ scope.row.sort }}</span>
                </template>
            </el-table-column>

            <el-table-column label="是否管理员">
                <template #default="scope">
                    <el-switch v-model="scope.row.is_admin" 
                    :active-value="1"
                    :inactive-value="0"
                    style="--el-switch-on-color:#e99d53">
                </el-switch>
                </template>
            </el-table-column>

            <el-table-column label="创建时间">
                <template #default="scope">
                    <span>{{ formatTime(scope.row.created_at,'yyyy-MM-dd') }}</span>
                </template>
            </el-table-column>

            <el-table-column label="更新时间">
                <template #default="scope">
                    <span>{{ formatTime(scope.row.created_at,'yyyy-MM-dd') }}</span>
                </template>
            </el-table-column>

            <el-table-column label="操作">
                <template #default="scope">
                    <el-button  size="small" @click="editRole(scope.row.id)">
                        <el-icon><Edit/></el-icon>
                        编辑
                    </el-button>
                    <el-popconfirm width="200px" confirm-button-text="确定"
                    cancel-button-text="取消" icon="Delete"
                    icon-color="#626AEF"
                    :title="'确定删除角色名为'+scope.row.name+'的角色吗？'"
                    @confirm="delRole(scope.row.id)">
                    <template #reference>
                        <el-button size="small" tyle="danger">
                            <el-icon><Delete/></el-icon>
                            删除
                        </el-button>
                    </template>
                    </el-popconfirm>
                </template>
            </el-table-column>

            </el-table> 
          </div>
         <!-- 表格 end -->

         <!-- 分页 start -->
          <el-pagination background layout="total,sizes, prev, pager, next,jumper"
                :total="total"
                v-model:page-size="pageSize"
                :page-sizes="[10,20,30,40]" @current-change="changePage" @size-change="changeSize"/>
          <!-- 分页 end -->
    </el-card>
</template>

<script setup lang="ts">
import { formatTime } from '@/utils/date';
import {ref,reactive,toRefs,onMounted} from 'vue'
import {getRoleListApi} from '@/api/system/role/role'
const state=reactive({
    loading:false,
    tableData:[],
    total:0,
    pageIndex:1,
    pageSize:10,
    searchValue:'',
})
const loadData= async (state:any)=>{ 
    state.loading=true
    state.tableData=[]
    const params={
        'pageIndex':state.pageIndex,
        'pageSize':state.pageSize,
    }
    const {data} =await getRoleListApi(params)
    state.tableData=data.result.list
    state.total=data.result.total
    state.loading=false
}
onMounted(()=>{
    loadData(state)
})
const {tableData,total,pageSize,searchValue,loading}=toRefs(state)


const changeSize=(val)=>{
    state.pageSize=val
    state.pageIndex = 1 
    loadData(state)
}
// 切换页码执行事件
const changePage = (val)=>{
    state.pageIndex=val
    loadData(state)
}
const search=()=>{
    console.log('search')
}
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
    color: #e99d53;
}
.my-button{
    display: flex;
    justify-content: space-between;
}
.el-pagination{
    justify-content: center;
    margin-top: 20px;
}
:deep(.el-pagination.is-background .el-pager li:not(.is-disabled).is-active){
    background-color: #e99d53;
}
</style>
