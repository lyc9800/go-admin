import  request  from '../../request'

// 获取树形菜单列表
export function getMenuListApi(params?: object){
    return request({
        url:'menu',
        method:'get',
        params
    })
}
// 增加菜单
export function addMenuApi(data:object){
    return request({
        url:'menu',
        method:'post',
        data
    })
}
// 编辑菜单
export function editMenuApi(data:object){
    return request({
        url:'menu',
        method:'put',
        data
    })
}
// 删除菜单
export function delMenuApi(id:number){
    return request({
        url:`menu/${id}`,
        method:'delete',
    })
}