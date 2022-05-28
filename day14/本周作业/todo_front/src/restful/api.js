import Axios from 'axios'

// Axios.defaults.baseURL='http://127.0.0.1:8999/'
Axios.defaults.baseURL='http://121.5.37.237:8999/'

// 添加请求拦截器
/* Axios.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    if (localStorage.getItem('access_token')) {
    	// Axios.defaults.headers.common['Authorization'] = localStorage.getItem('access_token');
    	// console.log(config.headers);
    	config.headers.Authorization = localStorage.getItem('access_token')
        console.log(config.headers);
    }
    return config;
  }, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }); */

//axios携带token
// 添加请求拦截器，在请求头中加token
Axios.interceptors.request.use(
    config => {
        if (localStorage.getItem('access_token')) {
            config.headers.Authorization = "Bearer " + localStorage.getItem('access_token');
            console.log(config)
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    });


// 登录
export const userLogin = (params)=>{
	return Axios.post('login',params).then(res=>res.data);
}
// 获取所有代办事项 
export const getTodolist = ()=>{
    return Axios.get('api/v1/todo').then(res=>res.data)
}
// 修改事项状态
export const editTodo = (params) =>{
    return Axios.put('api/v1/todo',params).then(res=>res.data)
}
// 添加事项
export const addTodo = (params) =>{
    return Axios.post('api/v1/todo',params).then(res=>res.data)
}
// 删除事项
export const delTodo = (id) =>{
    return Axios.delete('api/v1/todo/'+id).then(res=>res.data)
}

// 注册
export const register = (params)=>{
	return Axios.post('register',params).then(res=>res.data)
}
// 退出
export const logout = ()=>{
	return Axios.post('logout').then(res=>res.data)
}

