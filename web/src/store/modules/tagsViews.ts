import router from "@/router";
import { defineStore } from "pinia";

export const useTagsViewsStore=defineStore('tagsViewsState',{ 
    state:()=>({
        activeTabsValue:'',
        visitedViews:[{path:'/home',name:'Home',meta:{title:'首页',affix:true},title:'首页'}]
    }),
    getters:{},
    actions:{
        setTabsMenuValue(val:string){
            this.activeTabsValue=val
        },
        addVisitedView(views:any){
            this.setTabsMenuValue(views.path)
            if(this.visitedViews.some(v=>v.path=== views.path)) return
            this.visitedViews.push(Object.assign({},views,{
                title:views.meta.title || 'no-name'
            }))
        },
        addView(view:any){
            this.addVisitedView(view)
        },
        delVisitedView(path:any){
            return new Promise(resolve=>{
                this.visitedViews=this.visitedViews.filter(v=>{
                    return (v.path !== path || v.meta.affix)
                })
                resolve([...this.visitedViews])
            })
        },
        delView(activeTabPath:any){
            return new Promise(resolve=>{
                this.delVisitedView(activeTabPath)
                resolve({
                    visitedViews:[...this.visitedViews]
                })
            })
        },
        delOtherViews(path:any){
            this.visitedViews=this.visitedViews.filter(item=>{
                return item.path===path || item.meta.affix
            })
        },
        delAllViews(){
            return new Promise(resolve=>{
                this.visitedViews=this.visitedViews.filter(v=>v.meta.affix)
                resolve([...this.visitedViews])
            })
        },
        goHome(){
            this.activeTabsValue='/home'
            router.push({path:'/home'})
        },
    }
})