import axios from 'axios'

export default () => {
  const axiosInstance = axios.create({
    baseURL: "http://0.0.0.0:4000/api"
  })

/* Enviamos token */
  const admin = localStorage.getItem('admin')
  const token = localStorage.getItem('token')
  if (token ) {
    axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token}`
    axiosInstance.defaults.headers.common['Admin'] = `Admin ${admin}`
    console.log(token);
 
  } 

/* Enviamos usuario */

 
  
 
  axiosInstance.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response.status === 401) {
        localStorage.removeItem('token')
        localStorage.removeItem('admin')
        location.reload()
      }
      return Promise.reject(error) 
    }
  )

  return axiosInstance
}