import { tr } from 'element-plus/es/locales.mjs'
import {defineStore} from 'pinia'

export const useSettingStore = defineStore('settingState',{
    state:()=>({
        // 菜单是否收缩
        isCollapse:false,
    }),
    getters:{},
    actions:{
        // 切换Collapse
        setCollapse(value:boolean){
            this.isCollapse = value
        }
    },
})