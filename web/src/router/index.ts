import Nprogress from '@/config/nprogress'
// 1.导入vue-router模块
import {createRouter, createWebHashHistory} from 'vue-router'
import {useMenuStore} from '@/store/modules/menu'
import {useUserStore} from '@/store/modules/user'

// 定义路由和组件映射关系
const modules = import.meta.glob("@/views/**/**.vue")


// 2.定义一些路由地址，每个都需要映射到一个组件
const routes = [
    {
        path: '/',
        redirect: '/login',
        meta: {title: '后台管理系统-登陆'},
        component: () => import('../views/system/login/Login.vue')
    },
    {
        path: '/login',
        name: 'Login',
        meta: {title: '后台管理系统-登陆'},
        component: () => import('../views/system/login/Login.vue')
    },
    {
        path: '/index',
        name: 'Index',
        component: () => import('@/views/system/layout/Index.vue'),
        redirect: '/home',
        children: [
            {
                path: '/home',
                name: 'Home',
                meta: {title: '首页', icon: 'House'},
                component: () => import('@/views/system/home/Index.vue')
            }
        ]
    },
]

// 3.创建路由实例并传递'routes'配置
const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

// 防止首次或者刷新界面路由失效
let registerRouteFresh = true
// 设置白名单
const whiteList = ['/login']

// 辅助函数：动态导入组件
const loadComponent = (componentPath: string) => {
    if (!componentPath) return null
    
    const fullPath = `/src/views/${componentPath}`
    
    
    const importFn = modules[fullPath]
    if (importFn) {
        return importFn
    } else {
        
        // 返回一个默认组件
        return () => import('@/views/system/error/NotFound.vue')
    }
}

// 辅助函数：递归生成动态路由
const generateDynamicRoutes = (menus: any[], parentPath = '') => {
    const dynamicRoutes: any[] = []
    
    menus.forEach(menu => {
        // 处理父级菜单
        if (menu.level === 0) {
            const route: any = {
                path: menu.path,
                name: menu.name,
                meta: {
                    icon: menu.web_icon,
                    title: menu.name
                },
                component: () => import('@/views/system/layout/Index.vue'),
                children: []
            }
            
            // 如果有子菜单
            if (menu.sub_menus && menu.sub_menus.length > 0) {
                menu.sub_menus.forEach((subMenu: any) => {
                    route.children.push({
                        path: subMenu.path,
                        name: subMenu.name,
                        meta: {
                            icon: subMenu.web_icon,
                            title: subMenu.name
                        },
                        component: loadComponent(subMenu.component_name)
                    })
                })
            }
            
            dynamicRoutes.push(route)
        } 
        // 处理一级菜单（没有父级但有组件）
        else if (menu.level === 1 && menu.component_name) {
            // 找到对应的父级路由（path 去除最后一段）
            const parentPath = menu.path.substring(0, menu.path.lastIndexOf('/'))
            
            // 查找是否已存在父级路由
            let parentRoute = dynamicRoutes.find(r => r.path === parentPath)
            
            if (parentRoute) {
                // 如果父级存在，添加到父级的children
                parentRoute.children.push({
                    path: menu.path,
                    name: menu.name,
                    meta: {
                        icon: menu.web_icon,
                        title: menu.name
                    },
                    component: loadComponent(menu.component_name)
                })
            } else {
                // 如果父级不存在，创建一个带有布局的路由
                dynamicRoutes.push({
                    path: parentPath,
                    name: `${menu.name}Parent`,
                    component: () => import('@/views/system/layout/Index.vue'),
                    children: [{
                        path: menu.path,
                        name: menu.name,
                        meta: {
                            icon: menu.web_icon,
                            title: menu.name
                        },
                        component: loadComponent(menu.component_name)
                    }]
                })
            }
        }
    })
    
    return dynamicRoutes
}

// 路由拦截守卫
router.beforeEach(async (to, from, next) => {

    
    // 1.Nprogress开始
    Nprogress.start()

    // 如果是白名单的路径，直接放行
    const isWhiteList = whiteList.some(item => to.path.indexOf(item) !== -1)
    
    if (isWhiteList) {
        
        return next()
    } else {
        // 判断是否已经登陆
        const userStore = useUserStore()
        if (!userStore.token) {
            
            return next({path: `/login?redirect=${to.path}`, replace: true})
        }
    }

    // 获取菜单信息
    const menuStore = useMenuStore()
    

    // 如果routers为空，获取菜单数据
    if (menuStore.routers.length === 0) {
        
        await menuStore.generateRouter()
    }
    
    

    // 生成并添加动态路由
    if (registerRouteFresh && menuStore.routers.length > 0) {
        
        
        // 生成动态路由
        const dynamicRoutes = generateDynamicRoutes(menuStore.routers)
        
        
        // 添加动态路由
        dynamicRoutes.forEach(route => {
            router.addRoute('Index', route) // 添加到 Index 路由的 children 中
        })
        
        // 添加404路由（可选）
        if (!router.hasRoute('404')) {
            router.addRoute({
                path: '/:pathMatch(.*)*',
                name: '404',
                component: () => import('@/views/system/error/404.vue')
            })
        }
        
        registerRouteFresh = false
        
        // 重定向当前路由以激活新添加的路由
        return next({...to, replace: true})
    }
    
    next()
})

// 路由跳转结束
router.afterEach(() => {
    Nprogress.done()
})

// 路由跳转失败
router.onError(error => {
    Nprogress.done()
    
})

export default router