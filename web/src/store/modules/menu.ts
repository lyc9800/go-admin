import { getMenuListApi } from "@/api/system/menu/menu";
import { defineStore } from "pinia";

export const useMenuStore = defineStore('menuState',{
    state:()=>({
        register:false, //路由是否注册
        routers:[],     //路由列表
    }),
    getters:{},
    actions:{ 
        // 生成路由
        generateRouter: async function() {
            const {data}=await getMenuListApi()
            console.log('API 返回数据:', data)
            this.routers=data.result
            return data.result
        }
    },
    persist:true
})