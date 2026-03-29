<template>
    <el-card>
        <!-- 头部区域 start -->
        <template #header>
            <div class="card-header">
                <h3>
                    <el-icon style="margin-right: 10px;"><Document/></el-icon>日志管理
                </h3>
                <!-- 搜索 start -->
                 <div class="card-search">
                    <el-row :gutter="10">
                        <el-col :span="14">
                            <el-input prefix-icon="Search" v-model="searchValue"
                            @keyup.enter="search" placeholder="关键字搜索(回车)"/>
                        </el-col>

                        <el-col :span="7">
                            <el-button plain style="width: 100%;" color="#2fa7b9" @click="exportExcelAction">导出 Excel</el-button>
                        </el-col>

                        <el-col :span="3" style="display: inline-flex; justify-content: center; align-items: center; cursor: pointer;">
                            <el-icon style="font-size:20px;color: #b3b6bc;" @click="refresh"><Refresh/></el-icon>
                        </el-col>
                    </el-row>
                 </div>
                <!-- 搜索 end -->
            </div>
        </template>
        <!-- 头部区域 end -->
        <!-- 内容区域 start -->
         <div class="table-box">
            <el-table element-loading-text="数据加载中……" v-loading="loading" :data="tableData"
            style="width: 100%;text-align: center;"
            :cell-style="{textAlign:'center'}"
            :header-cell-style="{fontSize:'15px',background:'#e99d53',color:'white',textAlign:'center'}">
            <el-table-column type="index" label="序号" width="100px" :index="Nindex"/>
            <el-table-column label="IP">
                <template #default="scope">
                    <span v-if="scope.row.ip=='::1'">IP地址</span>
                    <span v-else>{{scope.row.province}}-{{scope.row.city}}</span>
                </template>
            </el-table-column>
            <el-table-column label="请求状态">
                <template #default="scope">
                    <span>{{scope.row.status_code}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作方式">
                <template #default="scope">
                    <span>{{scope.row.http_method}}</span>
                </template>
            </el-table-column>
            <el-table-column label="请求方法">
                <template #default="scope">
                    <span>{{scope.row.class_method}}</span>
                </template>
            </el-table-column>
            <el-table-column label="执行时间(毫秒)">
                <template #default="scope">
                    <el-tag>{{scope.row.use_time}}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="浏览器">
                <template #default="scope">
                    <el-tag>{{scope.row.browser}}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="访问时间">
                <template #default="scope">
                    <sapn>{{formatTime(scope.row.created_at,'yyyy-MM-dd HH:mm:sss')}}</sapn>
                </template>
            </el-table-column>
            </el-table>
         </div>
        <!-- 内容区域 end -->
        <!-- 分页区域 start -->
            <el-pagination
                background
                layout="total,sizes, prev, pager, next,jumper"
                v-model:page-size="pageSize" :page-sizes="[10,20,30,40]"     
                @current-change="changePage" @size-change="changeSize"     
                :total="total">
            </el-pagination>
        <!-- 分页区域 end -->
    </el-card>
</template>

<script lang='ts' setup>
import { ref, reactive, toRefs ,onMounted} from 'vue'
import { getLogListApi } from '@/api/system/log/log'
import { formatTime } from '@/utils/date';
import {exportExcel} from '@/utils/exportExcel'
const state = reactive({
    loading:false,
    tableData:[],
    total:0,
    pageIndex:1,
    pageSize:10,
    searchValue:''
})
const {tableData,pageSize,loading,total,searchValue} =toRefs(state)
const search = () => {
    console.log('search');
}
const loadData = async(state:Object) => { 
    state.loading=true
    state.tableData=[]
    const params = {
        'pageIndex':state.pageIndex,
        'pageSize':state.pageSize,
    }
    const {data} =await getLogListApi(params)
     console.log('📦 API返回的完整数据:', data)
    state.tableData=data.result.list
    state.total=data.result.count
    state.loading=false
}
onMounted(()=>{
    loadData(state)
})
const refresh=()=>{
    state.searchValue = ''
    loadData(state)
}   
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
const Nindex = (index)=>{
    const page=state.pageIndex
    const pageSize=state.pageSize
    return index+1+(page-1)*pageSize
}
const column=[
    {name:'id',label:'序号'},
    {name:'ip',label:'IP地址'},
    {name:'request_uri',label:'请求方法'},
    {name:'browser',label:'浏览器'},
    {name:'created_at',label:'访问时间'}
]
const exportExcelAction=()=>{
    exportExcel({
        column,
        data:state.tableData,
        fileName:'操作日志',
        format:'xlsx',
        autowidth:true,
    })
}
</script>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.card-header h3{ 
    display: flex;
    justify-content: center;
    align-items: center;
    color: #e99d53;
}
.el-pagination{
    margin-top: 20px;
    justify-content: center;
}
:deep(.el-pagination.is-background .el-pager li:not(.disabled).is-active){
    background-color: #e99d53;
}
</style>