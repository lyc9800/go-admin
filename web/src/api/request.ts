import { useUserStore } from "@/store/modules/user";
import axios from "axios";
import type {AxiosError, AxiosRequestConfig} from "axios";

// 获取用户token
let userStore = localStorage.getItem('userStore')
let token = userStore!=null ? JSON.parse(userStore).token : ''
const service = axios.create({
    baseURL:import.meta.env.VITE_BASE_URL,
    timeout:100000,
    withCredentials:true
})

service.interceptors.request.use((config:AxiosRequestConfig)=>{ 
    const userStore = useUserStore()
    const token:string = userStore.token
    if(token){config.headers['AccessToken'] = token}
    return config
},(error:AxiosError)=>{
    return Promise.reject(new Error('请求错误'))
})

export default service