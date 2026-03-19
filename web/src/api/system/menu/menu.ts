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