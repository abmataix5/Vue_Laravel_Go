import ApiGo from '@/services/Api.go'

 export default {

  getAll() {
    return ApiGo().get(`user/`)
  },
  getOneWorker(id) {
    return ApiGo().get(`user/${id}`)
  },
  createWorker(data) {
    console.log(data);
    return ApiGo().post('user/', data)
  },
  deleteWorker(id) {
    return ApiGo().delete(`user/${id}`)
  },
  updateWorker(id, data) {
    return ApiGo().put(`user/${id}`, data)
  }
 } 



