import ApiGo from '@/services/Api.go'
import ApiLaravel from './Api.laravel';

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
    return ApiGo().put(`workers/${id}`, data)
  },
  loginWorker(data){
    return ApiGo().post('login', data)
  },
  loginLaravelWorker(params){
    console.log("userService.loginLaravelWorkers")
    console.log(params)
    return ApiLaravel().post(`login/`, params)
  },
 } 



