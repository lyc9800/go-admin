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
