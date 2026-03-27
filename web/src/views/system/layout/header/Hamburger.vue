<template>
    <el-breadcrumb class="app-breadcrumb" separator="/">
        <transition-group name="breadcrumb" mode="out-in">
            <el-breadcrumb-item :to="{path:'/home'}" key="home" v-if="matched[0].meta.title!=='首页'">
                <div class="breadcrumb-item">
                    <span class="breadcrumb-item">首页</span>
                </div>
            </el-breadcrumb-item>
            <el-breadcrumb-item v-for="(item,index) in matched" :key="item.name">
                <span v-if="item.redirect ==='noRedirect' || index==matched.length-1" class="no-redirect">
                    {{item.meta.title}}
                </span>
                <a v-else @click.prevent="handleLink(item)">{{ item.meta.title }}</a>
            </el-breadcrumb-item>
        </transition-group>
    </el-breadcrumb>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRoute,useRouter } from 'vue-router'
const route = useRoute();
const router = useRouter();
const handleLink=(item:any)=>{
    try {
        if (!item?.path) {
            console.warn('路由路径不存在:', item)
            return
        }
        router.push(item.path).catch(err => {
            console.error('路由跳转失败:', err)
        })
    } catch (error) {
        console.error('跳转异常:', error)
    }
}
const matched=computed(()=>route.matched.filter(item=>item.meta && item.meta.title))
</script>

<style  scoped>

</style>