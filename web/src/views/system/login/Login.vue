<template>
    <div class="login-container">
        <div class="login-box">
            <!-- 左边背景图 start -->
             <div class="login-left">
                <img src="@/assets/login/side-logo.png" alt="">
             </div>
            <!-- 左边背景图 end -->

            <!-- 登录表单 start -->
             <div class="login-form">
                <!-- 标题和logo start -->
                 <div class="login-title">
                    <img src="@/assets/login/logo02.png" alt="">
                 </div>
                 <h2 class="title">后台管理系统</h2>
                <!-- 标题和logo end -->
                <!-- 登录表单组件 start -->
                 <LoginForm/>
                <!-- 登录表单组件 end -->
             </div>
            <!-- 登录表单 end -->
        </div>
    </div>
</template>
<script setup lang="ts">
import LoginForm from './components/LoginForm.vue'
</script>
<style  scoped>

</style>
