import  request  from "@/api/request";


export function getRoleListApi(params:object){
    return request({
        url:'role',
        method:'get',
        params
    })
}