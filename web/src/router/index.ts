import Nprogress from '@/config/nprogress'

// 1.导入vue-router模块
import {createRouter,createWebHistory} from 'vue-router'
// 2.定义一些路由地址，每个都需要映射到一个组件
const routes=[{
    path:'/',
    name:'Login',
    meta:{title:'后台管理系统-登陆'},
    component:()=> import ('../views/Login.vue')
}]
// 3.创建路由实例并传递'routes'配置
const router=createRouter({
    history:createWebHistory(),
    routes:routes
})
// 路由拦截守卫
router.beforeEach(async(to,form,next)=>{
    // 1.Nprogress开始
    Nprogress.start()
    next()
})
// 路由跳转结束
router.afterEach(()=>{
    Nprogress.done()
})
// 路由跳转失败
router.onError(error=>{
    Nprogress.done()
    console.warn("路由错误",error.message)
})
export default router