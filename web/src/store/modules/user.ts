import { defineStore } from "pinia";

export const useUserStore = defineStore("userStore", {
  state: () => {
    return{
        // 登陆token
        token: '',
        // 登陆用户信息
        userInfo: {
            avatar: '',
            nickName: '',
            sex: '',
        },
        // 角色
        roles:[]
    }
  },
  getters: {},
  actions: {
    // 设置登陆token
    setToken(token:string) {
        this.token = token
    },
    // 设置登陆用户的信息
    setUserInfo(userInfo:any) {
        this.userInfo = userInfo
    },
    // 更新部分信息
    setUserPartInfo(userInfo:any) {
        this.userInfo.avatar = userInfo.avatar
        this.userInfo.nickName = userInfo.nickName
        this.userInfo.sex = userInfo.sex
    },
  },
  persist:true,

})