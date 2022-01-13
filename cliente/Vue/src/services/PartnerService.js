import Apilaravel from "@/services/Api.laravel";

export default{

    get() {
      return Apilaravel().get(`users/`);
    },
    getOne(id) {
      return Apilaravel().get(`users/${id}`);
    },
    add(params) {
  
      return Apilaravel().post(`users/`, params);
    },
    update(id, params) {
      return Apilaravel().put(`users/${id}`, params); 
    },
    delete(id) {
   
      return Apilaravel().delete(`users/${id}`);
    }
  };