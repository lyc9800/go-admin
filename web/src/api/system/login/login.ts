import  request  from "@/api/request";

// 登陆系统
export function loginApi(data: Object) { 
    return request({ 
        url: '/login/password',
        method: 'post',
        data
    })
}