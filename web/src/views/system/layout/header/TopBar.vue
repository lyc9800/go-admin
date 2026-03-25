<template>
  <div class="main">
    <div class="tool-left">
      <CollapseIcon/>
      <Hamburger/>
    </div>
    <div class="linkBox">
      <!-- 登陆用户信息 start -->
        <el-popover :width="300">
          <!-- 用户头像 start-->
            <template #reference>
              <el-link :underline="false">
                <img v-if="userInfo.avatar" :src="url+'uploadFile/'+userInfo.avatar" style="width:40px;border-radius: 50px;">
                <img v-else src="@/assets/default.png" style="width:40px;border-radius: 50px;">
                <span>{{userInfo.userName}}</span>
              </el-link>
            </template>
          <!-- 用户头像 end-->
          <!-- 用户信息 start-->
            <template #default>
              <div style="display:flex;gap:16px;flex-direction: column;">
                <div class="info-card">
                  <!-- 用户头像 start-->
                   <img v-if="userInfo.avatar" :src="url+'uploadFile/'+userInfo.avatar">
                   <img v-else src="@/assets/default.png">
                   <p>用户名:{{ userInfo.userName }}</p>
                   <p>Email:{{ userInfo.email }}</p>
                   <p>注册时间:{{ formatTime(userInfo.CreatedAt,'yyyy-MM-dd') }}</p>
                   <p>注册天数:{{ calculateDays(userInfo.CreatedAt) }}天</p>
                  <!-- 用户头像 end-->
                </div>
                <div class="info-card-desc" style="margin: 0;">
                    <div style="float:left;width: 75px;padding: 10px;border-right: 1px;text-align: center;">
                    </div>
                </div>
              </div>
            </template>
          <!-- 用户信息 end-->
        </el-popover>
      <!-- 登陆用户信息 end -->
      <!-- 退出系统 start -->
       <el-popconfirm confirm-button-text="确定" cancel-button-text="取消" :icon="SwitchButton"
       icon-color="#30bcd7" title="确定要退出吗?" @confirm="exit" >
        <template #reference>
          <el-link :underline="false"  class="logout-btn">
            <el-icon>
              <SwitchButton/>
            </el-icon>
          </el-link>
        </template>
       </el-popconfirm>
      <!-- 退出系统 end -->
    </div>
  </div>
  <!-- tabs多标签页 start -->
   <TabsView/>
  <!-- tabs多标签页 end-->
</template>

<script setup lang="ts">
import { SwitchButton } from '@element-plus/icons-vue'
import CollapseIcon from './CollapseIcon.vue';
import Hamburger from './Hamburger.vue';
import TabsView from '@/views/system/layout/tags/Index.vue';
import { useUserStore } from '@/store/modules/user'
import { formatTime,calculateDays } from '@/utils/date';

// 获取登陆用户信息
const {userInfo}=useUserStore()
console.log('当前 userInfo 对象:', userInfo)
// 退出系统
const exit =()=>{
  // 清除用户登陆信息
  window.localStorage.removeItem("userStore")
  // 清除用户对应的菜单数据
  window.localStorage.removeItem("menuState")
  window.location.href='/'
}
// 服务器路径
const url=import.meta.env.VITE_BASE_URL
</script>

<style  scoped>
.main{
  height: 70px;
  box-shadow: rgb(0 0 0 /10%) 0px 0px 0px;
  background: width;
}
.linkBox{
  height: 100%;
  display: flex;
  align-items: center;
  text-align: center;
  float: right;
  
}
.linkBox .el-link{
  margin-right: 25px;
  color: #8c8c8c;
}
.linkBox .el-link:hover{
  color: #30bcd7;
}
.linkBox .el-link span{
  margin-left: 8px;
}
/* 鼠标滑向头像的样式 */
.info-card img{
  width: 55px;
  margin: 0px 15px 0px 0px;
  border-radius: 50px;
  float: left;
}
.info-card p{
  margin-right: 0px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;

}
.main{
  display: flex;
  justify-content: space-between;
  height: 70px;
  box-shadow: rgb(0 0 0 /10%) 0px 0px 10px ;
  background: white;
}
.tool-left{
  display: flex;
  align-items: center;
  margin-left: 10px;
}
:deep(.main-tabs-view) {
  background: white !important;
}

</style>

