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
                            <el-icon style="font-size: 20px;" color="#b3b6bc" @click="refresh"><Refresh/></el-icon>
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
                    <span>{{ scope.row.role_name }}</span>
                </template>
            </el-table-column>

            <el-table-column label="角色排序">
                <template #default="scope">
                    <span>{{ scope.row.sort }}</span>
                </template>
            </el-table-column>

            <el-table-column label="是否管理员">
                <template #default="scope">
                    <el-switch
                        v-model="scope.row.is_admin"
                        :active-value="1"
                        :inactive-value="0"
                        style="--el-switch-on-color:#e99d53"
                        :before-change="() => handleSwitchChange(scope.row)"
                        @click="(e) => console.log('🖱️ 开关被点击！', e, scope.row)"
                        />
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

     <el-dialog align-center v-model="addRoleFormVisible" width="45%" destroy-on-close>
        <template #header>
            <div class="my-header">
                <el-icon size="26px"><EditPen/></el-icon>
                <h1>新增角色</h1>
            </div>
        </template>
            <!-- 添加角色 start-->
             <AddRole @closeAddRoleForm="closeAddRoleForm" @success="success"/>
            <!-- 添加角色 end-->
     </el-dialog>
     <!-- 新增角色弹窗 start -->
      <el-dialog align-center v-model="editRoleFormVisible" width="50%" destroy-on-close>
        <template #header>
            <div class="my-header">
                <el-icon size="26px"><EditPen/></el-icon>
                <h1>编辑角色</h1>
            </div>
        </template>
            <!-- 编辑角色 start-->
             <EditRole :roleInfo="roleInfo" @closeEditRoleForm="closeEditRoleForm" @success="success"/>
      </el-dialog>
</template>

<script setup lang="ts">
import { formatTime } from '@/utils/date';
import {exportExcel} from '@/utils/exportExcel'
import {ref,reactive,toRefs,onMounted} from 'vue'
import {getRoleListApi,delRoleApi,getRoleApi,changeIsAdminApi} from '@/api/system/role/role'
import {ElMessage,ElMessageBox} from 'element-plus'
import AddRole from './components/AddRole.vue';
import EditRole from './components/EditRole.vue';

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
        'keyword':state.searchValue,
    }
    const {data} =await getRoleListApi(params)
    state.tableData=data.result.list
    state.total=data.result.total
    state.loading=false
}
onMounted(()=>{
    loadData(state)
})
// 关闭新增角色弹窗
const closeAddRoleForm=()=>{
    addRoleFormVisible.value=false
}
// 编辑角色弹窗
const editRoleFormVisible=ref(false )
const roleInfo=ref()
const editRole= async(id:number)=>{
    const {data } =await getRoleApi(id) 
    roleInfo.value=data.result  
    editRoleFormVisible.value=true
}
const closeEditRoleForm=()=>{
    editRoleFormVisible.value=false
}   
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
const Nindex = (index)=>{
    const page=state.pageIndex
    const pageSize=state.pageSize
    return index+1+(page-1)*pageSize
}
// 新增角色弹出框状态
const addRoleFormVisible=ref(false)

const addRole=()=>{
    addRoleFormVisible.value=true
}
const refresh=()=>{
    state.searchValue = ''
    loadData(state)
}   
const search=()=>{
    if(state.searchValue!=null&&state.searchValue!=''){
        loadData(state)
        ElMessage({
            type:'success',
            message:`关键字"${state.searchValue}"搜索内容如下`
        })
    }
}
// 新增角色成功刷新界面
const success=()=>{
    loadData(state)
    closeAddRoleForm()
    closeEditRoleForm()
}
const delRole=async(id:number) => { 
    const {data}=await delRoleApi(id)
        if(data.code===200){
            ElMessage({
                type:'success',
                message:'删除成功'
            })
            await loadData(state)
        }else{
            ElMessage({
                type:'error',
                message:'删除失败'
            })
        }
}   
const column=[
    {name:'id',label:'角色ID'},
    {name:'name',label:'角色名称'},
    {name:'remarks',label:'角色描述'},
    {name:'sort',label:'排序'},
    {name:'is_admin',label:'是否是管理员'}
]
const exportExcelAction=()=>{
    exportExcel({
        column,
        data:state.tableData,
        fileName:'角色信息',
        format:'xlsx',
        autowidth:true,
    })
}
const handleSwitchChange = async (row) => {
  console.log('✅ before-change 被调用！row:', row);
  console.log('✅ is_admin 原始值：', row.is_admin, '类型：', typeof row.is_admin);

  try {
    const confirmResult = await ElMessageBox.confirm(
      "确定修改当前角色的管理员状态吗？",
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        // cancelButtonClass: 'btnConfirm', // 👈 先注释掉
        autofocus: false
      }
    );

    console.log('✅ 用户点击了确定');
    // 继续执行 API 调用
    const newIsAdmin = row.is_admin === 1 ? 0 : 1;
    await changeIsAdminApi(row.id, newIsAdmin);
    row.is_admin = newIsAdmin; // 更新本地数据
    ElMessage.success('状态已更新');
  } catch (error) {
    if (error === 'cancel') {
      console.log('❌ 用户取消了操作');
    } else {
      console.error('⚠️ 其他错误：', error);
    }
  }
};

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
.my-header{
    display: flex;
    justify-content: flex-start;
    color: #e99d53;
}
:deep(.el-pagination.is-background .el-pager li:not(.is-disabled).is-active){
    background-color: #e99d53;
}
</style>
