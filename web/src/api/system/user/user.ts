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