import request  from "@/api/request";


// 获取用户列表数据
export function getUserListApi(params:object){
    return request({
        url:'user',
        method:'get',
        params
    })
}
// 新增管理员信息
export function addUserApi(data:object){
    return request({
        url:'user',
        method:'post',
        data
    })
}
// 获取用户信息
export function getUserApi(id:number){
    return request({
        url:`user/detail/${id}`,
        method:'get'
    })
}
// 修改用户信息
export function editUserApi(data:object){
    return request({
        url:'user',
        method:'put',
        data
    })
}
export function delUserApi(id:number){
    return request({
        url:`user/${id}`,
        method:'delete',
    })
}
// 获取角色列表
export function getAllRoleListAPi() {
    return request({
        url: 'role/all',
        method: 'get',
    })
}
// 修改用户信息
export function updateUserInfoApi(data:object){
    return request({
        url: 'user/info',
        method: 'put',
        data
    })
}
// 发送邮箱验证码
export function sendEmailApi(email:string){
    return request({
        url: 'user/sendemail',
        method: 'get',
        params: {
            email: email
        }
    })
}

export function verifyCode(email:string, code:string){
    return request({
        url: 'user/verifycode',
        method: 'get',
        params: {
            email: email,
            code: code
        }
    })
}
export function changeUserEmail(email: string, code: string) {
  return request({
    url: 'user/changeemail',  // 假设后端接口是这个
    method: 'put',  // 或 'post'，根据后端设计
    data: {  // 如果是PUT/POST请求，通常用data
      email: email,
      code: code
    }
  })
}
// 修改密码
export function updatePasswordApi(data: { oldPassword: string, newPassword: string }) {
  return request({
    url: '/user/change-password',
    method: 'put',
    data
  })
}
