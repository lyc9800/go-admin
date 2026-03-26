import { defineStore } from "pinia";
import { getUserApi } from '@/api/system/user/user'
import { id } from "element-plus/es/locale/index.mjs";

export const useUserStore = defineStore("userStore", {
  state: () => {
    return {
      // 登陆token
      token: '',
      // 登陆用户信息
      userInfo: {
        avatar: '',
        userName: '',
        sex: '',
        email: '',
      },
      // 角色
      roles: []
    }
  },
  getters: {},
  actions: {
    // 设置登陆token
    setToken(token: string) {
      this.token = token
    },
    // 设置登陆用户的信息
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
    },
    // 更新部分信息
    setUserPartInfo(userInfo: any) {
      this.userInfo.avatar = userInfo.avatar
      this.userInfo.userName = userInfo.userName
      this.userInfo.sex = userInfo.sex
      this.userInfo.email = userInfo.email
    },
    // 刷新用户信息
    async refreshUserInfo() {
      try {
        const response = await getUserApi(id)
        if (response.data.code === 200) {
          this.setUserPartInfo({
            avatar: response.data.result.avatar || '',
            userName: response.data.result.username || response.data.result.userName,
            sex: response.data.result.sex || '',
            email: response.data.result.email || ''
          })
          return true
        }
        return false
      } catch (error) {
        console.error('刷新用户信息失败:', error)
        return false
      }
    }
  },
  persist: true,
})