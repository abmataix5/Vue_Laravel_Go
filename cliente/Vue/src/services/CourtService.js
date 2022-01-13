import Apilaravel from "@/services/Api.laravel";

export default{

    get() {
      return Apilaravel().get(`courts/`);
    },
    add(params) {
 
      return Apilaravel().post(`courts/`, params);
    },
    update(id, params) {
      return Apilaravel().put(`courts/${id}`, params); 
    },
    delete(id) {

      return Apilaravel().delete(`courts/${id}`);
    }
  };