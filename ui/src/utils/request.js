import axios from 'axios'
var baseURL=import.meta.env.VITE_API_URL
// 创建一个 axios 实例
const service = axios.create({
	baseURL: baseURL, // 所有的请求地址前缀部分
	timeout: 60000, // 请求超时时间毫秒
	withCredentials: true, // 异步请求携带cookie
})

// 添加请求拦截器
service.interceptors.request.use(
	function (config) {
		// 在发送请求之前做些什么
		return config
	},
	function (error) {
		// 对请求错误做些什么
		return Promise.reject(error)
	}
)

// 添加响应拦截器
service.interceptors.response.use(
	function (response) {
		const dataAxios = response.data
		return dataAxios
	},
	function (error) {
		console.log(error)
		return Promise.reject(error)
	}
)
export default service
