import axios from 'axios'

export default () => {
  const axiosInstance = axios.create({
    baseURL: "http://0.0.0.0:4000/api"
  })


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