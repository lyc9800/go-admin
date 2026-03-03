import axios from "axios";
const service = axios.create({
    baseURL:'http://localhost:9999',
    timeout:100000
})

export default service