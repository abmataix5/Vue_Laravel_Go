import axios from 'axios'

export default () => {
  const axiosInstance = axios.create({
    baseURL: "http://0.0.0.0:8000/api"
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
      if (error.response.status === 202) {
        // localStorage.removeItem('token')
        // localStorage.removeItem('user')
        // location.reload()
        console.log("usuario eliminado")
      }

      if (error.response.status === 500) {
        // localStorage.removeItem('token')
        // localStorage.removeItem('user')
        // location.reload()
        console.log("usuario Creado")
      }
      return Promise.reject(error) 
    }
  )

  return axiosInstance
}