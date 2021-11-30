import ApiGo from '@/services/Api.go'

 export default {

  getAll() {
    return ApiGo().get(`workers/`)
  },
  getOneWorker(id) {
    return ApiGo().get(`workers/${id}`)
  },
  createWorker(data) {
    console.log(data);
    return ApiGo().post('workers/', data)
  },
  deleteWorker(id) {
    return ApiGo().delete(`workers/${id}`)
  },
  updateWorker(id, data) {
    return ApiGo().put(`workesr/${id}`, data)
  }
 } 



