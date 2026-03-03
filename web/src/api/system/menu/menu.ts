import  request  from '../../request'

// 获取树形菜单列表
export function getMenuListApi(params?: object){
    return request({
        url:'http://localhost:5173/menu.json',
        method:'get',
        params
    })
}