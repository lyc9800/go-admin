<template>
    <el-row :gutter="20">
        <!-- 左侧信息 start -->
         <div class="left-box">
            <h3 class="title">
                <el-icon><User/></el-icon>
                个人设置
            </h3>
            <!-- 基本信息 start -->
             <div class="set">
                <h4>基础设置</h4>
                <el-form ref="basicFormRef" :rules="basicRules" status-icon :model="basic">
                    <el-row :gutter="20">
                        <!-- 用户昵称 start -->
                         <el-col :span="8">
                            <el-form-item label="用户账号" prop="userName" style="width:100%">
                                <el-input v-model="basic.userName" disabled></el-input>
                            </el-form-item>
                         </el-col>
                        <!-- 用户昵称 end -->
                        <!-- 性别 -->
                         <el-col :span="8">
                            <el-form-item label="性别" prop="sex" style="width:100%">
                                <el-radio v-model="basic.sex" label="男"></el-radio>
                                <el-radio v-model="basic.sex" label="女"></el-radio>
                            </el-form-item>
                         </el-col>
                         <!-- 性别 end -->
                         <!--头像--> 
                         <el-col :span="5">
                            <el-form-item label="头像" style="margin: auto;">
                                <el-upload class="avatar-uploader" :action="uploadURL" :headers="headerObj"  
                                name="fileResource" :show-file-list="false" :on-success="handlerAvatarSuccess">
                                    <img v-if="basic.avatar" :src="url+'uploadFile/'+basic.avatar" style="width:50px;border-radius: 50px;">
                                    <img v-else src="@/assets/default.png" style="width: 50px;border-radius: 50px;"/>
                                    <span style="margin-left: 10px;font-size: 10px;text-decoration: underline;line-height: 20px;">点击更换</span>
                                </el-upload>
                            </el-form-item>
                         </el-col>
                         <!-- 提交按钮 -->
                         <el-col :span="3">
                            <el-form-item>
                                <el-button plain color="#2fa7b9" style="margin-left: 50px;">保存</el-button>
                            </el-form-item>
                         </el-col>
                    </el-row>
                </el-form>
             </div>
            <!-- 基本信息 end -->
             <el-divider border-style="dashed"/>
             <!-- 绑定邮箱 start -->
             <!-- 绑定邮箱 end -->
             <!-- 修改密码 start -->
             <!-- 修改密码 end -->

         </div>
        <!-- 左侧信息 end -->

         <!-- 右侧信息 start -->

         <!-- 右侧信息 end -->

    </el-row>
</template>

<script setup lang="ts">
import { ref,reactive,toRefs } from 'vue'
import { useUserStore } from '@/store/modules/user'
const  state=reactive({
    basic:{
        userName:'',
        sex:'',
        avatar:''
    },
})
// 服务器路径
const url=import.meta.env.VITE_BASE_URL
// 图片上传地址
const uploadURL=url+'upload/file'
// 上传图片使用的token
const userStore=useUserStore()
const headerObj={
    AccessToken:userStore.token
}
const handlerAvatarSuccess=(res)=>{
    if(res.code==200){
        state.basic.avatar=res.fileName
    }
}
const {basic} = toRefs(state)
</script>

<style  scoped>
.left-box{
    width: 100%;
    height: auto;
    background: white ;
    padding: 20px;
    box-sizing: border-box;
}
.left-box title{
    color: #e99d53;
    margin-bottom: 10px;
    padding: 20px;
    display: inline-flex;
    justify-content: center;
    align-items: center ;
}
.left-box .set{
    text-align: left;
    padding: 0px 20px;
    margin-bottom: 10px;
    color: #8f8f8f;
    line-height: 35px;
}
.left-box .set h4{
    line-height: 45px;
    color: #8f8f8f;
}
</style>
