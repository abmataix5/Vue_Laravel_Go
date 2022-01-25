import axios from 'axios'

export default () => {
  const axiosInstance = axios.create({
    baseURL: "http://0.0.0.0:8001/api"
  })

  /* Enviamos token */

  const token = localStorage.getItem('token')
  if (token ) {
    axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token}`
    console.log("Token enviado");
    console.log(token);

  } 

  axiosInstance.interceptors.response.use(
    (response) => response,
    (error) => {
      console.log("VALOR ERROR RECIBIDO");
      console.log(error);
      console.log(error.message);
      
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