import axios from 'axios'

export default () => {
  const axiosInstance = axios.create({
    baseURL: "http://0.0.0.0:4000/api"
  })


  const token = localStorage.getItem('token')
  if (token) {
    axiosInstance.defaults.headers.common.Authorization = `Bearer ${token}`
    console.log(token);
  } 

  axiosInstance.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response.status === 401) {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        location.reload()
      }
      return Promise.reject(error) 
    }
  )

  return axiosInstance
}