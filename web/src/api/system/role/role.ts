import  request  from "@/api/request";


export function getRoleListApi(params:object){
    return request({
        url:'role',
        method:'get',
        params
    })
}
// 新建角色
export function addRoleApi(data:object){
    return request({
        url:'role',
        method:'post',
        data
    })
}
// 根据ID获取角色信息
export function getRoleApi(id:number){
    return request({
        url:`role/detail/${id}`,
        method:'get',
    })
}
// 编辑角色
export function editRoleApi(data:object){
    return request({
        url:'role',
        method:'put',
        data
    })
}
// 删除角色
export function delRoleApi(id:number){
    return request({
        url:`role/${id}`,
        method:'delete',
    })
}
// 修改管理员状态
export function changeIsAdminApi(id:number,is_admin:number){
    return request({
        url:`role/${id}/${is_admin}`,
        method:'patch',
    })
}