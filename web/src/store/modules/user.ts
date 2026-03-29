import { defineStore } from "pinia";
import { getUserApi } from '@/api/system/user/user'
// 删除错误的 import: import { id } from "element-plus/es/locale/index.mjs";

export const useUserStore = defineStore("userStore", {
  state: () => {
    return {
      // 登陆token
      token: '',
      // 登陆用户信息
      userInfo: {
        id: 0,  // 添加 id 字段，用于刷新时调用接口
        avatar: '',
        userName: '',
        sex: '',
        email: '',
        role_name: '',
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
      this.userInfo = { ...this.userInfo, ...userInfo }
    },
    // 更新部分信息
    setUserPartInfo(userInfo: any) {
      this.userInfo.avatar = userInfo.avatar ?? this.userInfo.avatar
      this.userInfo.userName = userInfo.userName ?? userInfo.username ?? this.userInfo.userName
      this.userInfo.sex = userInfo.sex ?? this.userInfo.sex
      this.userInfo.email = userInfo.email ?? this.userInfo.email
      this.userInfo.role_name = userInfo.role_name ?? userInfo.roleName ?? this.userInfo.role_name
    },
    // 刷新用户信息
    async refreshUserInfo() {
      try {
        // 使用 userInfo 中的 id
        const userId = this.userInfo.id
        if (!userId) {
          console.warn('用户ID不存在，无法刷新用户信息')
          return false
        }
        
        const response = await getUserApi(userId)
        if (response.data.code === 200) {
          this.setUserPartInfo({
            avatar: response.data.result.avatar || '',
            userName: response.data.result.username || response.data.result.userName,
            sex: response.data.result.sex || '',
            email: response.data.result.email || '',
            role_name: response.data.result.role_name || response.data.result.roleName || ''  // 添加 role_name
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