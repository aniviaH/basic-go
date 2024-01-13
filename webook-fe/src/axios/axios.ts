import axios from "axios";

const instance = axios.create({
    // baseURL: "http://localhost:8080",
    baseURL: "http://localhost:81",
    withCredentials: true
})


instance.interceptors.response.use(function (resp) {
    const newToken = resp.headers["x-jwt-token"]
    console.log("resp headers", resp.headers)
    console.log("token" + newToken)
    console.log("resp", resp)
    if (newToken) {
        localStorage.setItem("token", newToken)
    }
    if (resp.status == 401) { 
        alert("请登录")
        // window.location.href="/users/login"
    }
    return resp
}, (err) => {
    console.log('http error:', err)
    if (err.response && err.response.status == 401) {
        alert("请登录")
        // window.location.href="/users/login"
    }
    return err
})

// 在这里让每一个请求都加上 authorization 的头部
instance.interceptors.request.use((req) => {
    const token = localStorage.getItem("token")
    req.headers.setAuthorization("Bearer " + token, true)
    return req
}, (err) => {
    console.log(err)
})

export default instance