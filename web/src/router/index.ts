import Nprogress from '@/config/nprogress'
// 1.导入vue-router模块
import {createRouter,createWebHashHistory} from 'vue-router'
import {useMenuStore} from '@/store/modules/menu'
import {useUserStore} from '@/store/modules/user'
// 定义路由和组件映射关系
const modules=import.meta.glob("@/views/**/**.vue")
console.log('🔍 所有可用模块键:', Object.keys(modules))
// 2.定义一些路由地址，每个都需要映射到一个组件
const routes=[{
    path:'/',
    redirect:'/login',
    meta:{title:'后台管理系统-登陆'},
    component:()=> import ('../views/system/login/Login.vue')
},
{
    path:'/login',
    name:'Login',
    meta:{title:'后台管理系统-登陆'},
    component:()=> import ('../views/system/login/Login.vue')
},
{
    path: '/index',
    name:'Index',
    component:()=>import('@/views/system/layout/Index.vue'),
    redirect:'/home',
    children:[
        {
            path:'/home',
            name:'Home',
            meta:{title:'首页',icon:'House'},
            component: ()=> import('@/views/system/home/Index.vue')
        }
    ]
}
]
// 3.创建路由实例并传递'routes'配置
const router=createRouter({
    history:createWebHashHistory(),
    routes:routes
})
// 防止首次或者刷新界面路由失效
let registerRouteFresh=true
// 设置白名单
const whiteList=['/login']
// 路由拦截守卫
router.beforeEach(async(to,from,next)=>{
    console.log('🚦 路由守卫开始，目标路径:', to.path)
    console.log(`🚦 路由守卫触发: ${from.path || 'null'} -> ${to.path}`)
    // 1.Nprogress开始
    Nprogress.start()

    // 如果是白名单的路径，直接放行
    const some=whiteList.some(function(item){ 
        console.log('✅ 白名单路径，直接放行')
        return to.path.indexOf(item)!==-1
 
    })
    if(some){
        return next()
    }else{
        // 判断是否已经登陆
        const userStore=useUserStore()
        if(userStore.token==''|| userStore.token==null){
            console.log('❌ 未登录，跳转到登录页')
            return next({path:`/login?redirect=${to.path}`,replace:true})
        }
    }


    // 获取菜单信息
    const menuStore=useMenuStore()
    console.log('🔄 检查菜单Store状态:', {
    routers长度: menuStore.routers.length,
    routers内容: menuStore.routers,
    store对象: menuStore
    })

    // 如果routers为空，获取菜单数据
    if(menuStore.routers.length==0){
        console.log('🔄 开始获取菜单数据...')
        await menuStore.generateRouter()
    }
    console.log('动态路由数据:', JSON.parse(JSON.stringify(menuStore.routers)))
    // 生成动态路由 start
    menuStore.routers.forEach((item:any)=>{
        // 组装动态路由地址 start
        let myRoute:any={}
        myRoute={
            path:item.path,
            name:item.name,
            meta:{
                icon:item.web_icon,
                title:item.name
            },
            component:()=>import('@/views/system/layout/Index.vue')
        }
        myRoute.children=[]
        if(item.level===1 && item.component_name.length!=0){
            myRoute.children.push({
                path:item.path,
                name:item.name,
                meta:{
                    icon:item.web_icon,
                    title:item.name
                },
                component:modules[`@/views/${item.component_name}`]
            })
        }
        if(item.sub_menus){
            item.sub_menus.forEach((subItem:any)=>{
                if(subItem.path){
                    myRoute.children.push({
                         path:subItem.path,
                         name:subItem.name,
                         meta:{
                            icon:subItem.web_icon,
                            title:subItem.name
                         },
                         component:modules[`@/views/${subItem.component_name}`]
                    })
                }
            })
        }
        routes.push(myRoute)
        // 组装动态路由地址 end
    })
    if(registerRouteFresh){
        // 添加动态路由
        routes.forEach(item=>{
            router.addRoute(item)
        })
        next({...to,replace:true})
        registerRouteFresh=false
    }else{
        next()

    }
    // 生成动态路由 end
    // next()
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